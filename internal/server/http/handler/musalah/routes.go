package musalah

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/yakubu-llc/jummah-api/internal/service"
	"go.uber.org/zap"
)

func RegisterHumaRoutes(
	musalahService service.MusalahService,
	humaApi huma.API,
	logger *zap.Logger,
) {

	handler := &httpHandler{
		musalahService: musalahService,
		logger:         logger,
	}

	// Register GET /test/{id}
	huma.Register(humaApi, huma.Operation{
		OperationID: "get-musalah-by-id",
		Method:      http.MethodGet,
		Path:        "/musalah/{id}",
		Summary:     "Get musalah by ID",
		Description: "Get musalah by ID.",
		Tags:        []string{"Musalah"},
	}, handler.getByID)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-all-musalahs",
		Method:      http.MethodGet,
		Path:        "/musalah",
		Summary:     "Get all musalahs",
		Description: "Get all musalahs.",
		Tags:        []string{"Musalah"},
	}, handler.getAll)

	huma.Register(humaApi, huma.Operation{
		OperationID: "create-musalah",
		Method:      http.MethodPost,
		Path:        "/musalah",
		Summary:     "Create a musalah",
		Description: "Create a musalah.",
		Tags:        []string{"Musalah"},
	}, handler.create)

	huma.Register(humaApi, huma.Operation{
		OperationID: "update-musalah",
		Method:      http.MethodPut,
		Path:        "/musalah/{id}",
		Summary:     "Update a musalah",
		Description: "Update a musalah.",
		Tags:        []string{"Musalah"},
	}, handler.update)

	huma.Register(humaApi, huma.Operation{
		OperationID: "delete-musalah",
		Method:      http.MethodDelete,
		Path:        "/musalah/{id}",
		Summary:     "Delete a musalah",
		Description: "Delete a musalah.",
		Tags:        []string{"Musalah"},
	}, handler.delete)

}
