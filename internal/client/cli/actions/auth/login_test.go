package auth

import (
	"context"
	"os"
	"testing"

	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/urfave/cli/v2"
)

type MockLoginAPI struct{}

func (a *MockLoginAPI) Login(user models.User) (*models.User, error) {
	return &models.User{
		Email:     user.Email,
		AuthToken: "test-token",
	}, nil
}

func TestLoginAction_Do_Basic(t *testing.T) {
	action := NewLoginAction(&MockLoginAPI{})

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "login",
				Action: action.Do,
			},
		},
	}
	args := []string{
		os.Args[0],
		"login",
		"user@example.com",
		"password123",
	}

	err := app.RunContext(context.Background(), args)

	if err != nil {
		t.Errorf("Do вернул ошибку: %v", err)
	}
}
