package actions

import (
	"github.com/srg-bnd/keeper/internal/client/api"
	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	uCLI "github.com/urfave/cli/v2"
)

type PingAction struct {
	api *api.API
}

func NewPingAction(api *api.API) *PingAction {
	return &PingAction{
		api: api,
	}
}

func (a *PingAction) Do(c *uCLI.Context) error {
	data, err := a.api.Ping()

	if err != nil {
		helpers.ErrorBadge(err.Error())
	} else {
		helpers.SuccessBadge(string(data))
	}

	return nil
}
