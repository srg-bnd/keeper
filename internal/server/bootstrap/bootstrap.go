// Server Bootstrap
package bootstrap

import (
	"github.com/gin-gonic/gin"
	config "github.com/srg-bnd/keeper/config/server"
	"github.com/srg-bnd/keeper/internal/server"
	"github.com/srg-bnd/keeper/internal/server/dbtool"
	"github.com/srg-bnd/keeper/internal/server/policy"
	"github.com/srg-bnd/keeper/internal/server/services/security"
	"github.com/srg-bnd/keeper/internal/server/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB(databaseDSN string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Run() error {
	// Setting
	setting := config.NewSetting()

	// DB
	db, err := connectDB(setting.DatabaseDSN)
	if err != nil {
		return err
	}

	err = dbtool.Migrate(db)
	if err != nil {
		return err
	}

	if setting.WithDBSeed {
		err = dbtool.Seed(db)
		if err != nil {
			return err
		}
	}

	// Dependencies
	config := config.Config{
		DB:       db,
		Storage:  storage.NewStorage(db),
		Router:   gin.Default(),
		Policy:   policy.NewPolicy(db),
		Security: security.NewSecurity(setting.SecretKey),
		Setting:  setting,
	}

	// Server
	server := server.NewServer(&config)
	err = server.Start(setting.HostAddr)
	if err != nil {
		return err
	}

	return nil
}
