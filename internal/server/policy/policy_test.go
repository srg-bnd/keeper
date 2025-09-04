package policy_test

import (
	"testing"

	"github.com/srg-bnd/keeper/internal/server/policy"
	"gorm.io/gorm"
)

func TestPolicy_HasPermission(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		policyData *policy.PolicyData
		want       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := policy.NewPolicy(tt.db)
			got := p.HasPermission(tt.policyData)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}
