package db_storage

import (
	"context"

	"github.com/srg-bnd/keeper/internal/server/models"
)

func (s *DBStorage) UserGet(ctx context.Context, user *models.User) error {
	if err := s.db.First(user, "email = ?", user.Email).Error; err != nil {
		return err
	}

	return nil
}

func (s *DBStorage) UserCreate(ctx context.Context, user *models.User) error {
	var clientRole models.Role

	if err := s.db.First(&clientRole, "name = ?", models.RoleNameClient).Error; err != nil {
		return err
	}

	user.Roles = []models.Role{clientRole}
	if result := s.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
