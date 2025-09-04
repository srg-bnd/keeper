package secrets

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
	"github.com/srg-bnd/keeper/internal/server/policy"
)

func (h *SecretHandler) ShowHandler(c *gin.Context) {
	user, error := h.getUser(c)
	if error != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	action := models.PermissionActionRead
	ownerType := models.PermissionOwnerTypeSecret
	if !h.Policy.HasPermission(&policy.PolicyData{User: user, Action: &action, OwnerType: ownerType}) {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	secret, err := h.initSecret(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// Finds secret
	if err := h.Storage.SecretGet(c.Request.Context(), secret); err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	// Returns secret
	c.JSON(http.StatusCreated, secret)
}

func (h *SecretHandler) initSecret(c *gin.Context) (*models.Secret, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	secret := models.Secret{}
	secret.ID = uint(id)

	return &secret, nil
}
