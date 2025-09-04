package security_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/server/services/security"
)

func TestNewSecurity(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secretKey string
		want      *security.Security
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := security.NewSecurity(tt.secretKey)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurity_CheckPasswordHash(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		secretKey string
		// Named input parameters for target function.
		password string
		hash     string
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := security.NewSecurity(tt.secretKey)
			got := s.CheckPasswordHash(tt.password, tt.hash)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurity_HashPassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		secretKey string
		// Named input parameters for target function.
		password string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := security.NewSecurity(tt.secretKey)
			got, gotErr := s.HashPassword(tt.password)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("HashPassword() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("HashPassword() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurity_GenerateToken(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		secretKey string
		// Named input parameters for target function.
		userID  uint
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := security.NewSecurity(tt.secretKey)
			got, gotErr := s.GenerateToken(tt.userID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GenerateToken() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GenerateToken() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
