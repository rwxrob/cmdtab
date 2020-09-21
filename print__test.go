package cmdtab

import (
	"testing"
)

func TestLinecount(t *testing.T) {
	txt := `
  some
  thing
  here
  `
	if linecount(txt) != 4 {
		t.Error("linecount failed")
	}
}
