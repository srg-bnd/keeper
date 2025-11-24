package db_storage_test

import (
	"context"
	"testing"

	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/storage/db_storage"
	"gorm.io/gorm"
)

func TestDBStorage_SecretGetAll(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		secret  models.Secret
		want    []models.Secret
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := db_storage.NewDBStorage(tt.db)
			got, gotErr := s.SecretGetAll(context.Background(), tt.secret)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SecretGetAll() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SecretGetAll() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SecretGetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
