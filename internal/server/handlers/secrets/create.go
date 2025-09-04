package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/policy"
)

func (h *SecretHandler) CreateHandler(c *gin.Context) {
	user, error := h.getUser(c)
	if error != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	action := models.PermissionActionCreate
	ownerType := models.PermissionOwnerTypeSecret
	if !h.Policy.HasPermission(&policy.PolicyData{User: user, Action: &action, OwnerType: ownerType}) {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	var secret models.Secret

	// Checks secret data
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// Saves to DB
	if err := h.Storage.SecretCreate(c.Request.Context(), &secret); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// Creates association
	if err := h.DB.Model(&user).Association("Secrets").Append(&secret); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// Returns secret
	c.JSON(http.StatusCreated, secret)
}
