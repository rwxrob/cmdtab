package cmdtab_test

import (
	"fmt"

	cmd "gitlab.com/rwxrob/cmdtab"
)

func ExampleComplete_plainnosubs() {
	cmd.Main = cmd.New("exe")
	cmd.CompLine = "exe" // set by shell
	cmd.Complete()
	// Unordered Output:
}

func ExampleComplete_plansubs() {
	cmd.New("list")
	cmd.New("add")
	cmd.New("remove")
	exe := cmd.New("exe1", "list", "add", "remove")
	cmd.Main = exe // simulate Execute()
	fmt.Println(exe.Has("list"))
	fmt.Println(exe.Has("add"))
	fmt.Println(exe.Has("remove"))
	fmt.Println(exe.Has("missing"))
	cmd.CompLine = "exe" // set by shell
	cmd.Complete()
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
	cmd.New("list")
	cmd.New("add")
	cmd.New("remove")
	cmd.Main = cmd.New("exe", "list", "add", "remove")
	cmd.CompLine = "exe l" // set by shell
	cmd.Complete()
	// Unordered Output:
	// list
}

func ExampleComplete_fullsub() {
	cmd.New("list")
	cmd.New("add")
	cmd.New("remove")
	cmd.Main = cmd.New("exe", "list", "add", "remove")
	cmd.CompLine = "exe list" // set by shell
	cmd.Complete()
	fmt.Println("nothing")
	// Unordered Output:
	// nothing
}

func ExampleComplete_subsubs() {
	cmd.New("sched", "all", "today")
	cmd.Main = cmd.New("exe2", "sched", "add", "remove")
	cmd.CompLine = "exe2 sched" // set by shell
	cmd.Complete()

	// Unordered Output:
	// all
	// today
}

/*
func TestComplete_hidden(t *testing.T) {
	cmd.KeepAlive = true
	cmd.New("hashidden", "another")
	cmd.CompLine = "hashidden"
	cmd.Execute("hashidden")
}
*/

func ExampleComplete_delegated() {
	list := cmd.New("list")
	list.Completion = func(compline string) []string {
		// could do fancier context analysis of compline
		return []string{"all", "today", "yesterday"}
	}
	cmd.Main = list // simulate Execute()
	cmd.New("add")
	cmd.New("remove")
	cmd.New("exe", "list", "add", "remove")
	cmd.CompLine = "exe list" // set by shell
	cmd.Complete()

	// Unordered Output:
	// all
	// today
	// yesterday
}
