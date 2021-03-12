package main

import (
	"fmt"
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
			fmt.Println("yes")
		}
		os.Exit(0)
	}
	if verbose {
		fmt.Println("no")
	}
	os.Exit(1)
}
