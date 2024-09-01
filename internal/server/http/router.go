package http

import (
	"github.com/yakubu-llc/jummah-api/internal/server/http/handler/health"
	"github.com/yakubu-llc/jummah-api/internal/server/http/handler/musalah"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() chi.Router {
	router := chi.NewMux()

	config := huma.DefaultConfig(s.apiName, s.apiVersion)
	// config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
	// 	"bearerAuth": {
	// 		Type:         "http",
	// 		Scheme:       "bearer",
	// 		BearerFormat: "JWT",
	// 	},
	// }

	humaApi := humachi.New(router, config)

	musalah.RegisterHumaRoutes(
		s.services.MusalahService,
		humaApi,
		s.logger,
	)

	health.RegisterHumaRoutes(
		s.services.HealthService,
		humaApi,
		s.logger,
	)

	return router
}
