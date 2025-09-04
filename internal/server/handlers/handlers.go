package handlers

import (
	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server/handlers/auth"
	"github.com/srg-bnd/keeper/internal/server/handlers/secrets"
)

func NewSecretHandler(config *config.Config) *secrets.SecretHandler {
	return &secrets.SecretHandler{
		Policy:  config.Policy,
		Storage: config.Storage,
		DB:      config.DB,
	}
}

func NewAuthHandler(config *config.Config) *auth.AuthHandler {
	return &auth.AuthHandler{
		Policy:   config.Policy,
		Storage:  config.Storage,
		Security: config.Security,
	}
}
