package cmdtab

import "testing"

func TestCommand_SubcommandUsage(t *testing.T) {
	one := New("one")
	one.Usage = "some args [here]"
	two := New("two")
	two.Usage = func() string {
		return "some args for two [here]"
	}
	New("three")
	// three has nil Usage
	command := New("scusage", "one", "two", "three")
	subcmds := command.Subcommands()
	for i, usage := range command.SubcommandUsage() {
		t.Logf("%v %v %v\n", "command", subcmds[i], String(usage))
	}
	t.Logf(command.SprintUsage())
}
