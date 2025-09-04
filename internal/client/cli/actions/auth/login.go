package auth

import (
	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type LoginAPI interface {
	Login(models.User) (*models.User, error)
}

type LoginAction struct {
	api LoginAPI
}

func NewLoginAction(api LoginAPI) *LoginAction {
	return &LoginAction{
		api: api,
	}
}

func (a *LoginAction) Do(c *uCLI.Context) error {
	email := c.Args().Get(0)
	password := c.Args().Get(1)
	if email == "" || password == "" {
		return cli.Exit("Некорректные данные", 1)
	}

	user, err := a.api.Login(models.User{Email: email, Password: password})

	if err != nil {
		helpers.ErrorBadge(err.Error())
		return nil
	} else {
		helpers.SuccessBadge("AUTH_TOKEN=" + user.AuthToken)
	}

	return nil
}
