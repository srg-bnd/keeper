package secrets

import (
	"context"
	"os"
	"testing"

	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/urfave/cli/v2"
)

type MockListAPI struct{}

func (a *MockListAPI) AllSecrets() ([]models.Secret, error) {
	return []models.Secret{}, nil
}

func TestListAction_Do(t *testing.T) {
	action := NewListAction(&MockListAPI{})

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "list",
				Action: action.Do,
			},
		},
	}
	args := []string{
		os.Args[0],
		"list",
	}

	err := app.RunContext(context.Background(), args)

	if err != nil {
		t.Errorf("Do вернул ошибку: %v", err)
	}
}
