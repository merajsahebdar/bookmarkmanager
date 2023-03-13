package main

import (
	"os"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/merajsahebdar/bookmarkmanager/internal/app/cfg"
	"github.com/merajsahebdar/bookmarkmanager/internal/app/manage"
)

const (
	appBundleID = "com.merajsahebdar.apps.bookmarkmanager"
)

var cli struct {
	Debug  bool           `help:"Enable the debug mode."`
	Manage manage.Command `cmd:"manage" help:"Start the bookmark manager."`
}

func main() {
	ctx := kong.Parse(&cli)

	// Prepare the data home directory...
	dataHomePath := xdg.DataHome + "/" + appBundleID
	if f, err := os.Stat(dataHomePath); os.IsNotExist(err) {
		if createErr := os.Mkdir(dataHomePath, 0755); createErr != nil {
			ctx.Fatalf("failed to create the data home: %s", createErr.Error())
		}
	} else if !f.IsDir() {
		ctx.Fatalf("the data home is already occupied")
	}

	if err := ctx.Run(&cfg.Context{
		Debug:        cli.Debug,
		DataHomePath: dataHomePath,
	}); err != nil {
		ctx.FatalIfErrorf(err)
	}
}
