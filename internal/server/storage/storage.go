package storage

import (
	"context"

	"gorm.io/gorm"

	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/storage/db_storage"
)

type SecretRepositories interface {
	SecretGetAll(context.Context, models.Secret) ([]models.Secret, error)
	SecretGet(context.Context, *models.Secret) error
	SecretCreate(context.Context, *models.Secret) error
}

type UserRepositories interface {
	UserGet(context.Context, *models.User) error
	UserCreate(context.Context, *models.User) error
}

type Repositories interface {
	SecretRepositories
	UserRepositories
}

func NewStorage(db *gorm.DB) Repositories {
	return db_storage.NewDBStorage(db)
}
