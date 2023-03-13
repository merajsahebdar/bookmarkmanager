package main

import (
	"os"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/merajsahebdar/bookmarkmanager/internal/app/cfg"
	"github.com/merajsahebdar/bookmarkmanager/internal/cmd/cli"
)

const (
	appBundleID = "com.merajsahebdar.apps.bookmarkmanager"
)

func main() {
	ctx := kong.Parse(&cli.CLI)

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
		Debug:        cli.CLI.Debug,
		DataHomePath: dataHomePath,
	}); err != nil {
		ctx.FatalIfErrorf(err)
	}
}
