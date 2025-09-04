// Auth API
package api

import (
	"encoding/json"
	"errors"

	"github.com/srg-bnd/keeper/internal/client/models"
)

func (c *API) Login(user models.User) (*models.User, error) {
	request := c.httpClient.R().
		SetBody(&user)
		// SetResult(&user)

	response, err := request.Post("/api/login")
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	var newUser models.User
	err = json.Unmarshal(response.Body(), &newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (c *API) Register(user models.User) (*models.User, error) {
	request := c.httpClient.R().SetBody(&user)

	response, err := request.Post("/api/register")
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(response.Status())
	}

	var newUser models.User
	err = json.Unmarshal(response.Body(), &newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
