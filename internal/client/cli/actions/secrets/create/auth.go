package create

import (
	"github.com/srg-bnd/keeper/internal/client/api"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"

	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type AuthAction struct {
	api *api.API
}

func NewAuthAction(api *api.API) *AuthAction {
	return &AuthAction{
		api: api,
	}
}

func (a *AuthAction) Do(c *uCLI.Context) error {
	login := c.Args().Get(0)
	password := c.Args().Get(1)
	if login == "" || password == "" {
		return cli.Exit("Некорректные данные", 1)
	}

	data := []byte(login + ":" + password)
	return CreateSecret(a.api, data, shared_models.SecretTypeAuth)
}
