package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/policy"
)

func (h *SecretHandler) ListHandler(c *gin.Context) {
	// Gets user
	user, error := h.getUser(c)
	if error != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// Checks permissions
	action := models.PermissionActionRead
	ownerType := models.PermissionOwnerTypeSecret
	if !h.Policy.HasPermission(&policy.PolicyData{User: user, Action: &action, OwnerType: ownerType}) {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	// Loads secrets
	var userWithSecrets models.User
	if err := h.DB.Preload("Secrets").First(&userWithSecrets, user.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// Returns secrets
	c.JSON(http.StatusCreated, userWithSecrets.Secrets)
}
