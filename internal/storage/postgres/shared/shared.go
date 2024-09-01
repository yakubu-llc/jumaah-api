package shared

import (
	"github.com/uptrace/bun"
)

func ExcludeInsertColumns(query *bun.InsertQuery) *bun.InsertQuery {
	query.ExcludeColumn("created_at", "updated_at")
	return query
}

func ExcludeUpdateColumns(query *bun.UpdateQuery) *bun.UpdateQuery {
	query.ExcludeColumn("created_at", "updated_at", "id")
	return query
}

type PaginationRequest struct {
	Cursor int `json:"cursor" default:"1" min:"1" required:"false"`
	Limit  int `json:"limit" default:"10" min:"1" max:"100" required:"false"`
}