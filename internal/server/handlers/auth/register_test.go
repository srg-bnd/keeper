package auth

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthHandler_RegisterHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		c *gin.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var h AuthHandler
			h.RegisterHandler(tt.c)
		})
	}
}
