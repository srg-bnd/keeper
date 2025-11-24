// Secret API
package api

import (
	"errors"
	"strconv"

	"github.com/srg-bnd/keeper/internal/client/models"
)

func (c *API) AllSecrets() ([]models.Secret, error) {
	var secrets []models.Secret
	request := c.httpClient.R().
		SetHeader("Authorization", c.authToken).
		SetResult(&secrets)

	response, err := request.Get("/api/secrets")
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	return secrets, nil
}

func (c *API) CreateSecret(secret models.Secret) (*models.Secret, error) {
	request := c.httpClient.R().
		SetHeader("Authorization", c.authToken).
		SetBody(&secret).SetResult(&secret)

	response, err := request.Post("/api/secrets")

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	return &secret, nil
}

func (c *API) GetSecret(secret models.Secret) (*models.Secret, error) {
	request := c.httpClient.R().
		SetHeader("Authorization", c.authToken).
		SetPathParams(map[string]string{"secretId": strconv.Itoa(int(secret.ID))}).
		SetResult(&secret)

	response, err := request.Get("/api/secrets/{secretId}")

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	return &secret, nil
}
