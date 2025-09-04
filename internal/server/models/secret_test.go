package models_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/server/models"
)

func TestSecret_TableName(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s models.Secret
			got := s.TableName()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
