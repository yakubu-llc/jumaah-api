package account

import (
	"github.com/uptrace/bun"
	"github.com/yakubu-llc/jumaah-api/internal/entities/shared"
)

// Account represents an account entity.
type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID   int    `json:"id"`
	Name string `json:"name"`
	shared.Timestamps
}

// CreateAccountParams contains the parameters for creating a new account.
type CreateAccountParams struct {
	Name string `json:"name"`
}

// UpdateAccountParams contains the parameters for updating a account.
type UpdateAccountParams struct {
	Name string `json:"name"`
}
