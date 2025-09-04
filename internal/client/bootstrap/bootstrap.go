package bootstrap

import (
	"context"

	config "github.com/srg-bnd/keeper/config/client"
	"github.com/srg-bnd/keeper/internal/client/cli"
	"github.com/srg-bnd/keeper/internal/client/storage"
	"github.com/srg-bnd/keeper/internal/shared/cipher"
)

// Run initializes and starts the application.
// Returns an error if configuration fails or CLI cannot start.
func Run(ctx context.Context) error {
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	config := config.Config{
		Cipher:        cipher.NewCipher([]byte(setting.EncryptionKey)),
		Setting:       setting,
		SecretStorage: storage.NewStorage(),
	}

	return cli.NewCLI(&config).Start(ctx)
}
