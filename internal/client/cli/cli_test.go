package cli

import (
	"context"
	"testing"

	config "github.com/srg-bnd/keeper/config/client"
)

func TestCLI_Start(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		config  *config.Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := NewCLI(tt.config)
			gotErr := cli.Start(context.Background())
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Start() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Start() succeeded unexpectedly")
			}
		})
	}
}

func TestNewCLI(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		config *config.Config
		want   *CLI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCLI(tt.config)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewCLI() = %v, want %v", got, tt.want)
			}
		})
	}
}
