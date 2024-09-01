package domain

import (
	"github.com/yakubu-llc/jummah-api/internal/service"
	"github.com/yakubu-llc/jummah-api/internal/service/domain/health"
	"github.com/yakubu-llc/jummah-api/internal/service/domain/musalah"
	"github.com/yakubu-llc/jummah-api/internal/storage"
)

// NewService implementation for storage of all services.
func NewService(
	repositories storage.Repository,
) *service.Service {
	return &service.Service{
		MusalahService: musalah.NewMusalahService(repositories),
		HealthService:  health.NewHealthService(repositories),
	}
}
