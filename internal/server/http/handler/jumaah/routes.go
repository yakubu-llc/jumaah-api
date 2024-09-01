package jumaah

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/yakubu-llc/jumaah-api/internal/service"
	"go.uber.org/zap"
)

func RegisterHumaRoutes(
	jumaahService service.JumaahService,
	humaApi huma.API,
	logger *zap.Logger,
) {

	handler := &httpHandler{
		jumaahService: jumaahService,
		logger:        logger,
	}

	// Register GET /test/{id}
	huma.Register(humaApi, huma.Operation{
		OperationID: "get-jumaah-by-id",
		Method:      http.MethodGet,
		Path:        "/jumaah/{id}",
		Summary:     "Get jumaah by ID",
		Description: "Get jumaah by ID.",
		Tags:        []string{"Jumaah"},
	}, handler.getByID)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-all-jumaahs",
		Method:      http.MethodGet,
		Path:        "/jumaah",
		Summary:     "Get all jumaahs",
		Description: "Get all jumaahs.",
		Tags:        []string{"Jumaah"},
	}, handler.getAll)

	huma.Register(humaApi, huma.Operation{
		OperationID: "create-jumaah",
		Method:      http.MethodPost,
		Path:        "/jumaah",
		Summary:     "Create a jumaah",
		Description: "Create a jumaah.",
		Tags:        []string{"Jumaah"},
	}, handler.create)

	huma.Register(humaApi, huma.Operation{
		OperationID: "update-jumaah",
		Method:      http.MethodPut,
		Path:        "/jumaah/{id}",
		Summary:     "Update a jumaah",
		Description: "Update a jumaah.",
		Tags:        []string{"Jumaah"},
	}, handler.update)

	huma.Register(humaApi, huma.Operation{
		OperationID: "delete-jumaah",
		Method:      http.MethodDelete,
		Path:        "/jumaah/{id}",
		Summary:     "Delete a jumaah",
		Description: "Delete a jumaah.",
		Tags:        []string{"Jumaah"},
	}, handler.delete)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-attendees",
		Method:      http.MethodGet,
		Path:        "/jumaah/{id}/attendees",
		Summary:     "Get attendees",
		Description: "Get attendees.",
		Tags:        []string{"Jumaah"},
	}, handler.getAttendees)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-attendee-count",
		Method:      http.MethodGet,
		Path:        "/jumaah/{id}/attendees/count",
		Summary:     "Get attendee count",
		Description: "Get attendee count.",
		Tags:        []string{"Jumaah"},
	}, handler.getAttendeeCount)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-attendee",
		Method:      http.MethodGet,
		Path:        "/jumaah/{id}/attendees/{account_id}",
		Summary:     "Get attendee",
		Description: "Get attendee.",
		Tags:        []string{"Jumaah"},
	}, handler.getAttendee)

	huma.Register(humaApi, huma.Operation{
		OperationID: "create-attendee",
		Method:      http.MethodPost,
		Path:        "/jumaah/attendees",
		Summary:     "Create attendee",
		Description: "Create attendee.",
		Tags:        []string{"Jumaah"},
	}, handler.createAttendee)

	huma.Register(humaApi, huma.Operation{
		OperationID: "update-attendee",
		Method:      http.MethodPut,
		Path:        "/jumaah/{id}/attendees/{account_id}",
		Summary:     "Update attendee",
		Description: "Update attendee.",
		Tags:        []string{"Jumaah"},
	}, handler.updateAttendee)

	huma.Register(humaApi, huma.Operation{
		OperationID: "delete-attendee",
		Method:      http.MethodDelete,
		Path:        "/jumaah/{id}/attendees/{account_id}",
		Summary:     "Delete attendee",
		Description: "Delete attendee.",
		Tags:        []string{"Jumaah"},
	}, handler.deleteAttendee)

}
