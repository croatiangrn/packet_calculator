package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/croatiangrn/packet_calculator/src/container"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/config"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/database"
	appHttp "github.com/croatiangrn/packet_calculator/src/infrastructure/http"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Load("./")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
		return
	}

	spew.Dump(cfg)

	dbDsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := database.InitDB(dbDsn)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
		return
	}

	if err := database.RunMigrations(db, "./migrations"); err != nil {
		log.Fatalf("Error running migrations: %v\n", err)
		return
	}

	appContainer := container.NewContainer(db)

	appHttp.SetContainer(appContainer)
	routerHandler := appHttp.InitRouter()

	srv := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: routerHandler,
	}

	go func() {
		log.Printf("Server starting on %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}
