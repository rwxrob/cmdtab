package cmdtab_test

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func ExampleComplete_plainnosubs() {
	cmdtab.Main = cmdtab.New("exe")
	cmdtab.CompLine = "exe" // set by shell
	cmdtab.Complete()
	// Unordered Output:
}

func ExampleComplete_plansubs() {
	cmdtab.New("list")
	cmdtab.New("add")
	cmdtab.New("remove")
	exe := cmdtab.New("exe1", "list", "add", "remove")
	cmdtab.Main = exe // simulate Execute()
	fmt.Println(exe.Has("list"))
	fmt.Println(exe.Has("add"))
	fmt.Println(exe.Has("remove"))
	fmt.Println(exe.Has("missing"))
	cmdtab.CompLine = "exe" // set by shell
	cmdtab.Complete()
	// Output:
	// true
	// true
	// true
	// false
	// list
	// add
	// remove
}

func ExampleComplete_presubs() {
	cmdtab.New("list")
	cmdtab.New("add")
	cmdtab.New("remove")
	cmdtab.Main = cmdtab.New("exe", "list", "add", "remove")
	cmdtab.CompLine = "exe l" // set by shell
	cmdtab.Complete()
	// Unordered Output:
	// list
}

func ExampleComplete_fullsub() {
	cmdtab.New("list")
	cmdtab.New("add")
	cmdtab.New("remove")
	cmdtab.Main = cmdtab.New("exe", "list", "add", "remove")
	cmdtab.CompLine = "exe list" // set by shell
	cmdtab.Complete()
	fmt.Println("nothing")
	// Unordered Output:
	// nothing
}

func ExampleComplete_subsubs() {
	cmdtab.New("sched", "all", "today")
	cmdtab.Main = cmdtab.New("exe2", "sched", "add", "remove")
	cmdtab.CompLine = "exe2 sched" // set by shell
	cmdtab.Complete()

	// Unordered Output:
	// all
	// today
}

/*
func TestComplete_hidden(t *testing.T) {
	cmdtab.KeepAlive = true
	cmdtab.New("hashidden", "another")
	cmdtab.CompLine = "hashidden"
	cmdtab.Execute("hashidden")
}
*/

func ExampleComplete_delegated() {
	x := cmdtab.New("list")
	x.Completion = func(compline string) []string {
		// could do fancier context analysis of compline
		return []string{"all", "today", "yesterday"}
	}
	cmdtab.Main = x // simulate Execute()
	cmdtab.New("add")
	cmdtab.New("remove")
	cmdtab.New("exe", "list", "add", "remove")
	cmdtab.CompLine = "exe list" // set by shell
	cmdtab.Complete()

	// Unordered Output:
	// all
	// today
	// yesterday
}
