package main

import (
	"context"
	"errors"
	"github.com/croatiangrn/packet_calculator/src/container"
	appHttp "github.com/croatiangrn/packet_calculator/src/infrastructure/http"
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

	appContainer := container.NewContainer()

	appHttp.SetContainer(appContainer)
	routerHandler := appHttp.InitRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routerHandler,
	}

	go func() {
		log.Println("Server starting on :8080")
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
