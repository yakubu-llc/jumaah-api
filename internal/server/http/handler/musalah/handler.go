package musalah

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
	"github.com/yakubu-llc/jumaah-api/internal/server/http/handler/shared"
	"github.com/yakubu-llc/jumaah-api/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

type httpHandler struct {
	musalahService service.MusalahService
	logger         *zap.Logger
}

func newHTTPHandler(musalahService service.MusalahService, logger *zap.Logger) *httpHandler {
	return &httpHandler{
		musalahService: musalahService,
		logger:         logger,
	}
}

type SingleMusalahResponse struct {
	Body struct {
		shared.MessageResponse
		Musalah *musalah.Musalah `json:"musalah"`
	}
}

func (h *httpHandler) getByID(ctx context.Context, input *shared.PathIDParam) (*SingleMusalahResponse, error) {
	user := shared.GetAuthenticatedUser(ctx)
	h.logger.Debug("user", zap.Any("user", user))

	musalah, err := h.musalahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Musalah not found")
		default:
			h.logger.Error("failed to fetch musalah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the musalah")
		}
	}

	resp := &SingleMusalahResponse{}
	resp.Body.Message = "Musalah fetched successfully"
	resp.Body.Musalah = &musalah

	return resp, nil
}

type GetAllMusalahOutput struct {
	Body struct {
		shared.MessageResponse
		Musalahs []musalah.Musalah `json:"musalahs"`
		shared.PaginationResponse
	}
}

func (h *httpHandler) getAll(ctx context.Context, input *shared.PaginationRequest) (*GetAllMusalahOutput, error) {
	LIMIT := input.Limit + 1

	musalahs, err := h.musalahService.GetAll(ctx, LIMIT, input.Cursor)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Musalahs not found")
		default:
			h.logger.Error("failed to fetch musalahs", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the musalahs")
		}
	}

	resp := &GetAllMusalahOutput{}
	resp.Body.Message = "Musalahs fetched successfully"
	resp.Body.Musalahs = musalahs

	if len(musalahs) == LIMIT {
		resp.Body.Cursor = &musalahs[len(musalahs)-1].ID
		resp.Body.HasMore = true
		resp.Body.Musalahs = resp.Body.Musalahs[:len(resp.Body.Musalahs)-1]
	}

	return resp, nil
}

type CreateMusalahInput struct {
	Body musalah.CreateMusalahParams `json:"musalah"`
}

func (h *httpHandler) create(ctx context.Context, input *CreateMusalahInput) (*SingleMusalahResponse, error) {
	musalah, err := h.musalahService.Create(ctx, input.Body)
	if err != nil {
		h.logger.Error("failed to create musalah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while creating the musalah")
	}

	resp := &SingleMusalahResponse{}
	resp.Body.Message = "Musalah created successfully"
	resp.Body.Musalah = &musalah

	return resp, nil
}

type UpdateMusalahInput struct {
	shared.PathIDParam
	Body musalah.UpdateMusalahParams `json:"musalah"`
}

func (h *httpHandler) update(ctx context.Context, input *UpdateMusalahInput) (*SingleMusalahResponse, error) {
	_, err := h.musalahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Musalah not found")
		default:
			h.logger.Error("failed to fetch musalah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the musalah")
		}
	}

	musalah, err := h.musalahService.Update(ctx, input.ID, input.Body)

	if err != nil {
		h.logger.Error("failed to update musalah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while updating the musalah")
	}

	resp := &SingleMusalahResponse{}
	resp.Body.Message = "Musalah updated successfully"
	resp.Body.Musalah = &musalah

	return resp, nil
}

type DeleteMusalahResponse struct {
	Body shared.MessageResponse
}

func (h *httpHandler) delete(ctx context.Context, input *shared.PathIDParam) (*DeleteMusalahResponse, error) {
	_, err := h.musalahService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Musalah not found")
		default:
			h.logger.Error("failed to fetch musalah", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the musalah")
		}
	}

	err = h.musalahService.Delete(ctx, input.ID)
	if err != nil {
		h.logger.Error("failed to delete musalah", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while deleting the musalah")
	}

	resp := &DeleteMusalahResponse{}
	resp.Body.Message = "Musalah deleted successfully"

	return resp, nil
}
