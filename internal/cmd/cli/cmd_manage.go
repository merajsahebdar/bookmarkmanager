package cli

import (
	"fmt"

	"github.com/merajsahebdar/bookmarkmanager/pkg/zapcfg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// The ManageCommand represents the cli command that is used to manage bookmarks.
type ManageCommand struct{}

// Run runs the command.
func (cmd *ManageCommand) Run(cc *Context) error {
	fx.New(
		fx.Provide(
			func() (*zap.Logger, error) {
				logger, err := zapcfg.NewDevelopment(cc.Debug, cc.DataHomePath+"/log").Build()
				if err != nil {
					return nil, fmt.Errorf("failed to build logger: %w", err)
				}

				return logger, nil
			},
		),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	).Run()

	return nil
}
