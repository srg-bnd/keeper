package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSetting_WithEnvVars(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "test-token")
	t.Setenv("ENCRYPTION_KEY", "test-key")
	t.Setenv("SERVER_HOST", "custom-host:8080")
	t.Setenv("SYNCER_PERIOD", "15")

	setting, err := NewSetting()

	assert.NoError(t, err)
	assert.Equal(t, "test-token", setting.AuthToken)
	assert.Equal(t, "test-key", setting.EncryptionKey)
	assert.Equal(t, "custom-host:8080", setting.ServerHost)
	assert.Equal(t, uint(15), setting.SyncerPeriod)
}
