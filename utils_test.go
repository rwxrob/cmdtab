package cmdtab

import "testing"

func TestFound(t *testing.T) {
	t.Log(Found("README.md"))
	t.Log(Found("READMME.md"))
}
