package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSetting_Defaults(t *testing.T) {
	setting, err := NewSetting()

	assert.NoError(t, err)
	assert.Empty(t, setting.AuthToken)
	assert.Empty(t, setting.EncryptionKey)
	assert.Equal(t, "localhost:3000", setting.ServerHost)
	assert.Equal(t, uint(10), setting.SyncerPeriod)
}
