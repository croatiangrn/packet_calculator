package repositories

type PackageRepository interface {
	CalculatePacks(order int, packs []int) (map[int]int, error)
}
