package cmdtab_test

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func ExampleCommand_VersionLine() {

	x := cmdtab.New("foo")
	fmt.Println(x.VersionLine())
	x.Version = `v1.0.0`
	fmt.Println(x.VersionLine())
	x.Copyright = `© Rob`
	fmt.Println(x.VersionLine())
	x.License = `Apache 2.0`
	fmt.Println(x.VersionLine())
	x.Copyright = ""
	fmt.Println(x.VersionLine())

	// Output:
	//
	// foo v1.0.0
	// foo v1.0.0 © Rob
	// foo v1.0.0 © Rob (Apache 2.0)
	// foo v1.0.0 (Apache 2.0)
}

func ExampleCommand_MarshalJSON() {

	x := cmdtab.New("mycmd", "subcmd1", "subcmd2")
	x.Author = `Rob`
	x.Copyright = `© Rob`
	x.Description = `Just for testing.`
	x.Examples = `examples`
	x.Git = `github.com/rwxrob/cmdtab`
	x.Issues = `https://github.com/rwxrob/cmdtab/issues`
	x.License = `Apache 2.0`
	//x.SeeAlso = ``
	x.Summary = `summary`
	//x.Usage = ``
	x.Version = `v1.0.0`

	// without these Usage won't return because they do not exist
	cmdtab.New("subcmd1")
	cmdtab.New("subcmd2")

	fmt.Println(x)

	// Unordered Output:
	// {
	//   "Author": "Rob",
	//   "Copyright": "© Rob",
	//   "Description": "Just for testing.",
	//   "Examples": "examples",
	//   "Git": "github.com/rwxrob/cmdtab",
	//   "Issues": "https://github.com/rwxrob/cmdtab/issues",
	//   "License": "Apache 2.0",
	//   "Name": "mycmd",
	//   "Subcommands": [
	//     "subcmd1",
	//     "subcmd2"
	//   ],
	//   "Summary": "summary",
	//   "Version": "v1.0.0"
	// }

}
