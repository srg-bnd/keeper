package auth

import (
	"context"
	"os"
	"testing"

	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/urfave/cli/v2"
)

type MockRegisterAPI struct{}

func (a *MockRegisterAPI) Register(user models.User) (*models.User, error) {
	return &models.User{
		Email:     user.Email,
		AuthToken: "test-token",
	}, nil
}

func TestRegisterAction_Do(t *testing.T) {
	action := NewRegisterAction(&MockRegisterAPI{})

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "register",
				Action: action.Do,
			},
		},
	}
	args := []string{
		os.Args[0],
		"register",
		"user@example.com",
		"password123",
	}

	err := app.RunContext(context.Background(), args)

	if err != nil {
		t.Errorf("Do вернул ошибку: %v", err)
	}
}
