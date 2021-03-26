package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func main() {
	args := cmdtab.ArgsFromStdin("one", "two")
	for n, r := range args {
		fmt.Printf("arg%v: %v\n", n, r)
	}
}
