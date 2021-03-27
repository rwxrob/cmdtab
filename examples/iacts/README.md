# Import Another [cmdtab](https://github.com/rwxrob/cmdtab) Subcommand (iacts)

This example demonstrates the power of utilizing the existing underlying
infrastructure that Go and modules provide.

1. Include the 3rd party module in the imports as an anonymous import
```Go
import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdtab-pomo
)
```

2. Add the subcommand to the root command
```Go
x := cmdtab.New("<root>", "pomo")
```

3. `go mod tidy`

4. `go install`

5. `complete -C <root> <root>`

There is the potential to embed a hidden command that is built into
the library itself, allowing for *any* cmdtab-based CLI to be able
to import and manage 3rd party cmdtab-based subcommands.

## Getting Started

`go mod tidy`

`go install

`iacts help links`
