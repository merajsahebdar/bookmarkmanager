package manage

import (
	"context"
	"fmt"

	"git.mills.io/prologic/bitcask"
	"github.com/merajsahebdar/bookmarkmanager/internal/app/cfg"
	"github.com/merajsahebdar/bookmarkmanager/pkg/zapcfg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// The Command represents the cli command that is used to manage bookmarks.
type Command struct{}

// Run runs the command.
func (cmd *Command) Run(cc *cfg.Context) error {
	fx.New(
		fx.Provide(
			newLogger,
			newDB,
			func() *cfg.Context {
				return cc
			},
		),
		fx.Invoke(registerManageHook),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	).Run()

	return nil
}

func newLogger(cc *cfg.Context) (*zap.Logger, error) {
	logger, err := zapcfg.NewDevelopment(cc.Debug, cc.DataHomePath+"/log").Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return logger, nil
}

func newDB(cc *cfg.Context) (*bitcask.Bitcask, error) {
	db, err := bitcask.Open(cc.DataHomePath + "/database")
	if err != nil {
		return nil, fmt.Errorf("failed to open the local database: %w", err)
	}

	return db, nil
}

func registerManageHook(lc fx.Lifecycle, db *bitcask.Bitcask) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
