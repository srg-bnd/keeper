// Server setting
package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

const (
	hostUsage        = "base URL of target address"
	databaseDSNUsage = "base DSN of database"
	secretKeyUsage   = "secret key"
)

const (
	hostDefault        = ":3000"
	databaseDSNDefault = "host=127.0.0.1 user=keeper password=keeper dbname=keeper sslmode=disable"
)

type Setting struct {
	HostAddr    string `env:"HOST"`
	DatabaseDSN string `env:"DATABASE"`
	WithDBSeed  bool
	SecretKey   string `env:"SECRET_KEY"`
}

func NewSetting() *Setting {
	var setting Setting
	flag.StringVar(&setting.HostAddr, "a", hostDefault, hostUsage)
	flag.StringVar(&setting.DatabaseDSN, "d", databaseDSNDefault, databaseDSNUsage)
	flag.StringVar(&setting.SecretKey, "k", "", secretKeyUsage)
	flag.BoolVar(&setting.WithDBSeed, "seed", false, "")

	flag.Parse()
	err := env.Parse(&setting)
	if err != nil {
		panic(err)
	}

	return &setting
}
