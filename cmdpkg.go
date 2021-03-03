package cmdtab

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Index contains all the commands available as subcommands with one of them
// being set to Main. Commands are created and registered with New().
//
// WARNING: Index is made public only to cover the unforeseen needs of command
// creators to inspect their own Command set but should not be relied upon
// heavily to avoid unnecessary coupling dependencies between commands. While
// the Index type is guaranteed never to change, direct manipulation of the
// Index is strongly discouraged. Use one of the helper package methods instead
// (or request the addition of one).
var Index = map[string]*Command{}

// Visible returns a new map containing only pointers to the visible Commands.
func Visible() map[string]*Command {
	vis := make(map[string]*Command)
	for k, v := range Index {
		if !v.Hidden() {
			vis[k] = v
		}
	}
	return vis
}

// Hidden returns a new map containing only pointers to the hidden Commands.
func Hidden() map[string]*Command {
	vis := make(map[string]*Command)
	for k, v := range Index {
		if v.Hidden() {
			vis[k] = v
		}
	}
	return vis
}

// CompLine is set if a completion context from the shell is detected. (For
// Bash it is COMP_LINE. See Programmable Completion in the bash man page.)
var CompLine string

func init() {
	// leave room for other potention shell completion environment variable
	// names to be added in the future
	CompLine = os.Getenv("COMP_LINE")
}

// Args returns a reliable collection of arguments to the executable.
//
// WARNING: Although the first the element of os.Args is usually the binary
// of the compiled program executed it is never reliable and significantly
// differs depending on operating system and method of program execution.
// The first argument is therefore stripped completely leaving only the
// arguments to be processed. The cmd.Args package variable can also be
// set during testing to check cmd.Execute() behavior.
var Args = os.Args[1:]

// Call allows any indexed subcommand to be called directly by name. Avoid
// using this method as much as possible since it creates very tight coupling
// dependencies between commands. It is included primarily publicly so that
// builtin commands like help, usage, and version can be wrapped with
// internationalized aliases.
func Call(name string, args []string) error {
	defer TrapPanic()
	command, has := Index[name]
	if !has {
		return Unimplemented(name)
	}
	return command.Call(args)
}

// FixCompLine activates an attempt to correct the CompLine to work best with
// completion. For example, when an executable that uses the cmd package is
// renamed or is called as a path and not just the single command name. True by
// default. Set to false to leave the CompLine exactly as it is detected but
// note that depending on a specific form of CompLine may not be consistent
// across operating systems.
var FixCompLine bool = true

// Complete calls complete on the Main command passing it CompLine. No
// verification of Main's existence is checked. The CompLine is always changed
// to match the actual name of the Main command even if the executable name has
// been changed or called as an alias. This ensures proper tab completion no
// matter what the actual executable is called.
func Complete() {
	if !FixCompLine {
		Main.Complete(CompLine)
		return
	}
	i := strings.Index(CompLine, " ")
	if i < 0 {
		Main.Complete(Main.Name)
	} else {
		Main.Complete(Main.Name + CompLine[i:])
	}
}

// Main contains the main command passed to Execute to start the
// program. While it can be changed by Subcommands it usually should not be.
var Main *Command

// Execute traps all panics (see Panic), detects completion and does it, or
// sets Main to the command name passed, injects the Builtin subcommands
// (unless OmitBuiltins is true), looks up the named command from the
// command Index and calls it passing cmd.Args. Execute alway exits
// the program.
func Execute(name string) {
	defer TrapPanic()
	command, has := Index[name]
	if !has {
		ExitUnimplemented(name)
	}
	Main = command
	if !OmitBuiltins {
		for _, name := range builtins {
			Main.Add(name)
		}
	}
	if CompLine != "" {
		Complete()
		Exit()
	}
	err := command.Call(Args)
	if err != nil {
		ExitError(err)
	}
	Exit()
}

// ExitExec exits the currently running Go program and hands off memory
// and control to the executable passed as the first in a string of
// arguments along with the arguments to pass along to the called
// executable. This is only supported on systems that support Go's
// syscall.Exec() and underlying execve() system call.
func ExitExec(xnargs ...string) error {
	return syscall.Exec(xnargs[0], xnargs, os.Environ())
}

// Has looks for the named subcommand in the Index.
func Has(name string) bool {
	_, has := Index[name]
	return has
}

// New initializes a new command and subcommands (adding them to the
// internal subcommand index) and returns a pointer to the command. Note
// that the subcmds added do not create a new Command in the Index, they
// are merely added to the list returned by Subcommands.
func New(name string, subcmds ...string) *Command {
	c := new(Command)
	c.Name = name
	c.Other = map[string]Stringer{}
	c.subcommands = []string{}
	if len(subcmds) > 0 {
		c.Add(subcmds...)
	}
	c.Usage = func() string { return c.defaultUsage() }
	Index[name] = c
	return c
}

// Version contains the semantic version of the cmd package used. This value
// is printed with the version builtin subcommand.
const Version = `v0.1.0`

// JSON returns a JSON representation of the state of the cmd package including
// the main command and all subcommands from the internal index. This can be
// useful when providing documentation in a structured data format that can be
// easily shared and rendered in different ways. The json builtin simply calls
// this and prints it. Empty values are always omitted. (See
// Command.MarshalJSON() as well.)
func JSON() string {
	s := make(map[string]interface{})
	s["PackageVersion"] = Version
	s["Main"] = Main.Name
	s["Commands"] = Visible()
	return ConvertToJSON(s)
}

// WaitForInterrupt just blocks until an interrupt signal is received.
// It should only be called from the main goroutine. It takes a single
// context.CancelFunc that is designed to signal everything to stop
// cleanly before exiting.
func WaitForInterrupt(cancel context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer cancel()
	log.Println("waiting for interrupt")
	<-interrupt
	log.Println("received interrupt")
}
