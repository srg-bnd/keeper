package server_test

import (
	"testing"

	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		config *config.Config
		want   *server.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := server.NewServer(tt.config)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		config *config.Config
		// Named input parameters for target function.
		addr    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := server.NewServer(tt.config)
			gotErr := server.Start(tt.addr)
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
