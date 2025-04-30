package repositories

type PackageRepository interface {
	GetPacks() ([]int, error)
}
