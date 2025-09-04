package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
)

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var user models.User

	// Checks user data
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.Username = strings.TrimSpace(user.Username)

	// Checks user
	if err := h.Storage.UserGet(c.Request.Context(), &user); err == nil {
		c.JSON(http.StatusConflict, nil)
		return
	}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// Hashes password
	hashedPassword, err := h.Security.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	user.Password = hashedPassword

	// Saves to DB
	if err := h.Storage.UserCreate(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// Returns user
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
