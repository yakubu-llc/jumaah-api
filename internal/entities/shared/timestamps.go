package shared

import "time"

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" bun:",soft_delete"`
}