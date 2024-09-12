package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/yakubu-llc/jumaah-api/internal/entities/account"
	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"
)

type MusalahRepository interface {
	Create(ctx context.Context, params musalah.CreateMusalahParams) (musalah.Musalah, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]musalah.Musalah, error)
	Update(ctx context.Context, id int, params musalah.UpdateMusalahParams) (musalah.Musalah, error)
	GetById(ctx context.Context, id int) (musalah.Musalah, error)
}

type JumaahRepository interface {
	Create(ctx context.Context, params jumaah.CreateJumaahParams) (jumaah.Jumaah, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]jumaah.Jumaah, error)
	Update(ctx context.Context, id int, params jumaah.UpdateJumaahParams) (jumaah.Jumaah, error)
	GetById(ctx context.Context, id int) (jumaah.Jumaah, error)

	GetAttendee(ctx context.Context, jumaahId int, accountId int) (jumaah.Attendee, error)
	GetAttendees(ctx context.Context, jumaahId int, paginationParams shared.PaginationRequest) ([]jumaah.Attendee, error)
	GetAttendeeCount(ctx context.Context, jumaahId int) (int, error)
	CreateAttendee(ctx context.Context, params jumaah.CreateAttendeeParams) (jumaah.Attendee, error)
	UpdateAttendee(ctx context.Context, jumaahId int, accountId int, params jumaah.UpdateAttendeeParams) (jumaah.Attendee, error)
	DeleteAttendee(ctx context.Context, jumaahId int, accountId int) error
}

type AccountRepository interface {
	Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]account.Account, error)
	Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error)
	GetById(ctx context.Context, id int) (account.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error)
}

type RepositoryProvider interface {
	Musalah() MusalahRepository
	Jumaah() JumaahRepository
	Account() AccountRepository
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
