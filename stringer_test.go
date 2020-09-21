package cmdtab_test

import (
	"fmt"

	cmd "gitlab.com/rwx.gg/cmdtab"
)

type ImmaStringer struct{}

func (s ImmaStringer) String() string {
	return "Hello"
}

func ExampleString() {

	f := func() string { return "Hello" }
	fmt.Println(cmd.String(f))

	s := "Hello"
	fmt.Println(cmd.String(s))

	st := ImmaStringer{} // st.String()
	fmt.Println(cmd.String(st))

	// Output:
	// Hello
	// Hello
	// Hello
}
