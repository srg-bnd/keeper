// Syncs data in storage
package syncer

import (
	"context"
	"log"
	"time"

	config "github.com/srg-bnd/keeper/config/client"
	keeperAPI "github.com/srg-bnd/keeper/internal/client/api"
)

type Syncer struct {
	keeperAPI    *keeperAPI.API
	updatePeriod time.Duration
	IsRunning    bool
}

// Returns new syncer
func NewSyncer(keeperAPI *keeperAPI.API, config *config.Config) *Syncer {
	return &Syncer{
		keeperAPI:    keeperAPI,
		IsRunning:    false,
		updatePeriod: time.Duration(config.Setting.SyncerPeriod) * time.Second,
	}
}

// Starts synchronization
func (s *Syncer) Start(ctx context.Context) error {
	s.IsRunning = true

	for s.IsRunning {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := s.syncData(ctx); err != nil {
				log.Printf("Sync error: %v", err)
			}
			time.Sleep(s.updatePeriod * time.Second)
		}
	}

	return nil
}

// Stops synchronization
func (s *Syncer) Stop(ctx context.Context) error {
	s.IsRunning = false

	return nil
}

// Syncs data
func (s *Syncer) syncData(ctx context.Context) error {
	secrets, err := s.keeperAPI.AllSecrets()

	if err != nil {
	} else {
		s.keeperAPI.SecretStorage.AddList(ctx, secrets)
	}

	return nil
}
