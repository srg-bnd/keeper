package secrets

import (
	"fmt"

	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	"github.com/srg-bnd/keeper/internal/client/models"

	uCLI "github.com/urfave/cli/v2"
)

type ListAPI interface {
	AllSecrets() ([]models.Secret, error)
}

type ListAction struct {
	api ListAPI
}

func NewListAction(api ListAPI) *ListAction {
	return &ListAction{
		api: api,
	}
}

func (a *ListAction) Do(ctx *uCLI.Context) error {
	list, _ := a.api.AllSecrets()
	if len(list) == 0 {
		helpers.WarnBadge("Empty")
	} else {
		helpers.SuccessBadge("List:")
		for id, secret := range list {
			fmt.Println("-", id, secret.Metadata, secret.Type)
		}
	}

	return nil
}
