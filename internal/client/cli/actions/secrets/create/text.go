package create

import (
	"github.com/srg-bnd/keeper/internal/client/api"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"

	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type TextAction struct {
	api *api.API
}

func NewTextAction(api *api.API) *TextAction {
	return &TextAction{
		api: api,
	}
}

func (a *TextAction) Do(c *uCLI.Context) error {
	text := c.Args().Get(0)
	if text == "" {
		return cli.Exit("Некорректные данные", 1)
	}

	data := []byte(text)
	return CreateSecret(a.api, data, shared_models.SecretTypeText)
}
