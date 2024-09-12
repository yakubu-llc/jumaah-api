package jumaah

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/server/http/handler/shared"
	"github.com/yakubu-llc/jumaah-api/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

type httpHandler struct {
	jumaahService service.JumaahService
	logger        *zap.Logger
}

func newHTTPHandler(jumaahService service.JumaahService, logger *zap.Logger) *httpHandler {
	return &httpHandler{
		jumaahService: jumaahService,
		logger:        logger,
	}
}

type SingleJumaahResponse struct {
	Body struct {
		shared.MessageResponse
		Jumaah *jumaah.Jumaah `json:"jumaah"`
	}
}

func (h *httpHandler) getByID(ctx context.Context, input *shared.PathIDParam) (*SingleJumaahResponse, error) {
	jumaah, err := h.jumaahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Jumaah not found")
		default:
			h.logger.Error("failed to fetch jumaah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the jumaah")
		}
	}

	resp := &SingleJumaahResponse{}
	resp.Body.Message = "Jumaah fetched successfully"
	resp.Body.Jumaah = &jumaah

	return resp, nil
}

type GetAllJumaahOutput struct {
	Body struct {
		shared.MessageResponse
		Jumaahs []jumaah.Jumaah `json:"jumaahs"`
		shared.PaginationResponse
	}
}

func (h *httpHandler) getAll(ctx context.Context, input *shared.PaginationRequest) (*GetAllJumaahOutput, error) {
	LIMIT := input.Limit + 1

	jumaahs, err := h.jumaahService.GetAll(ctx, LIMIT, input.Cursor)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Jumaahs not found")
		default:
			h.logger.Error("failed to fetch jumaahs", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the jumaahs")
		}
	}

	resp := &GetAllJumaahOutput{}
	resp.Body.Message = "Jumaahs fetched successfully"
	resp.Body.Jumaahs = jumaahs

	if len(jumaahs) == LIMIT {
		resp.Body.NextCursor = &jumaahs[len(jumaahs)-1].ID
		resp.Body.HasMore = true
		resp.Body.Jumaahs = resp.Body.Jumaahs[:len(resp.Body.Jumaahs)-1]
	}

	return resp, nil
}

type CreateJumaahInput struct {
	Body jumaah.CreateJumaahParams `json:"jumaah"`
}

func (h *httpHandler) create(ctx context.Context, input *CreateJumaahInput) (*SingleJumaahResponse, error) {
	jumaah, err := h.jumaahService.Create(ctx, input.Body)
	if err != nil {
		h.logger.Error("failed to create jumaah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while creating the jumaah")
	}

	resp := &SingleJumaahResponse{}
	resp.Body.Message = "Jumaah created successfully"
	resp.Body.Jumaah = &jumaah

	return resp, nil
}

type UpdateJumaahInput struct {
	shared.PathIDParam
	Body jumaah.UpdateJumaahParams `json:"jumaah"`
}

func (h *httpHandler) update(ctx context.Context, input *UpdateJumaahInput) (*SingleJumaahResponse, error) {
	_, err := h.jumaahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Jumaah not found")
		default:
			h.logger.Error("failed to fetch jumaah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the jumaah")
		}
	}

	jumaah, err := h.jumaahService.Update(ctx, input.ID, input.Body)

	if err != nil {
		h.logger.Error("failed to update jumaah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while updating the jumaah")
	}

	resp := &SingleJumaahResponse{}
	resp.Body.Message = "Jumaah updated successfully"
	resp.Body.Jumaah = &jumaah

	return resp, nil
}

type DeleteJumaahResponse struct {
	Body shared.MessageResponse
}

func (h *httpHandler) delete(ctx context.Context, input *shared.PathIDParam) (*DeleteJumaahResponse, error) {
	_, err := h.jumaahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Jumaah not found")
		default:
			h.logger.Error("failed to fetch jumaah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the jumaah")
		}
	}

	err = h.jumaahService.Delete(ctx, input.ID)
	if err != nil {
		h.logger.Error("failed to delete jumaah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while deleting the jumaah")
	}

	resp := &DeleteJumaahResponse{}
	resp.Body.Message = "Jumaah deleted successfully"

	return resp, nil
}

type GetAttendeeInput struct {
	shared.PathIDParam
	AccountID int `path:"account_id" minimum:"1"`
}

type GetAttendeeResponse struct {
	Body struct {
		shared.MessageResponse
		Attendee *jumaah.Attendee `json:"attendee"`
	}
}

