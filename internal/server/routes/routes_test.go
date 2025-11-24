package routes_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server/routes"
)

func TestSetRoutes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		router *gin.Engine
		config *config.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routes.SetRoutes(tt.router, tt.config)
		})
	}
}
