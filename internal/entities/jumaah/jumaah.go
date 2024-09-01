package jumaah

import (
	"time"

	"github.com/uptrace/bun"
	"github.com/yakubu-llc/jumaah-api/internal/entities/shared"
)

type JumaahStatus string

const (
	JumaahStatusSuggested JumaahStatus = "suggested"
	JumaahStatusScheduled JumaahStatus = "scheduled"
)

// Jumaah represents a jumaah entity.
type Jumaah struct {
	bun.BaseModel `bun:"table:jumaahs"`

	ID        int          `json:"id"`
	Name      string       `json:"name"`
	AccountID int          `json:"accountId"`
	MusalahID int          `json:"musalahId"`
	Status    JumaahStatus `json:"status"`
	BeginsAt  time.Time    `json:"beginsAt"`
	shared.Timestamps
}

type MutateStatusParam struct {
	Status JumaahStatus `json:"status" enum:"suggested,scheduled"`
}

// CreateJumaahParams contains the parameters for creating a new jumaah.
type CreateJumaahParams struct {
	Name      string `json:"name"`
	AccountID int    `json:"accountId" minimum:"1"`
	MusalahID int    `json:"musalahId" minimum:"1"`
	MutateStatusParam
	BeginsAt time.Time `json:"beginsAt"`
}

// UpdateJumaahParams contains the parameters for updating a jumaah.
type UpdateJumaahParams struct {
	Name      string `json:"name"`
	MusalahID int    `json:"musalah_id" minimum:"1"`
	MutateStatusParam
	BeginsAt time.Time `json:"begins_at"`
}
