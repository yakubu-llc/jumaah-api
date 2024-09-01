package jumaah

import (
	"context"
	"database/sql"

	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"
)

func (r *JumaahRepository) CreateAttendee(ctx context.Context, params jumaah.CreateAttendeeParams) (jumaah.Attendee, error) {
	resp := jumaah.Attendee{}

	err := r.db.
		NewInsert().
		Model(&params).
		ModelTableExpr("jumaah_attendees").
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *JumaahRepository) UpdateAttendee(ctx context.Context, jumaahId int, accountId int, params jumaah.UpdateAttendeeParams) (jumaah.Attendee, error) {
	resp := jumaah.Attendee{}

	err := r.db.
		NewUpdate().
		Model(&params).
		ModelTableExpr("jumaah_attendees").
		Where("jumaah_id = ?", jumaahId).
		Where("account_id = ?", accountId).
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *JumaahRepository) DeleteAttendee(ctx context.Context, jumaahId int, accountId int) error {
	res, err :=
		r.db.
			NewDelete().
			Model(&jumaah.Attendee{}).
			Where("jumaah_id = ?", jumaahId).
			Where("account_id = ?", accountId).
			Exec(ctx)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *JumaahRepository) GetAttendee(ctx context.Context, jumaahId int, accountId int) (jumaah.Attendee, error) {
	resp := jumaah.Attendee{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("jumaah_id = ?", jumaahId).
		Where("account_id = ?", accountId).
		Scan(ctx)

	return resp, err
}

func (r *JumaahRepository) GetAttendees(ctx context.Context, jumaahId int, paginationParams shared.PaginationRequest) ([]jumaah.Attendee, error) {
	resp := []jumaah.Attendee{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("account_id >= ?", paginationParams.Cursor).
		Where("jumaah_id = ?", jumaahId).
		Order("account_id").
		Limit(paginationParams.Limit).
		Scan(ctx)

	return resp, err
}

func (r *JumaahRepository) GetAttendeeCount(ctx context.Context, jumaahId int) (int, error) {
	count, err := r.db.
		NewSelect().
		Model(&jumaah.Attendee{}).
		Where("jumaah_id = ?", jumaahId).
		Count(ctx)

	return count, err
}
