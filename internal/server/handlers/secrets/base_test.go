package secrets

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
)

func TestSecretHandler_getUser(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		c       *gin.Context
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var h SecretHandler
			got, gotErr := h.getUser(tt.c)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getUser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getUser() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("getUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
