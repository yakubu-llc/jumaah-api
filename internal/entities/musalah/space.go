package musalah

import (
	"github.com/uptrace/bun"
	"github.com/yakubu-llc/jumaah-api/internal/entities/shared"
)

// Musalah represents a musalah entity.
type Musalah struct {
	bun.BaseModel `bun:"table:musalahs"`

	ID   int    `json:"id"`
	Name string `json:"name"`
	shared.Timestamps
}

// CreateMusalahParams contains the parameters for creating a new musalah.
type CreateMusalahParams struct {
	Name string `json:"name"`
}

// UpdateMusalahParams contains the parameters for updating a musalah.
type UpdateMusalahParams struct {
	Name string `json:"name"`
}
