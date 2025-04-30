package http

import (
	"github.com/croatiangrn/packet_calculator/src/infrastructure/http/dto"
	"github.com/croatiangrn/packet_calculator/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PacksController handles HTTP requests related to packages
type PacksController struct {
	service *service.Pack
}

func NewPacksController(pkgService *service.Pack) *PacksController {
	return &PacksController{
		service: pkgService,
	}
}

func (pc *PacksController) RegisterRoutes(router *gin.RouterGroup) {
	packageAPI := router.Group("/packs")
	{
		packageAPI.GET("", pc.GetPackSizes)
		packageAPI.POST("", pc.AddPackSize)
		packageAPI.DELETE("/:id", pc.DeletePackSize)
	}
}

func (pc *PacksController) GetPackSizes(ginCtx *gin.Context) {
	sizes, err := pc.service.GetAllPacks()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get packs"})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"sizes": sizes,
	})
}

func (pc *PacksController) AddPackSize(ginCtx *gin.Context) {
	var req dto.AddPackSizeRequest

	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := pc.service.AddPackSize(req.ItemSize)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add pack size"})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": "Pack size added successfully"})
}

func (pc *PacksController) DeletePackSize(ginCtx *gin.Context) {
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pack ID"})
		return
	}

	if err := pc.service.DeletePackSize(id); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pack size"})
		return
	}

	ginCtx.JSON(http.StatusNoContent, nil)
}
