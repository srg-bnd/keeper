package routes

import (
	"github.com/gin-gonic/gin"
	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server/handlers"
	"github.com/srg-bnd/keeper/internal/server/middleware"
)

func SetRoutes(router *gin.Engine, config *config.Config) {
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(config.SecretKey))

	authHandler := handlers.NewAuthHandler(config)
	secretHandler := handlers.NewSecretHandler(config)

	registerHealthCheckRoutes(router)

	// Auth
	router.POST("api/register", authHandler.RegisterHandler)
	router.POST("api/login", authHandler.LoginHandler)

	// Secrets
	authorized.GET("api/secrets", secretHandler.ListHandler)
	authorized.GET("api/secrets/:id", secretHandler.ShowHandler)
	authorized.POST("api/secrets", secretHandler.CreateHandler)
}
