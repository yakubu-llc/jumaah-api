package storage

import (
	"context"

	"github.com/yakubu-llc/jummah-api/internal/entities/musalah"
	"github.com/yakubu-llc/jummah-api/internal/storage/postgres/shared"
)

type MusalahRepository interface {
	Create(ctx context.Context, params musalah.CreateParams) (musalah.Musalah, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]musalah.Musalah, error)
	Update(ctx context.Context, id int, params musalah.UpdateParams) (musalah.Musalah, error)
	GetById(ctx context.Context, id int) (musalah.Musalah, error)
}

type RepositoryProvider interface {
	Musalah() MusalahRepository
}

type Transaction interface {
	RepositoryProvider
	Commit() error
	Rollback() error
	SubTransaction() (Transaction, error)
}

type Repository interface {
	RepositoryProvider
	HealthCheck(ctx context.Context) error
	NewTransaction() (Transaction, error)
	RunInTx(ctx context.Context, fn func(ctx context.Context, tx Transaction) error) error
}
