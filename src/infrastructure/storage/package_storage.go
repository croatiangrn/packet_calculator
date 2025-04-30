package storage

import "database/sql"

type PackageStorage struct {
	db *sql.DB
}

func NewPackageStorage(db *sql.DB) *PackageStorage {
	return &PackageStorage{
		db: db,
	}
}

func (p *PackageStorage) CalculatePacks(order int, packs []int) (map[int]int, error) {
	packsMap := make(map[int]int)
	for _, pack := range packs {
		if order >= pack {
			count := order / pack
			packsMap[pack] = count
			order = order % pack
		}
	}
	return packsMap, nil
}
