package http

import (
	"github.com/croatiangrn/packet_calculator/src/container"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var appContainer *container.Container

func SetContainer(container *container.Container) {
	appContainer = container
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	// Configure CORS middleware (Gin's official implementation)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://calculator.mikulic.dev"}, // TODO: This should be moved to a config file
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1API := router.Group("/v1")
	{
		appContainer.CalculatorController.RegisterRoutes(v1API)
		appContainer.PacksController.RegisterRoutes(v1API)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
