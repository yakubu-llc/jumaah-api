package jumaah

import (
	"time"

	"github.com/uptrace/bun"
	"github.com/yakubu-llc/jumaah-api/internal/entities/shared"
)

// Attendee represents a jumaah attendee entity.
type Attendee struct {
	bun.BaseModel `bun:"table:jumaah_attendees"`

	AccountID  int        `json:"accountId"`
	JumaahID   int        `json:"jumaahId"`
	DelayedETA *time.Time `json:"delayedETA"`
	shared.Timestamps
}

// CreateAttendeeParams contains the parameters for creating a new attendee.
type CreateAttendeeParams struct {
	AccountID  int        `json:"accountId" minimum:"1"`
	JumaahID   int        `json:"jumaahId" minimum:"1"`
	DelayedETA *time.Time `json:"delayedETA"`
}

// UpdateAttendeeParams contains the parameters for updating an attendee.
type UpdateAttendeeParams struct {
	AccountID  int        `json:"accountId" minimum:"1"`
	JumaahID   int        `json:"jumaahId" minimum:"1"`
	DelayedETA *time.Time `json:"delayedETA"`
}
