package jumaah

import (
	"context"

	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/storage"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"
)

type JumaahService struct {
	repositories storage.Repository
}

// NewTestService returns a new instance of test service.
func NewJumaahService(repositories storage.Repository) *JumaahService {
	return &JumaahService{
		repositories: repositories,
	}
}

func (s *JumaahService) GetById(ctx context.Context, id int) (jumaah.Jumaah, error) {
	return s.repositories.Jumaah().GetById(ctx, id)
}

func (s *JumaahService) GetAll(ctx context.Context, limit int, cursor int) ([]jumaah.Jumaah, error) {
	paginationParams := shared.PaginationRequest{
		Limit:  limit,
		Cursor: cursor,
	}

	return s.repositories.Jumaah().GetAll(ctx, paginationParams)
}

func (s *JumaahService) Create(ctx context.Context, params jumaah.CreateJumaahParams) (jumaah.Jumaah, error) {
	return s.repositories.Jumaah().Create(ctx, params)
}

func (s *JumaahService) Delete(ctx context.Context, id int) error {
	return s.repositories.Jumaah().Delete(ctx, id)
}

func (s *JumaahService) Update(ctx context.Context, id int, params jumaah.UpdateJumaahParams) (jumaah.Jumaah, error) {
	return s.repositories.Jumaah().Update(ctx, id, params)
}
