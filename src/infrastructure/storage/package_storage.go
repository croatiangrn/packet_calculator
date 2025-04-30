package storage

import (
	"database/sql"
	"github.com/croatiangrn/packet_calculator/src/domain/pack"
)

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

func (p *PackageStorage) GetAll() ([]pack.Pack, error) {
	query := "SELECT id, item_size FROM packs ORDER BY id ASC"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var packs []pack.Pack
	for rows.Next() {
		var p pack.Pack
		if err := rows.Scan(&p.ID, &p.ItemSize); err != nil {
			return nil, err
		}
		packs = append(packs, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return packs, nil
}

func (p *PackageStorage) AddPackSize(size int) error {
	query := "INSERT INTO packs (item_size) VALUES (?)"
	_, err := p.db.Exec(query, size)
	if err != nil {
		return err
	}
	return nil
}
