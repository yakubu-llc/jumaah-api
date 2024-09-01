package musalah

import (
	"context"

	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
	"github.com/yakubu-llc/jumaah-api/internal/storage"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"
)

type MusalahService struct {
	repositories storage.Repository
}

// NewTestService returns a new instance of test service.
func NewMusalahService(repositories storage.Repository) *MusalahService {
	return &MusalahService{
		repositories: repositories,
	}
}

func (s *MusalahService) GetById(ctx context.Context, id int) (musalah.Musalah, error) {
	return s.repositories.Musalah().GetById(ctx, id)
}

func (s *MusalahService) GetAll(ctx context.Context, limit int, cursor int) ([]musalah.Musalah, error) {
	paginationParams := shared.PaginationRequest{
		Limit:  limit,
		Cursor: cursor,
	}

	return s.repositories.Musalah().GetAll(ctx, paginationParams)
}

func (s *MusalahService) Create(ctx context.Context, params musalah.CreateParams) (musalah.Musalah, error) {
	return s.repositories.Musalah().Create(ctx, params)
}

func (s *MusalahService) Delete(ctx context.Context, id int) error {
	return s.repositories.Musalah().Delete(ctx, id)
}

func (s *MusalahService) Update(ctx context.Context, id int, params musalah.UpdateParams) (musalah.Musalah, error) {
	return s.repositories.Musalah().Update(ctx, id, params)
}
