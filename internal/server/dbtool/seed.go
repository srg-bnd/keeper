package dbtool

import (
	"github.com/srg-bnd/keeper/internal/server/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	admin := models.Role{Name: models.RoleNameAdmin}
	client := models.Role{Name: models.RoleNameClient}
	db.Create(&admin)
	db.Create(&client)

	db.Create(&models.Permission{RoleID: client.ID, Action: models.PermissionAction(models.PermissionActionCreate), OwnerType: string(models.PermissionOwnerTypeSecret)})
	db.Create(&models.Permission{RoleID: client.ID, Action: models.PermissionAction(models.PermissionActionRead), OwnerType: string(models.PermissionOwnerTypeSecret)})

	return nil
}
