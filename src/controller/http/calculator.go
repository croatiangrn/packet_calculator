package http

import (
	"github.com/croatiangrn/packet_calculator/src/infrastructure/http/dto"
	"github.com/croatiangrn/packet_calculator/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalculatorController struct {
	service *service.Calculator
}

func NewCalculatorController(service *service.Calculator) *CalculatorController {
	return &CalculatorController{
		service: service,
	}
}

func (cc *CalculatorController) RegisterRoutes(router *gin.RouterGroup) {
	packageAPI := router.Group("/calculate")
	{
		packageAPI.POST("", cc.CalculatePacks)
	}
}

func (cc *CalculatorController) CalculatePacks(ginCtx *gin.Context) {
	var req dto.CalculatePacksRequest

	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	packSizes, err := cc.service.GetPacksSizes()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get packs"})
		return
	}

	orderRes, err := cc.service.CalculatePacks(req.Items, packSizes)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"order": orderRes,
	})
}
