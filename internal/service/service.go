package service

import (
	"context"

	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
)

type MusalahService interface {
	Create(ctx context.Context, params musalah.CreateParams) (musalah.Musalah, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params musalah.UpdateParams) (musalah.Musalah, error)
	GetById(ctx context.Context, id int) (musalah.Musalah, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]musalah.Musalah, error)
}
type FileService interface {
	GeneratePresignedUrl(ctx context.Context, objectKey string) (string, error)
}

type HealthService interface {
	HealthCheck(ctx context.Context) error
}

// Service storage of all services.
type Service struct {
	MusalahService MusalahService
	HealthService  HealthService
}
