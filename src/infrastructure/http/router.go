package http

import (
	"github.com/croatiangrn/packet_calculator/src/container"
	"github.com/gin-gonic/gin"
	"net/http"
)

var appContainer *container.Container

func SetContainer(container *container.Container) {
	appContainer = container
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

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
