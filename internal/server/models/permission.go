package models

import (
	"errors"
)

type PermissionAction uint
type PermissionOwnerType string

const (
	PermissionActionCreate PermissionAction = iota
	PermissionActionRead
	PermissionActionUpdate
	PermissionActionDelete
)

const (
	PermissionOwnerTypeSecret PermissionOwnerType = "Secret"
)

func (a PermissionAction) Validate() error {
	switch a {
	case PermissionActionCreate, PermissionActionRead, PermissionActionUpdate, PermissionActionDelete:
		return nil
	default:
		return errors.New("invalid permission action")
	}
}

type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedAt uint64 `gorm:"not null"`
	UpdatedAt uint64 `gorm:"not null"`

	OwnerType string           `gorm:"index:idx_permission_owner,priority:1;not null"`
	OwnerID   uint             `gorm:"index:idx_permission_owner,priority:2;not null"`
	Action    PermissionAction `gorm:"not null"`
	RoleID    uint             `gorm:"index:idx_permission_role;not null"`

	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Permission) TableName() string {
	return "permissions"
}
