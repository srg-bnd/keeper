package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_registerHealthCheckRoutes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		router *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerHealthCheckRoutes(tt.router)
		})
	}
}
