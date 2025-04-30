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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/calculate-packs", func(c *gin.Context) {})

	return router
}
