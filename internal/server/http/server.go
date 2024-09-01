package http

import (
	"net/http"

	"github.com/yakubu-llc/jumaah-api/internal/service"

	"go.uber.org/zap"
)

type Server struct {
	services   *service.Service
	apiName    string
	apiVersion string
	logger     *zap.Logger
}

func NewServer(
	services *service.Service,
	apiName, apiVersion string,
	logger *zap.Logger,
) *Server {
	return &Server{
		services:   services,
		apiName:    apiName,
		apiVersion: apiVersion,
		logger:     logger,
	}
}

func (s *Server) Serve(port string) error {
	router := s.routes()

	return http.ListenAndServe(port, router)
}
