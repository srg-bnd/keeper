package dbtool

import (
	"github.com/srg-bnd/keeper/internal/server/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Secret{},
		&models.Role{},
		&models.Permission{},
	); err != nil {
		return err
	}

	return nil
}
