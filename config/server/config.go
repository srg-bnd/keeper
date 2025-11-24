// Server dependencies
package config

import (
	"github.com/gin-gonic/gin"
	"github.com/srg-bnd/keeper/internal/server/policy"
	"github.com/srg-bnd/keeper/internal/server/services/security"
	"github.com/srg-bnd/keeper/internal/server/storage"
	"gorm.io/gorm"
)

type Config struct {
	DB        *gorm.DB
	Policy    *policy.Policy
	Router    *gin.Engine
	SecretKey string
	Setting   *Setting
	Storage   storage.Repositories
	Security  *security.Security
}
