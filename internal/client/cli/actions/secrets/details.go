package secrets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	"github.com/srg-bnd/keeper/internal/client/models"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"
	"github.com/urfave/cli/v2"
	uCLI "github.com/urfave/cli/v2"
)

type DetailsAPI interface {
	GetSecret(models.Secret) (*models.Secret, error)
}

type Decryptor interface {
	Decrypt([]byte) ([]byte, error)
}

type DetailsAction struct {
	api       DetailsAPI
	decryptor Decryptor
}

func NewDetailsAction(api DetailsAPI, decryptor Decryptor) *DetailsAction {
	return &DetailsAction{
		api:       api,
		decryptor: decryptor,
	}
}

func (a *DetailsAction) Do(c *uCLI.Context) error {
	secretID, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		return cli.Exit("Некорректные данные", 1)
	}

	secret, err := a.api.GetSecret(models.Secret{ID: uint(secretID)})
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	decryptedData, err := a.decryptor.Decrypt(secret.Data)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	data, err := a.formatdData(secret.Type, decryptedData)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	helpers.SuccessBadge("Secret data:")
	fmt.Printf(
		"- ID: %d\n- Type: %s\n- Metadata: %s\n",
		secret.ID,
		secret.Type,
		secret.Metadata,
	)

	helpers.SecretData(string(data))

	return nil
}

// format secret data
func (a *DetailsAction) formatdData(secretType shared_models.SecretType, decryptedData []byte) ([]byte, error) {
	switch secretType {
	case shared_models.SecretTypeAuth:
		data := strings.Split(string(decryptedData), ":")
		return []byte(fmt.Sprintf("login=%s & password=%s", data[0], data[1])), nil
	case shared_models.SecretTypeText:
		return []byte(decryptedData), nil
	case shared_models.SecretTypeBin:
		return []byte(decryptedData), nil
	case shared_models.SecretTypeBank:
		data := strings.Split(string(decryptedData), ":")
		return []byte(fmt.Sprintf("№=%s & date=%s & ccv=%s", data[0], data[1], data[2])), nil
	default:
		return nil, errors.New("unknowTypeError")
	}
}
