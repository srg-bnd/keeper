package secrets

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/policy"
	"github.com/srg-bnd/keeper/internal/server/storage"
	"gorm.io/gorm"
)

type SecretHandler struct {
	Policy  *policy.Policy
	Storage storage.SecretRepositories
	DB      *gorm.DB
}

func (h *SecretHandler) getUser(c *gin.Context) (*models.User, error) {
	userIDInt, exists := c.Get("user_id")
	if !exists {
		return nil, errors.New("not exists")
	}

	userID, ok := userIDInt.(uint)
	if !ok {
		return nil, errors.New("not valid")
	}

	user := models.User{}
	if err := h.DB.Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
