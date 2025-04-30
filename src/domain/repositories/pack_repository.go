package repositories

// PackageRepository defines the interface for package repositories
type PackageRepository interface {
	GetPacks() ([]int, error)
}
