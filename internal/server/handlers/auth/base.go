package auth

import (
	"github.com/srg-bnd/keeper/internal/server/policy"
	"github.com/srg-bnd/keeper/internal/server/services/security"
	"github.com/srg-bnd/keeper/internal/server/storage"
)

type AuthHandler struct {
	Policy   *policy.Policy
	Storage  storage.UserRepositories
	Security *security.Security
}
