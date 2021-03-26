package cmdtab_test

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

func ExamplePrintln() {

	stringer := func() string { return "something 1" }

	cmdtab.Println("something 1")
	cmdtab.Println("some%v %v", "thing", 1)
	cmdtab.Println(stringer)
	cmdtab.Println()
	cmdtab.Println(nil)

	// Output:
	// something 1
	// something 1
	// something 1
	//
	//
}

func ExampleSprint() {

	stringer := func() string { return "something 1" }

	fmt.Println(cmdtab.Sprint("something 1"))
	fmt.Println(cmdtab.Sprint("some%v %v", "thing", 1))
	fmt.Println(cmdtab.Sprint(stringer))
	fmt.Println(cmdtab.Sprint())

	// Output:
	// something 1
	// something 1
	// something 1
	//
}

func ExampleConvertToJSON() {

	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}

	fmt.Println(cmdtab.ConvertToJSON(sample))

	// Unordered output:
	//
	// {
	//   "array": [
	//     "blah",
	//     "another"
	//   ],
	//   "float": 1,
	//   "int": 1,
	//   "map": {
	//     "blah": "another"
	//   },
	//   "string": "some thing"
	// }
}
