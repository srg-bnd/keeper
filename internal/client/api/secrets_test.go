package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	config "github.com/srg-bnd/keeper/config/client"
	"github.com/srg-bnd/keeper/internal/client/models"
	"github.com/stretchr/testify/assert"
)

func TestAPIAllSecrets_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &config.Config{
		Setting: &config.Setting{
			AuthToken:  "test-token",
			ServerHost: server.URL[7:],
		},
		Cipher:        nil,
		SecretStorage: nil,
	}

	api := NewAPI(cfg)
	api.httpClient = resty.New().SetBaseURL(server.URL)
	_, err := api.AllSecrets()

	assert.NoError(t, err)
}

func TestGetSecret_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &config.Config{
		Setting: &config.Setting{
			AuthToken:  "test-token",
			ServerHost: server.URL[7:],
		},
		Cipher:        nil,
		SecretStorage: nil,
	}

	api := NewAPI(cfg)
	api.httpClient = resty.New().SetBaseURL(server.URL)
	_, err := api.GetSecret(models.Secret{})

	assert.NoError(t, err)
}

func TestCreateSecret_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &config.Config{
		Setting: &config.Setting{
			AuthToken:  "test-token",
			ServerHost: server.URL[7:],
		},
		Cipher:        nil,
		SecretStorage: nil,
	}

	api := NewAPI(cfg)
	api.httpClient = resty.New().SetBaseURL(server.URL)
	_, err := api.CreateSecret(models.Secret{})

	assert.NoError(t, err)
}
