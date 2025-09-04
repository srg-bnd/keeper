package create

import (
	"github.com/srg-bnd/keeper/internal/client/api"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"

	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type CardAction struct {
	api *api.API
}

func NewCardAction(api *api.API) *CardAction {
	return &CardAction{
		api: api,
	}
}

func (a *CardAction) Do(c *uCLI.Context) error {
	number := c.Args().Get(0)
	date := c.Args().Get(1)
	cvv := c.Args().Get(2)
	if number == "" || date == "" || cvv == "" {
		return cli.Exit("Некорректные данные", 1)
	}

	data := []byte(number + ":" + date + ":" + cvv)
	return CreateSecret(a.api, data, shared_models.SecretTypeBank)
}
