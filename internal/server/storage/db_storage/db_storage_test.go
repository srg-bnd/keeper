package db_storage_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/server/storage/db_storage"
	"gorm.io/gorm"
)

func TestNewDBStorage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db   *gorm.DB
		want *db_storage.DBStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := db_storage.NewDBStorage(tt.db)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewDBStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
