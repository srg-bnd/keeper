package auth

import (
	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	"github.com/srg-bnd/keeper/internal/client/models"

	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type RegisterAPI interface {
	Register(models.User) (*models.User, error)
}

type RegisterAction struct {
	api RegisterAPI
}

func NewRegisterAction(api RegisterAPI) *RegisterAction {
	return &RegisterAction{
		api: api,
	}
}

func (a *RegisterAction) Do(c *uCLI.Context) error {
	login := c.Args().Get(0)
	password := c.Args().Get(1)
	if login == "" || password == "" {
		return cli.Exit("Некорректные данные", 1)
	}

	user, err := a.api.Register(models.User{Email: login, Username: login, Password: password})

	if err != nil {
		helpers.ErrorBadge(err.Error())
		return nil
	} else {
		helpers.SuccessBadge("email: " + user.Email)
	}

	return nil
}
