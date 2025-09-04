package secrets

import (
	"context"
	"os"
	"testing"

	"github.com/srg-bnd/keeper/internal/client/models"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"
	"github.com/urfave/cli/v2"
)

type MockDetailsAPI struct{}
type MockDecryptor struct{}

func (a *MockDetailsAPI) GetSecret(secret models.Secret) (*models.Secret, error) {
	secret.Type = shared_models.SecretTypeText
	return &secret, nil
}

func (a *MockDecryptor) Decrypt(data []byte) ([]byte, error) {
	return data, nil
}

func TestLoginAction(t *testing.T) {
	action := NewDetailsAction(&MockDetailsAPI{}, &MockDecryptor{})

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "details",
				Action: action.Do,
			},
		},
	}
	args := []string{
		os.Args[0],
		"details",
		"1",
	}

	err := app.RunContext(context.Background(), args)

	if err != nil {
		t.Errorf("Do вернул ошибку: %v", err)
	}
}
