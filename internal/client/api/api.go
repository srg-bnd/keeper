// API
package api

import (
	"errors"

	"github.com/go-resty/resty/v2"
	config "github.com/srg-bnd/keeper/config/client"
)

type API struct {
	authToken     string
	Cipher        config.Cipher
	httpClient    *resty.Client
	SecretStorage config.SecretStorage
	URL           string
}

func NewAPI(config *config.Config) *API {
	return &API{
		authToken:     config.Setting.AuthToken,
		Cipher:        config.Cipher,
		httpClient:    newHTTPClient(config),
		SecretStorage: config.SecretStorage,
		URL:           config.Setting.ServerHost,
	}
}

func newHTTPClient(config *config.Config) *resty.Client {
	return resty.New().
		SetBaseURL(config.Setting.ServerHost).
		SetHeader("Content-Type", "application/json")
}

func apiResponse(response *resty.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	return response.Body(), nil
}

func (api *API) Ping() ([]byte, error) {
	request := api.httpClient.R()
	return apiResponse(request.Get("/ping"))
}
