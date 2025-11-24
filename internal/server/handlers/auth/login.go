package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/models"
)

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Checks input data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var user models.User
	user.Email = input.Email
	// Checks user
	if err := h.Storage.UserGet(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusConflict, nil)
		return
	}

	if !h.Security.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	token, err := h.Security.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// Returns user with token
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	})
}
