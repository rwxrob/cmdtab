package main

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/oglinuk/cmdtab-links" // Import the 3rd party subcommand
)

func init() {
	// Add subcommand root command
	x := cmdtab.New("iact", "links")
	x.Version = `v1.0.0`
	x.License = `MPL-v2.0`
}
