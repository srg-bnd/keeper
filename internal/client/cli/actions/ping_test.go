package actions

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/client/api"
	uCLI "github.com/urfave/cli/v2"
)

func TestNewPingAction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		api  *api.API
		want *PingAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPingAction(tt.api)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewPingAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPingAction_Do(t *testing.T) {
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
			a := NewPingAction(tt.api)
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
