package service

import (
	"fmt"
	"github.com/croatiangrn/packet_calculator/src/domain/pack"
	"github.com/croatiangrn/packet_calculator/src/domain/repositories"
	"github.com/croatiangrn/packet_calculator/src/infrastructure/http/dto"
)

type Pack struct {
	repo repositories.PackageRepository
}

func NewPack(repo repositories.PackageRepository) *Pack {
	return &Pack{
		repo: repo,
	}
}

type dpEntry struct {
	count    int
	lastPack int
}

func (p *Pack) GetPacksSizes() ([]int, error) {
	return p.repo.GetPacks()
}

func (p *Pack) GetAllPacks() ([]dto.PackResponse, error) {
	packs, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var packResponses []dto.PackResponse
	for _, singlePack := range packs {
		packResponses = append(packResponses, dto.PackResponse{
			ID:       singlePack.ID,
			ItemSize: singlePack.ItemSize,
		})
	}

	return packResponses, nil
}

func (p *Pack) AddPackSize(size int) error {
	if size <= 0 {
		return fmt.Errorf("size must be greater than zero")
	}

	packDomain := pack.Pack{
		ItemSize: size,
	}

	if err := packDomain.Validate(); err != nil {
		return err
	}

	return p.repo.AddPackSize(size)
}

func (p *Pack) DeletePackSize(id int) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater than zero")
	}

	return p.repo.DeletePackSize(id)
}
