package middleware_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/middleware"
)

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		secretKey string
		want      gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleware.AuthMiddleware(tt.secretKey)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("AuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
