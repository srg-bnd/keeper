package create

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/client/api"
	uCLI "github.com/urfave/cli/v2"
)

func TestNewCardAction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		api  *api.API
		want *CardAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCardAction(tt.api)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewCardAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardAction_Do(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		api *api.API
		// Named input parameters for target function.
		c       *uCLI.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewCardAction(tt.api)
			gotErr := a.Do(tt.c)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Do() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Do() succeeded unexpectedly")
			}
		})
	}
}
