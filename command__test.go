package cmdtab

import "testing"

func TestCommand_SubcommandUsage(t *testing.T) {
	x1 := New("one")
	x1.Usage = "some args [here]"
	x2 := New("two")
	x2.Usage = func() string {
		return "some args for two [here]"
	}
	New("three")
	// three has nil Usage
	x3 := New("scusage", "one", "two", "three")
	subcmds := x3.Subcommands()
	for i, usage := range x3.SubcommandUsage() {
		t.Logf("%v %v %v\n", "command", subcmds[i], String(usage))
	}
	t.Logf(x3.SprintUsage())
}
