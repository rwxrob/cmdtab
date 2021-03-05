package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("hello")
	x.Version = `v1.0.0` // Use semantic versioning
	x.Usage = `[<word>]`
	x.Summary = `this is a brief description of the command`
	x.Description = `This is a more detailed summary of the command. In
	this context, the *hello* command outputs 'Hello [args]!'`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			fmt.Println("Hello world!")
		} else {
			fmt.Printf("Hello %s!\n", args)
		}
		return nil
	}
}
