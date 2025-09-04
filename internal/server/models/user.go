package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`

	Secrets []Secret `gorm:"many2many:user_secrets"`
	Roles   []Role   `gorm:"many2many:user_roles"`
}

func (User) TableName() string {
	return "users"
}
