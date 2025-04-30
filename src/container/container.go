package container

import (
	"database/sql"
	"github.com/croatiangrn/packet_calculator/src/controller/http"
	"github.com/croatiangrn/packet_calculator/src/domain/repositories"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/storage"
	"github.com/croatiangrn/packet_calculator/src/service"
)

// TODO: Consider using a DI library for better management of dependencies e.g. Google Wire (https://github.com/google/wire)

// Container holds all the dependencies for the application
type Container struct {
	DB                   *sql.DB
	PackageRepo          repositories.PackageRepository
	PackageService       *service.Pack
	CalculatorService    *service.Calculator
	PacksController      *http.PacksController
	CalculatorController *http.CalculatorController
}

func NewContainer(db *sql.DB) *Container {
	c := &Container{DB: db}
	c.PackageRepo = storage.NewPackageStorage(c.DB)
	c.PackageService = service.NewPack(c.PackageRepo)
	c.CalculatorService = service.NewCalculator(c.PackageRepo)
	c.PacksController = http.NewPacksController(c.PackageService)
	c.CalculatorController = http.NewCalculatorController(c.CalculatorService)
	return c
}
