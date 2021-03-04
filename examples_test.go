package cmdtab_test

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func ExampleIndex() {
	cmdtab.New("foo", "bar")
	dex := cmdtab.Index()
	fmt.Println(dex["foo"])
	// Output:
	// {
	//   "Name": "foo",
	//   "Subcommands": [
	//     "bar"
	//   ]
	// }

}
