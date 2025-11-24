package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	config "github.com/srg-bnd/keeper/config/client"
	"github.com/stretchr/testify/assert"
)

func TestAPIPing_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}))
	defer server.Close()

	hostPort := server.URL[7:] // "127.0.0.1:8080"

	cfg := &config.Config{
		Setting: &config.Setting{
			AuthToken:  "test-token",
			ServerHost: hostPort,
		},
		Cipher:        nil,
		SecretStorage: nil,
	}

	api := NewAPI(cfg)
	api.httpClient = resty.New().SetBaseURL(server.URL)
	response, err := api.Ping()

	assert.NoError(t, err)
	assert.Equal(t, []byte("pong"), response)
}
