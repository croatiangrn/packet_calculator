package repositories

import "github.com/croatiangrn/packet_calculator/src/domain/pack"

// PackageRepository defines the interface for package repositories
type PackageRepository interface {
	GetPacks() ([]int, error)
	GetAll() ([]pack.Pack, error)
	AddPackSize(size int) error
}
