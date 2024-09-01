package musalah

import (
	"context"
	"database/sql"
	"log"

	"github.com/yakubu-llc/jumaah-api/internal/entities/musalah"
	"github.com/yakubu-llc/jumaah-api/internal/storage/postgres/shared"

	"github.com/uptrace/bun"
)

type MusalahRepository struct {
	db  bun.IDB
	ctx context.Context
}

// NewAccountRepository returns a new instance of the repository.
func NewMusalahRepository(db bun.IDB, ctx context.Context) *MusalahRepository {
	return &MusalahRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *MusalahRepository) Create(ctx context.Context, params musalah.CreateParams) (musalah.Musalah, error) {
	resp := musalah.Musalah{}

	err := r.db.
		NewInsert().
		Model(&params).
		ModelTableExpr("musalahs").
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *MusalahRepository) Update(ctx context.Context, id int, params musalah.UpdateParams) (musalah.Musalah, error) {
	resp := musalah.Musalah{}

	log.Println(r.db.
		NewUpdate().
		Model(&params).
		ModelTableExpr("musalahs").
		Where("id = ?", id).
		Returning("*").
		OmitZero().String())
	err :=
		r.db.
			NewUpdate().
			Model(&params).
			ModelTableExpr("musalahs").
			Where("id = ?", id).
			Returning("*").
			OmitZero().
			Scan(ctx, &resp)

	return resp, err
}

func (r *MusalahRepository) Delete(ctx context.Context, id int) error {
	res, err :=
		r.db.
			NewDelete().
			Model(&musalah.Musalah{}).
			Where("id = ?", id).
			Exec(ctx)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *MusalahRepository) GetById(ctx context.Context, id int) (musalah.Musalah, error) {
	resp := musalah.Musalah{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id = ?", id).
		Scan(ctx)

	return resp, err
}

func (r *MusalahRepository) GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]musalah.Musalah, error) {
	resp := []musalah.Musalah{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id >= ?", paginationParams.Cursor).
		Order("id").
		Limit(paginationParams.Limit).
		Scan(ctx)

	return resp, err
}
