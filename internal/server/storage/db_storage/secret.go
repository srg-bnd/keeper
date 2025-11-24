package db_storage

import (
	"context"

	"github.com/srg-bnd/keeper/internal/server/models"
)

func (s *DBStorage) SecretGetAll(ctx context.Context, secret models.Secret) ([]models.Secret, error) {
	var secrets []models.Secret
	err := s.db.Model(&models.Secret{}).Find(&secrets).Error

	return secrets, err
}

func (s *DBStorage) SecretGet(ctx context.Context, secret *models.Secret) error {
	if err := s.db.First(secret, secret.ID).Error; err != nil {
		return err
	}

	return nil
}

func (s *DBStorage) SecretCreate(ctx context.Context, secret *models.Secret) error {
	if result := s.db.Create(&secret); result.Error != nil {
		return result.Error
	}

	return nil
}
