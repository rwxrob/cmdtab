package cmdtab

import (
	"sort"
	"strings"
)

// OmitBuiltins turns off the injection of the Builtin subcommands into the
// Main command when Execute is called. It can be assigned in any init() or
// from main() before calling Execute().
var OmitBuiltins bool

// Builtins are subcommands that are added to every Main command when
// Execute is called. This can be prevented by setting OmitBuiltins to false.
//
// Most of the builtins are hidden (beginning with underscore '_') but the
// following are so standardized they are included by default:
//
// * help [subcmd]  - very long, formatted documentation
// * version - version, copyright, license, authors, git source
//
// Any of these can be overridden by command authors simply by naming their
// own version the same. This may be desirable when creating commands in
// other languages although keeping these standard English names is
// strongly recommended due to their ubiquitous usage.
//
// The following are hidden but can be promoted by encapsulating them in
// other subcommands each with its own file, name, title, and documentation:
func Builtins() []string {
	return builtins
}

var builtins []string

func init() {

	allnames := func() []string {
		names := []string{}
		for name, _ := range Index {
			names = append(names, name)
		}
		sort.Strings(names)
		return names
	}

	// seq scan is just fine for this scale
	notbuiltin := func(name string) bool {
		for _, n := range builtins {
			if n == name {
				return false
			}
		}
		return true
	}

	names := func() []string {
		alln := allnames()
		justnames := []string{}
		for _, name := range alln {
			if notbuiltin(name) {
				justnames = append(justnames, name)
			}
		}
		return justnames
	}

	// ----------------------------------------------------------------

	_version := New("_cmdversion")
	_version.Summary = `Print the cmd package version`
	_version.Method = func(ignored []string) error { Println(Version); return nil }
	builtins = append(builtins, "_cmdversion")

	_builtins := New("_builtins")
	_builtins.Summary = `List all cmd package builtins names and summaries`
	_builtins.Method = func(ignored []string) error {
		sort.Strings(builtins)
		for _, name := range builtins {
			Print("%-14v %v\n", name, strings.TrimSpace(Index[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_builtins")

	// ----------------------------------------------------------------

	_complete := New("_complete")
	_complete.Summary = `Force completion context`
	_complete.Method = func(args []string) error {
		words := []string{Main.Name}
		words = append(words, args...)
		CompLine = strings.Join(words, " ")
		Complete()
		return nil
	}
	builtins = append(builtins, "_complete")

	// ----------------------------------------------------------------

	_index := New("_index")
	_index.Summary = `List all names and summaries from cmd package index`
	_index.Method = func(ignored []string) error {
		for _, name := range allnames() {
			Print("%-14v %v\n", name, strings.TrimSpace(Index[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_index")

	// ----------------------------------------------------------------

	_names := New("_names")
	_names.Summary = `List names, main first`
	_names.Method = func(ignored []string) error {
		Println(Main.Name)
		for _, name := range names() {
			if name != Main.Name {
				Println(name)
			}
		}
		return nil
	}
	builtins = append(builtins, "_names")

	// ----------------------------------------------------------------

	_summaries := New("_summaries")
	_summaries.Summary = `List names and summaries`
	_summaries.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(Index[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_summaries")

	// ----------------------------------------------------------------

	_versions := New("_versions")
	_versions.Summary = `List names and versions`
	_versions.Method = func(args []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(String(Index[name].Version)))
		}
		return nil
	}
	builtins = append(builtins, "_versions")

	// ----------------------------------------------------------------

	_copyrights := New("_copyrights")
	_copyrights.Summary = `List names and copyrights`
	_copyrights.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(string(String(Index[name].Copyright))))
		}
		return nil
	}
	builtins = append(builtins, "_copyrights")

	// ----------------------------------------------------------------

	_licenses := New("_licenses")
	_licenses.Summary = `List names and licenses`
	_licenses.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(String(Index[name].License)))
		}
		return nil
	}
	builtins = append(builtins, "_licenses")

	// ----------------------------------------------------------------

	_authors := New("_authors")
	_authors.Summary = `List names and authors`
	_authors.Method = func(ignored []string) error {
		for _, name := range names() {
			author := Index[name].Author
			if author == nil {
				author = ""
			}
			Print("%-14v %v\n", name, strings.TrimSpace(String(author)))
		}
		return nil
	}
	builtins = append(builtins, "_authors")

	// ----------------------------------------------------------------

	_gits := New("_gits")
	_gits.Summary = `List names and git source repos`
	_gits.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(String(Index[name].Git)))
		}
		return nil
	}
	builtins = append(builtins, "_gits")

	// ----------------------------------------------------------------

	_issues := New("_issues")
	_issues.Summary = `List names and issue reporting URLs`
	_issues.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(String(Index[name].Issues)))
		}
		return nil
	}
	builtins = append(builtins, "_issues")

	// ----------------------------------------------------------------

	_usages := New("_usages")
	_usages.Summary = `List names and usages`
	_usages.Method = func(ignored []string) error {
		for _, name := range names() {
			Print("%-14v %v\n", name, strings.TrimSpace(String(Index[name].Usage)))
		}
		return nil
	}
	builtins = append(builtins, "_usages")

	// ----------------------------------------------------------------

	_desc := New("_descriptions")
	_desc.Summary = `List names and descriptions`
	_desc.Method = func(ignored []string) error {
		emph := DisableEmphasis
		DisableEmphasis = true
		for _, name := range names() {
			Print("DESCRIPTION %v\n\n", name)
			Println(Format(String(Index[name].Description), 4, int(WinSize.Col)))
			Println()
		}
		DisableEmphasis = emph
		return nil
	}
	builtins = append(builtins, "_descriptions")

	// ----------------------------------------------------------------

	_examples := New("_examples")
	_examples.Summary = `List names and examples`
	_examples.Method = func(ignored []string) error {
		emph := DisableEmphasis
		DisableEmphasis = true
		for _, name := range names() {
			Print("EXAMPLE %v\n\n", name)
			Print(Sprint(Format(String(Index[name].Examples), 4, int(WinSize.Col))))
			Print("\n\n")
		}
		DisableEmphasis = emph
		return nil
	}
	builtins = append(builtins, "_examples")

	// ----------------------------------------------------------------

	_help_json := New("_help_json")
	_help_json.Summary = `Dump help documentation as JSON`
	_help_json.Method = func(args []string) error {
		emph := DisableEmphasis
		DisableEmphasis = true
		Println(JSON())
		DisableEmphasis = emph
		return nil
	}
	builtins = append(builtins, "_help_json")

	// ----------------------------------------------------------------

	help := New("help")
	help.Summary = "Display detailed help documentation"
	help.Method = func(args []string) error {
		c := Main
		if len(args) > 0 && Has(args[0]) {
			c = Index[args[0]]
		}
		output := Sprint(clearScreen + termHeading(
			bold+c.Name, reset+"DOCUMENTATION", bold+c.Name,
			int(WinSize.Col)) + "\n\n" +
			bold + "NAME" + reset + "\n" +
			Format(c.Title(), 4, int(WinSize.Col)) + "\n\n" +
			bold + "USAGE" + reset + "\n" +
			Indent(Emphasize(c.SprintUsage()), 4) + "\n\n")
		if len(c.vsubcommands()) > 0 {
			output += bold + "COMMANDS" + reset + "\n" +
				Indent(Emphasize(c.SprintCommandSummaries()), 4) + "\n\n"
		}
		output += bold + "DESCRIPTION" + reset + "\n" +
			Format(String(c.Description), 4, int(WinSize.Col)) + "\n\n"

		// TODO finish output
		if PagedOut {
			PrintPaged(output, "")
		} else {
			Print(output)
		}
		return nil
	}
	builtins = append(builtins, "help")

	// ----------------------------------------------------------------

	version := New("version")
	version.Summary = `Display version, author, and legal information`
	version.Method = func(args []string) error {
		emph := DisableEmphasis
		DisableEmphasis = true
		if len(args) > 0 {
			v := Index[args[0]].VersionLine()
			if v != "" {
				Println(v)
			}
			return nil
		}
		vl := Main.VersionLine()
		if vl == "" {
			return nil
		}
		Println(vl)
		for _, name := range names() {
			if name == Main.Name {
				continue
			}
			line := Index[name].VersionLine()
			if line == "" {
				continue
			}
			Println(line)
		}
		DisableEmphasis = emph
		return nil
	}
	builtins = append(builtins, "version")

	// ----------------------------------------------------------------

	_usage := New("_usage")
	_usage.Summary = "display usage summaries"
	_usage.Method = func(ignored []string) error {
		Print(Main.SprintUsage())
		return nil
	}
	builtins = append(builtins, "_usage")

}
