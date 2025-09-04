package storage_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/server/storage"
	"gorm.io/gorm"
)

func TestNewStorage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db   *gorm.DB
		want storage.Repositories
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := storage.NewStorage(tt.db)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