func (h *httpHandler) getAttendee(ctx context.Context, input *GetAttendeeInput) (*GetAttendeeResponse, error) {
	attendee, err := h.jumaahService.GetAttendee(ctx, input.ID, input.AccountID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Attendee not found")
		default:
			h.logger.Error("failed to get attendee", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while getting the attendee")
		}
	}

	resp := &GetAttendeeResponse{}
	resp.Body.Message = "Attendee fetched successfully"
	resp.Body.Attendee = &attendee

	return resp, nil
}

type GetAttendeesInput struct {
	shared.PathIDParam
	shared.PaginationRequest
}

type GetAttendeesResponse struct {
	Body struct {
		shared.MessageResponse
		Attendees []jumaah.Attendee `json:"attendees"`
		shared.PaginationResponse
	}
}

func (h *httpHandler) getAttendees(ctx context.Context, input *GetAttendeesInput) (*GetAttendeesResponse, error) {
	attendees, err := h.jumaahService.GetAttendees(ctx, input.ID, input.PaginationRequest.Limit, input.PaginationRequest.Cursor)
	if err != nil {
		h.logger.Error("failed to get attendees", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while getting the attendees")
	}

	resp := &GetAttendeesResponse{}
	resp.Body.Message = "Attendees fetched successfully"
	resp.Body.Attendees = attendees

	if len(attendees) == input.Limit+1 {
		resp.Body.NextCursor = &attendees[len(attendees)-1].AccountID
		resp.Body.HasMore = true
		resp.Body.Attendees = resp.Body.Attendees[:len(resp.Body.Attendees)-1]
	}

	return resp, nil
}

type GetAttendeeCountInput struct {
	shared.PathIDParam
}

type GetAttendeeCountResponse struct {
	Body struct {
		shared.MessageResponse
		AttendeeCount int `json:"attendee_count"`
	}
}

func (h *httpHandler) getAttendeeCount(ctx context.Context, input *GetAttendeeCountInput) (*GetAttendeeCountResponse, error) {
	attendeeCount, err := h.jumaahService.GetAttendeeCount(ctx, input.ID)
	if err != nil {
		h.logger.Error("failed to get attendee count", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while getting the attendee count")
	}

	resp := &GetAttendeeCountResponse{}
	resp.Body.Message = "Attendee count fetched successfully"
	resp.Body.AttendeeCount = attendeeCount

	return resp, nil
}

type CreateAttendeeInput struct {
	Body jumaah.CreateAttendeeParams `json:"attendee"`
}

type CreateAttendeeResponse struct {
	Body struct {
		shared.MessageResponse
		Attendee *jumaah.Attendee `json:"attendee"`
	}
}

func (h *httpHandler) createAttendee(ctx context.Context, input *CreateAttendeeInput) (*CreateAttendeeResponse, error) {
	attendee, err := h.jumaahService.CreateAttendee(ctx, input.Body)
	if err != nil {
		h.logger.Error("failed to create attendee", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while creating the attendee")
	}

	resp := &CreateAttendeeResponse{}
	resp.Body.Message = "Attendee created successfully"
	resp.Body.Attendee = &attendee

	return resp, nil
}

type UpdateAttendeeInput struct {
	shared.PathIDParam
	AccountID int                         `query:"account_id" minimum:"1"`
	Body      jumaah.UpdateAttendeeParams `json:"attendee"`
}

type UpdateAttendeeResponse struct {
	Body struct {
		shared.MessageResponse
		Attendee *jumaah.Attendee `json:"attendee"`
	}
}

func (h *httpHandler) updateAttendee(ctx context.Context, input *UpdateAttendeeInput) (*UpdateAttendeeResponse, error) {
	attendee, err := h.jumaahService.UpdateAttendee(ctx, input.ID, input.AccountID, input.Body)
	if err != nil {
		h.logger.Error("failed to update attendee", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while updating the attendee")
	}

	resp := &UpdateAttendeeResponse{}
	resp.Body.Message = "Attendee updated successfully"
	resp.Body.Attendee = &attendee

	return resp, nil
}

type DeleteAttendeeInput struct {
	shared.PathIDParam
	AccountID int `query:"account_id" minimum:"1"`
}

type DeleteAttendeeResponse struct {
	Body shared.MessageResponse
}

func (h *httpHandler) deleteAttendee(ctx context.Context, input *DeleteAttendeeInput) (*DeleteAttendeeResponse, error) {
	err := h.jumaahService.DeleteAttendee(ctx, input.ID, input.AccountID)
	if err != nil {
		h.logger.Error("failed to delete attendee", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while deleting the attendee")
	}

	resp := &DeleteAttendeeResponse{}
	resp.Body.Message = "Attendee deleted successfully"

	return resp, nil
}
