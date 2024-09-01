package jumaah

import (
	"context"

	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"
)

func (s *JumaahService) GetAttendee(ctx context.Context, jumaahId int, accountId int) (jumaah.Attendee, error) {
	return s.repositories.Jumaah().GetAttendee(ctx, jumaahId, accountId)
}

func (s *JumaahService) GetAttendees(ctx context.Context, jumaahId int, limit int, cursor int) ([]jumaah.Attendee, error) {
	paginationParams := shared.PaginationRequest{
		Limit:  limit,
		Cursor: cursor,
	}

	return s.repositories.Jumaah().GetAttendees(ctx, jumaahId, paginationParams)
}

func (s *JumaahService) GetAttendeeCount(ctx context.Context, jumaahId int) (int, error) {
	return s.repositories.Jumaah().GetAttendeeCount(ctx, jumaahId)
}

func (s *JumaahService) CreateAttendee(ctx context.Context, params jumaah.CreateAttendeeParams) (jumaah.Attendee, error) {
	return s.repositories.Jumaah().CreateAttendee(ctx, params)
}

func (s *JumaahService) DeleteAttendee(ctx context.Context, jumaahId int, accountId int) error {
	return s.repositories.Jumaah().DeleteAttendee(ctx, jumaahId, accountId)
}

func (s *JumaahService) UpdateAttendee(ctx context.Context, jumaahId int, accountId int, params jumaah.UpdateAttendeeParams) (jumaah.Attendee, error) {
	return s.repositories.Jumaah().UpdateAttendee(ctx, jumaahId, accountId, params)
}
