// Package server provides implementation of HTTP server for handling incoming requests.
//
// The package contains core components for running a web server:
// - Server structure representing HTTP server instance
// - Methods for server initialization and management
// - Routing and request handling capabilities
//
// This package serves as a central component for handling client requests
// and managing server operations
package server

import (
	"github.com/gin-gonic/gin"
	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server/routes"
)

// Server represents an HTTP server instance with routing capabilities.
//
// The Server struct is responsible for:
// - Handling incoming HTTP requests
// - Routing requests to appropriate handlers
// - Managing server lifecycle
//
// Fields:
// router - HTTP handler responsible for routing incoming requests
type Server struct {
	router *gin.Engine
	config *config.Config
}

// NewServer creates a new instance of the HTTP server.
//
// Parameters:
// - router: HTTP handler responsible for routing incoming requests
//
// Returns a pointer to the newly created Server instance.
func NewServer(config *config.Config) *Server {
	return &Server{
		router: config.Router,
		config: config,
	}
}

// Start launches the HTTP server (Gin) on the specified address.
//
// Parameters:
// - addr: network address to listen on (e.g., ":8080")
//
// Returns:
// - error: nil if server started successfully, error otherwise
//
// This method initializes and starts the HTTP server,
// binding it to the specified address and using the configured router.
func (server *Server) Start(addr string) error {
	routes.SetRoutes(server.router, server.config)

	if err := server.router.Run(addr); err != nil {
		return err
	}

	return nil
}
