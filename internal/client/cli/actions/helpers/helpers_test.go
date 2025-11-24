package helpers

import "testing"

func TestSuccessBadge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessBadge(tt.str)
		})
	}
}

func TestWarnBadge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WarnBadge(tt.str)
		})
	}
}

func TestErrorBadge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorBadge(tt.str)
		})
	}
}

func TestSecretData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SecretData(tt.str)
		})
	}
}
