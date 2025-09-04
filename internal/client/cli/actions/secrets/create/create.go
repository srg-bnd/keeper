package create

import (
	"strconv"

	"github.com/srg-bnd/keeper/internal/client/api"
	"github.com/srg-bnd/keeper/internal/client/cli/actions/helpers"
	"github.com/srg-bnd/keeper/internal/client/models"
	shared_models "github.com/srg-bnd/keeper/internal/shared/models"
)

func CreateSecret(api *api.API, data []byte, secretType shared_models.SecretType) error {
	newSecret := models.Secret{Metadata: string(secretType), Type: secretType}
	encryptedData, err := api.Cipher.Encrypt(data)
	if err != nil {
		helpers.ErrorBadge(err.Error())
		return nil
	}

	newSecret.Data = encryptedData
	secret, err := api.CreateSecret(newSecret)
	if err != nil {
		helpers.ErrorBadge(err.Error())
		return nil
	}

	helpers.SuccessBadge("Secret created, ID=" + strconv.Itoa(int(secret.ID)))
	return nil
}
