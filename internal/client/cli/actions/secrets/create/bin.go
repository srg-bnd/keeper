package create

import (
	"fmt"
	"os"

	"github.com/srg-bnd/keeper/internal/client/api"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"

	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type BinAction struct {
	api *api.API
}

func NewBinAction(api *api.API) *BinAction {
	return &BinAction{
		api: api,
	}
}

func (a *BinAction) Do(c *uCLI.Context) error {
	filePath := c.Args().Get(0)
	if filePath == "" {
		return cli.Exit("Ошибка: путь к файлу не указан", 1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return cli.Exit(fmt.Sprintf("Ошибка чтения файла: %v", err), 1)
	}

	return CreateSecret(a.api, data, shared_models.SecretTypeBin)
}
