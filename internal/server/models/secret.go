package models

import (
	"time"

	"github.com/srg-bnd/keeper/internal/shared/models"
	"gorm.io/gorm"
)

type Secret struct {
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Type models.SecretType `gorm:"not null" json:"type"`

	Version  uint   `gorm:"integer;not null" json:"version"`
	Metadata string `gorm:"type:string" json:"metadata"`

	Data  []byte `gorm:"type:bytea;not null" json:"data"`
	Users []User `gorm:"many2many:user_secrets" json:"-"`
}

func (Secret) TableName() string {
	return "secrets"
}
