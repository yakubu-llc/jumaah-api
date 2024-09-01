package musalah

import (
	"github.com/uptrace/bun"
	"github.com/yakubu-llc/jummah-api/internal/entities/shared"
)

// Musalah represents a musalah entity.
type Musalah struct {
	bun.BaseModel `bun:"table:musalahs"`

	ID   int    `json:"id"`
	Name string `json:"name"`
	shared.Timestamps
}

// CreateParams contains the parameters for creating a new musalah.
type CreateParams struct {
	Name string `json:"name"`
}

// UpdateParams contains the parameters for updating a musalah.
type UpdateParams struct {
	Name string `json:"name"`
}
