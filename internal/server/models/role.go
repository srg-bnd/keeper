package models

type RoleName string

const (
	RoleNameAdmin  RoleName = "admin"
	RoleNameClient RoleName = "client"
)

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedAt uint64 `gorm:"not null"`
	UpdatedAt uint64 `gorm:"not null"`

	Name        RoleName     `gorm:"uniqueIndex:idx_role_name;not null"`
	Permissions []Permission `gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string {
	return "roles"
}
