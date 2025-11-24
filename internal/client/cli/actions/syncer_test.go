package actions

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/client/syncer"
	uCLI "github.com/urfave/cli/v2"
)

func TestNewSyncerAction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		syncer *syncer.Syncer
		want   *SyncerAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSyncerAction(tt.syncer)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewSyncerAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncerAction_Do(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		syncer *syncer.Syncer
		// Named input parameters for target function.
		c       *uCLI.Context
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewSyncerAction(tt.syncer)
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
