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

func (p *PackageStorage) GetPacks() ([]int, error) {
	query := "SELECT item_size FROM packs ORDER BY item_size ASC"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var packs []int
	for rows.Next() {
		var packSize int
		if err := rows.Scan(&packSize); err != nil {
			return nil, err
		}
		packs = append(packs, packSize)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return packs, nil
}
