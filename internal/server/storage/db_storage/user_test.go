package db_storage_test

import (
	"context"
	"testing"

	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/storage/db_storage"
	"gorm.io/gorm"
)

func TestDBStorage_UserGet(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := db_storage.NewDBStorage(tt.db)
			gotErr := s.UserGet(context.Background(), tt.user)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UserGet() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UserGet() succeeded unexpectedly")
			}
		})
	}
}

func TestDBStorage_UserCreate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := db_storage.NewDBStorage(tt.db)
			gotErr := s.UserCreate(context.Background(), tt.user)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UserCreate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UserCreate() succeeded unexpectedly")
			}
		})
	}
}
