package http

import (
	"github.com/croatiangrn/packet_calculator/src/infrastructure/http/dto"
	"github.com/croatiangrn/packet_calculator/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PackageController handles HTTP requests related to packages
type PackageController struct {
	service *service.Pack
}

func NewPackageController(pkgService *service.Pack) *PackageController {
	return &PackageController{
		service: pkgService,
	}
}

func (pc *PackageController) RegisterRoutes(router *gin.RouterGroup) {
	packageAPI := router.Group("/packages")
	{
		packageAPI.POST("/calculate", pc.CalculatePacks)

		// This should be separated to another controller to be RESTful
		// TODO: Move this to another controller
		packageAPI.GET("/sizes", pc.GetPackSizes)
		packageAPI.POST("/sizes", pc.AddPackSize)

	}
}

func (pc *PackageController) CalculatePacks(ginCtx *gin.Context) {
	var req dto.CalculatePacksRequest

	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	packSizes, err := pc.service.GetPacksSizes()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get packs"})
		return
	}

	orderRes, err := pc.service.CalculatePacks(1, packSizes)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"order": orderRes,
	})
}

func (pc *PackageController) GetPackSizes(ginCtx *gin.Context) {
	sizes, err := pc.service.GetAllPacks()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get packs"})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"sizes": sizes,
	})
}

func (pc *PackageController) AddPackSize(ginCtx *gin.Context) {
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
