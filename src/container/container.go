package container

import (
	"database/sql"
	"github.com/croatiangrn/packet_calculator/src/controller/http"
	"github.com/croatiangrn/packet_calculator/src/domain/repositories"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/storage"
)

type Container struct {
	DB                *sql.DB
	PackageRepo       repositories.PackageRepository
	PackageController *http.PackageController
}

func NewContainer(db *sql.DB) *Container {
	c := &Container{DB: db}
	c.PackageRepo = storage.NewPackageStorage(c.DB)
	return c
}
