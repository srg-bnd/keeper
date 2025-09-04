// Client dependencies
package config

import (
	"context"

	"github.com/srg-bnd/keeper/internal/client/models"
)

type SecretStorage interface {
	AddList(ctx context.Context, secret []models.Secret) error
	Add(ctx context.Context, secret models.Secret) error
	Delete(ctx context.Context, secretID uint) error
	Get(ctx context.Context, secretID uint) (*models.Secret, error)
	List(ctx context.Context) (map[uint]models.Secret, error)
}

type Cipher interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type Config struct {
	Cipher        Cipher
	Setting       *Setting
	SecretStorage SecretStorage
	WithSync      bool
}
