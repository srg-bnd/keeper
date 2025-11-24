package storage_test

import (
	"context"
	"testing"

	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/srg-bnd/keeper/internal/client/storage"
)

func TestStorage_Add(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secret  models.Secret
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage.NewStorage()
			gotErr := s.Add(context.Background(), tt.secret)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Add() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Add() succeeded unexpectedly")
			}
		})
	}
}

func TestStorage_Delete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secretID uint
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage.NewStorage()
			gotErr := s.Delete(context.Background(), tt.secretID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Delete() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Delete() succeeded unexpectedly")
			}
		})
	}
}

func TestStorage_Get(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secretID uint
		want     *models.Secret
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage.NewStorage()
			got, gotErr := s.Get(context.Background(), tt.secretID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Get() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Get() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
