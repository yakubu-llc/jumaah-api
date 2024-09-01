package domain

import (
	"github.com/yakubu-llc/jumaah-api/internal/service"
	"github.com/yakubu-llc/jumaah-api/internal/service/domain/account"
	"github.com/yakubu-llc/jumaah-api/internal/service/domain/health"
	"github.com/yakubu-llc/jumaah-api/internal/service/domain/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/service/domain/musalah"
	"github.com/yakubu-llc/jumaah-api/internal/storage"
)

// NewService implementation for storage of all services.
func NewService(
	repositories storage.Repository,
) *service.Service {
	return &service.Service{
		MusalahService: musalah.NewMusalahService(repositories),
		AccountService: account.NewAccountService(repositories),
		HealthService:  health.NewHealthService(repositories),
		JumaahService:  jumaah.NewJumaahService(repositories),
	}
}
