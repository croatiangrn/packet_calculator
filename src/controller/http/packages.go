package http

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type PackageController struct {
	db *sql.DB
}

func NewPackageController(db *sql.DB) *PackageController {
	return &PackageController{
		db: db,
	}
}

func (pc *PackageController) RegisterRoutes(router *gin.RouterGroup) {
	packageAPI := router.Group("/packages")
	{
		packageAPI.POST("/calculate", pc.CalculatePacks)
	}
}

func (pc *PackageController) CalculatePacks(ginCtx *gin.Context) {

}
