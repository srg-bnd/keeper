// CLI
package cli

import (
	"context"
	"os"

	config "github.com/srg-bnd/keeper/config/client"
	keeperAPI "github.com/srg-bnd/keeper/internal/client/api"
	"github.com/srg-bnd/keeper/internal/client/cli/actions"
	auth_actions "github.com/srg-bnd/keeper/internal/client/cli/actions/auth"
	secret_actions "github.com/srg-bnd/keeper/internal/client/cli/actions/secrets"
	create_secret_actions "github.com/srg-bnd/keeper/internal/client/cli/actions/secrets/create"
	"github.com/srg-bnd/keeper/internal/client/syncer"

	uCLI "github.com/urfave/cli/v2"
)

// Type for CLI
type CLI struct {
	api     *keeperAPI.API
	storage *config.SecretStorage
	syncer  *syncer.Syncer
}

// Returns new CLI
func NewCLI(config *config.Config) *CLI {
	keeperAPI := keeperAPI.NewAPI(config)

	return &CLI{
		api:    keeperAPI,
		syncer: syncer.NewSyncer(keeperAPI, config),
	}
}

// Starts CLI
func (cli *CLI) Start(ctx context.Context) error {
	app := &uCLI.App{
		Name:  "Keeper",
		Usage: "Клиент-серверная система, позволяющая надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.",
		Commands: []*uCLI.Command{
			{
				Name:   "ping",
				Usage:  "Проверить сервер",
				Action: actions.NewPingAction(cli.api).Do,
			},
			{
				Name:   "syncer",
				Usage:  "Запустить синхронизацию",
				Action: actions.NewSyncerAction(cli.syncer).Do,
			},
			// Auth
			{
				Name:   "register",
				Usage:  "Зарегистрироваться",
				Action: auth_actions.NewRegisterAction(cli.api).Do,
			},
			{
				Name:   "login",
				Usage:  "Логин в хранилище",
				Action: auth_actions.NewLoginAction(cli.api).Do,
			},
			// Gets secrets
			{
				Name:   "list",
				Usage:  "Показать список секретов",
				Action: secret_actions.NewListAction(cli.api).Do,
			},
			{
				Name:   "details",
				Usage:  "Детализация секрета",
				Action: secret_actions.NewDetailsAction(cli.api, cli.api.Cipher).Do,
			},
			// Creates secrets
			{
				Name:   "secret auth",
				Usage:  "Сохранить логин/пароль",
				Action: create_secret_actions.NewAuthAction(cli.api).Do,
			},
			{
				Name:   "secret bin",
				Usage:  "Сохранить бинарник",
				Action: create_secret_actions.NewBinAction(cli.api).Do,
			},
			{
				Name:   "secret card",
				Usage:  "Сохранить банковскую карту",
				Action: create_secret_actions.NewCardAction(cli.api).Do,
			},
			{
				Name:   "secret text",
				Usage:  "Сохранить текст",
				Action: create_secret_actions.NewTextAction(cli.api).Do,
			},
		},
	}

	return app.Run(os.Args)
}
