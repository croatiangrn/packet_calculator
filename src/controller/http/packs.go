package http

import (
	"github.com/croatiangrn/packet_calculator/src/service"
	"github.com/gin-gonic/gin"
)

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
	}
}

func (pc *PackageController) CalculatePacks(ginCtx *gin.Context) {
	packSizes, err := pc.service.GetPacks()
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": "Failed to get packs"})
		return
	}

	orderRes, err := pc.service.CalculatePacks(1, packSizes)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(200, gin.H{
		"order": orderRes,
	})
}
