package create_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/client/api"
	"github.com/srg-bnd/keeper/internal/client/cli/actions/secrets/create"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"
)

func TestCreateSecret(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		api        *api.API
		data       []byte
		secretType shared_models.SecretType
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := create.CreateSecret(tt.api, tt.data, tt.secretType)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateSecret() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateSecret() succeeded unexpectedly")
			}
		})
	}
}
