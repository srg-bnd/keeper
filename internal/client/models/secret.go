// Secret model
package models

import (
	"time"

	"github.com/srg-bnd/keeper/internal/shared/models"
)

const unknowTypeError = "unknown type"

type Secret struct {
	ID        uint              `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Type      models.SecretType `json:"type"`
	Version   uint              `json:"version"`
	Metadata  string            `json:"metadata"`
	Data      []byte            `json:"data"`
}
