package db_storage

import (
	"gorm.io/gorm"
)

type DBStorage struct {
	db *gorm.DB
}

func NewDBStorage(db *gorm.DB) *DBStorage {
	return &DBStorage{
		db: db,
	}
}
