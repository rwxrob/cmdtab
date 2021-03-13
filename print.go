package cmdtab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Print calls Sprint and prints it.
func Print(stuff ...interface{}) {
	fmt.Print(Sprint(stuff...))
}

// Println calls Sprint and prints it with a line return.
func Println(stuff ...interface{}) {
	fmt.Println(Sprint(stuff...))
}

// Sprint returns nothing if empty, acts like fmt.Sprint if it has one argument,
// or acts like Sprintf if it has more than one argument. Print can also print
// Stringers. Use Dump instead for debugging.
func Sprint(stuff ...interface{}) string {
	switch {
	case len(stuff) == 1:
		switch s := stuff[0].(type) {
		case string:
			if len(s) > 0 {
				return fmt.Sprint(s)
			}
		case Stringer:
			return fmt.Sprint(String(s))
		}
	case len(stuff) > 1:
		return fmt.Sprintf(String(stuff[0]), stuff[1:]...)
	}
	return ""
}

// Dump simply dumps the stuff passed to it to standard output. Use for
// debugging. Use Print for general printing.
func Dump(stuff ...interface{}) {
	fmt.Printf("%v\n", stuff)
}

// ConvertToJSON converts any object to its JSON string equivalent with two spaces of
// human-readable indenting. While this inflates the size for most purposes this
// is desirable even when dealing with large data sets. Technologies like
// GraphQL and offline-first progressive web apps have reduced the concern for
// total size of JSON. Usually human readability is more important. If an error
// is encountered while marshalling an ERROR key will be created with the
// string value of the error as its value.
func ConvertToJSON(thing interface{}) string {
	byt, err := json.MarshalIndent(thing, "", "  ")
	if err != nil {
		return fmt.Sprintf("{\"ERROR\": \"%v\"}", err)
	}
	return string(byt)
}

// PagedOut sets the default use of a pager application and calling
// PrintPaged() instead of Print() for non-hidden builtin subcommands.
var PagedOut = true

// PagedDefStatus is the status line passed to `less` to provide information at
// the bottom of the screen prompting the user what to do. Helpful with
// implementing help in languages besides English.
var PagedDefStatus = `Line %lb [<space>(down), b(ack), h(elp), q(quit)]`

// PrintPaged prints a string to the system pager (usually less) using the
// second argument as the custom status string (usually at the bottom). Control
// is returned to the calling program after completion. If no pager application
// is detected the regular Print() will be called intead. If status string is
// empty PagedDefStatus will be used (use " " to empty). Currently only the
// less pager is supported.
func PrintPaged(buf, status string) {
	if status == "" {
		status = PagedDefStatus
	}
	_, err := exec.LookPath("less")
	if err != nil || linecount(buf) < int(WinSize.Row) {
		Print(buf)
		return
	}
	less := exec.Command("less", "-r", "-Ps"+status)
	less.Stdin = strings.NewReader(buf)
	less.Stdout = os.Stdout
	less.Run()
}

func linecount(buf string) int {
	return bytes.Count([]byte(buf), []byte{'\n'})
}

// SmartPrintln calls Println() or Print() based on if IsTerminal()
// returns true or not.
func SmartPrintln(a ...interface{}) {
	if IsTerminal() {
		Println(a...)
		return
	}
	Print(a...)
}
