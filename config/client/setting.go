// Client setting
package config

import (
	"github.com/caarlos0/env/v11"
)

type Setting struct {
	AuthToken     string `env:"AUTH_TOKEN"`
	EncryptionKey string `env:"ENCRYPTION_KEY"`
	ServerHost    string `env:"SERVER_HOST" envDefault:"localhost:3000"`
	SyncerPeriod  uint   `env:"SYNCER_PERIOD" envDefault:"10"`
}

func NewSetting() (*Setting, error) {
	var setting Setting

	err := env.Parse(&setting)
	if err != nil {
		return nil, err
	}

	return &setting, nil
}
