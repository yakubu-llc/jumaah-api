package jumaah

import (
	"context"
	"database/sql"

	"github.com/yakubu-llc/jumaah-api/internal/entities/jumaah"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"

	"github.com/uptrace/bun"
)

type JumaahRepository struct {
	db  bun.IDB
	ctx context.Context
}

// NewAccountRepository returns a new instance of the repository.
func NewJumaahRepository(db bun.IDB, ctx context.Context) *JumaahRepository {
	return &JumaahRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *JumaahRepository) Create(ctx context.Context, params jumaah.CreateJumaahParams) (jumaah.Jumaah, error) {
	resp := jumaah.Jumaah{}

	err := r.db.
		NewInsert().
		Model(&params).
		ModelTableExpr("jumaahs").
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *JumaahRepository) Update(ctx context.Context, id int, params jumaah.UpdateJumaahParams) (jumaah.Jumaah, error) {
	resp := jumaah.Jumaah{}

	err :=
		r.db.
			NewUpdate().
			Model(&params).
			ModelTableExpr("jumaahs").
			Where("id = ?", id).
			Returning("*").
			OmitZero().
			Scan(ctx, &resp)

	return resp, err
}

func (r *JumaahRepository) Delete(ctx context.Context, id int) error {
	res, err :=
		r.db.
			NewDelete().
			Model(&jumaah.Jumaah{}).
			Where("id = ?", id).
			Exec(ctx)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *JumaahRepository) GetById(ctx context.Context, id int) (jumaah.Jumaah, error) {
	resp := jumaah.Jumaah{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id = ?", id).
		Scan(ctx)

	return resp, err
}

func (r *JumaahRepository) GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]jumaah.Jumaah, error) {
	resp := []jumaah.Jumaah{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id >= ?", paginationParams.Cursor).
		Order("id").
		Limit(paginationParams.Limit).
		Scan(ctx)

	return resp, err
}
