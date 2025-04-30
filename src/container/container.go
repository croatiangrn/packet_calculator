package container

import (
	"database/sql"
	"github.com/croatiangrn/packet_calculator/src/controller/http"
	"github.com/croatiangrn/packet_calculator/src/domain/repositories"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/storage"
	"github.com/croatiangrn/packet_calculator/src/service"
)

type Container struct {
	DB                *sql.DB
	PackageRepo       repositories.PackageRepository
	PackageService    *service.Pack
	PackageController *http.PackageController
}

func NewContainer(db *sql.DB) *Container {
	c := &Container{DB: db}
	c.PackageRepo = storage.NewPackageStorage(c.DB)
	c.PackageService = service.NewPackService(c.PackageRepo)
	c.PackageController = http.NewPackageController(c.PackageService)
	return c
}
