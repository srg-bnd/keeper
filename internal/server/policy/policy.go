package policy

import (
	"github.com/srg-bnd/keeper/internal/server/models"
	"gorm.io/gorm"
)

type Policy struct {
	db *gorm.DB
}

func NewPolicy(db *gorm.DB) *Policy {
	return &Policy{
		db: db,
	}
}

type PolicyData struct {
	User      *models.User
	OwnerType models.PermissionOwnerType
	Action    *models.PermissionAction
	OwnerID   *uint
}

func (p *Policy) HasPermission(policyData *PolicyData) bool {
	if policyData.Action == nil || policyData.OwnerType == "" {
		return false
	}

	roles := policyData.User.Roles
	if len(roles) == 0 {
		return false
	}

	roleIDs := make([]uint, 0, len(roles))
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}

	query := p.db.
		Where("action = ? AND owner_type = ?", policyData.Action, policyData.OwnerType).
		Where("role_id IN ?", roleIDs)

	if policyData.OwnerID != nil {
		query = query.Where("owner_id = ?", &policyData.OwnerID)
	}

	var permission models.Permission
	if err := query.First(&permission).Error; err != nil {
		return false
	}

	var count int64
	if err := query.Model(&models.Permission{}).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
