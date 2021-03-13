package main

import (
	"os"

	"github.com/rwxrob/cmdtab"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	if cmdtab.IsTerminal() {
		if verbose {
			cmdtab.SmartPrintln("yes")
		}
		os.Exit(0)
	}
	if verbose {
		cmdtab.SmartPrintln("no")
	}
	os.Exit(1)
}
