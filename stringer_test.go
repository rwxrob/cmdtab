package cmdtab_test

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

type ImmaStringer struct{}

func (s ImmaStringer) String() string {
	return "Hello"
}

func ExampleString() {

	f := func() string { return "Hello" }
	fmt.Println(cmdtab.String(f))

	s := "Hello"
	fmt.Println(cmdtab.String(s))

	st := ImmaStringer{} // st.String()
	fmt.Println(cmdtab.String(st))

	// Output:
	// Hello
	// Hello
	// Hello
}
