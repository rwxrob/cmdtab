package cmdtab_test

import (
	"fmt"

	cmd "gitlab.com/rwxrob/cmdtab"
)

func ExampleCommand_VersionLine() {

	c := cmd.New("foo")
	fmt.Println(c.VersionLine())
	c.Version = `v1.0.0`
	fmt.Println(c.VersionLine())
	c.Copyright = `© Rob`
	fmt.Println(c.VersionLine())
	c.License = `Apache 2.0`
	fmt.Println(c.VersionLine())
	c.Copyright = ""
	fmt.Println(c.VersionLine())

	// Output:
	//
	// foo v1.0.0
	// foo v1.0.0 © Rob
	// foo v1.0.0 © Rob (Apache 2.0)
	// foo v1.0.0 (Apache 2.0)
}

func ExampleCommand_MarshalJSON() {

	c := cmd.New("mycmd", "subcmd1", "subcmd2")
	c.Author = `Rob`
	c.Copyright = `© Rob`
	c.Description = `Just for testing.`
	c.Examples = `examples`
	c.Git = `gitlab.com/rwxrob/cmdtab`
	c.Issues = `https://gitlab.com/rwxrob/cmdtab/issues`
	c.License = `Apache 2.0`
	//c.SeeAlso = ``
	c.Summary = `summary`
	//c.Usage = ``
	c.Version = `v1.0.0`

	// without these Usage won't return because they do not exist
	cmd.New("subcmd1")
	cmd.New("subcmd2")

	fmt.Println(c)

	// Unordered Output:
	// {
	//   "Author": "Rob",
	//   "Copyright": "© Rob",
	//   "Description": "Just for testing.",
	//   "Examples": "examples",
	//   "Git": "gitlab.com/rwxrob/cmdtab",
	//   "Issues": "https://gitlab.com/rwxrob/cmdtab/issues",
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
