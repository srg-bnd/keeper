package actions

import (
	"github.com/srg-bnd/keeper/internal/client/syncer"

	uCLI "github.com/urfave/cli/v2"
)

type SyncerAction struct {
	syncer *syncer.Syncer
}

func NewSyncerAction(syncer *syncer.Syncer) *SyncerAction {
	return &SyncerAction{
		syncer: syncer,
	}
}

func (a *SyncerAction) Do(c *uCLI.Context) error {
	a.syncer.Start(c.Context)

	return nil
}
