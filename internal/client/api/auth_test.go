package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/srg-bnd/keeper/config/client"
	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/stretchr/testify/assert"
)

func getAPI(url string) *API {
	cfg := &config.Config{
		Setting: &config.Setting{
			AuthToken:  "test-token",
			ServerHost: "http://" + url,
		},
	}

	return NewAPI(cfg)
}

func startMockServer(fakeDate []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if string(fakeDate) == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write(fakeDate)
			w.WriteHeader(http.StatusOK)
		}
	}))
}

func TestAPILogin_Success(t *testing.T) {
	server := startMockServer([]byte(`
		{
			"id": 1,
			"username": "srg",
			"email": "srg@yandex.ru",
			"token": "mock-token"
		}
	`))
	defer server.Close()

	api := getAPI(server.URL[7:])
	user := models.User{Email: "srg@yandex.ru", Password: "password123"}
	result, err := api.Login(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.AuthToken)
	assert.Equal(t, result.Password, "")
	assert.Equal(t, user.Email, result.Email)
}

func TestAPIRegister_Success(t *testing.T) {
	server := startMockServer([]byte(`
		{
			"id": 1,
			"username": "srg",
			"email": "srg@yandex.ru"
		}
	`))
	defer server.Close()

	api := getAPI(server.URL[7:])
	user := models.User{Email: "srg@yandex.ru", Password: "password123"}
	result, err := api.Register(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, result.AuthToken, "")
	assert.Equal(t, result.Password, "")
}
