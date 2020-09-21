package cmdtab_test

import (
	"fmt"

	cmd "gitlab.com/rwx.gg/cmdtab"
)

func ExamplePrintln() {

	stringer := func() string { return "something 1" }

	cmd.Println("something 1")
	cmd.Println("some%v %v", "thing", 1)
	cmd.Println(stringer)
	cmd.Println()
	cmd.Println(nil)

	// Output:
	// something 1
	// something 1
	// something 1
	//
	//
}

func ExampleSprint() {

	stringer := func() string { return "something 1" }

	fmt.Println(cmd.Sprint("something 1"))
	fmt.Println(cmd.Sprint("some%v %v", "thing", 1))
	fmt.Println(cmd.Sprint(stringer))
	fmt.Println(cmd.Sprint())

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

	fmt.Println(cmd.ConvertToJSON(sample))

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
