package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/yakubu-llc/jumaah-api/internal/entities/account"
	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
)

type MusalahService interface {
	Create(ctx context.Context, params musalah.CreateMusalahParams) (musalah.Musalah, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params musalah.UpdateMusalahParams) (musalah.Musalah, error)
	GetById(ctx context.Context, id int) (musalah.Musalah, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]musalah.Musalah, error)
}

type JumaahService interface {
	Create(ctx context.Context, params jumaah.CreateJumaahParams) (jumaah.Jumaah, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params jumaah.UpdateJumaahParams) (jumaah.Jumaah, error)
	GetById(ctx context.Context, id int) (jumaah.Jumaah, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]jumaah.Jumaah, error)

	GetAttendee(ctx context.Context, jumaahId int, accountId int) (jumaah.Attendee, error)
	GetAttendees(ctx context.Context, jumaahId int, limit int, cursor int) ([]jumaah.Attendee, error)
	GetAttendeeCount(ctx context.Context, jumaahId int) (int, error)
	CreateAttendee(ctx context.Context, params jumaah.CreateAttendeeParams) (jumaah.Attendee, error)
	UpdateAttendee(ctx context.Context, jumaahId int, accountId int, params jumaah.UpdateAttendeeParams) (jumaah.Attendee, error)
	DeleteAttendee(ctx context.Context, jumaahId int, accountId int) error
}

type AccountService interface {
	Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error)
	GetById(ctx context.Context, id int) (account.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]account.Account, error)
}

type HealthService interface {
	HealthCheck(ctx context.Context) error
}

// Service storage of all services.
type Service struct {
	AccountService AccountService
	JumaahService  JumaahService
	MusalahService MusalahService
	HealthService  HealthService
}
