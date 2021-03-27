package main

import "github.com/rwxrob/cmdtab"

func init() {
	x := cmdtab.New("basic", "hello")
	x.Version = `v1.0.0` // Use semantic versioning
	x.Usage = `[hello <word>]`
	x.Summary = `this is a short description of the command`
	x.Description = `This is a more detailed summary of the command. In
	this context, the *minimal* command is an example of the minimal
	implementation of the cmdtab library. It has one subcommand (hello)
	which can be used with or without a <word>. If a word is provided
	it will output 'Hello <word>!'. If a word is not provided it will
	output 'Hello world!'.`
	x.License = `MPL-v2.0`
	// For all of the available variables please reference the `command.go`
	// file found at	https://github.com/rwxrob/cmdtab/blob/main/command.go#L9
}
