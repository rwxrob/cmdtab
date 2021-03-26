package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdtab"
)

func main() {
	args := cmdtab.ArgsFromStdin(os.Args[1:]...)
	for _, r := range args {
		fmt.Println(r)
	}
}
