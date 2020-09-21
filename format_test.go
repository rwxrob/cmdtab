package cmdtab

import "testing"

var tests = [][]string{
	{
		`plain`,
		`plain`,
	},
	{
		`
    initial and trailing newlines
    `,
		`initial and trailing newlines`,
	},
	{
		`

    two or more initial and trailing

    `,
		`two or more initial and trailing`,
	},
	{
		`
    Here is a raw:

        Roses are red
        Violets are blue

    That was indented.
    `,
		"Here is a raw:\n\n    Roses are red\n    Violets are blue\n\nThat was indented.",
	},
	{
		`
    Here is a list using trailing spaces:

    * one  
    * two  
    * three  
    
    Same as:
    
    * four  
    * five  
    * six  
    
    Ok.`,
		"Here is a list using trailing spaces:\n\n* one\n* two\n* three\n\nSame as:\n\n* four\n* five\n* six\n\nOk.",
	},
	{
		`
    Here is a paragraph
    with several lines in it
    that should be put into one line.

    And
    another one here.
    `,
		"Here is a paragraph with several lines in it that should be put into one\nline.\n\nAnd another one here.",
	},
	{
		`
    but does this  
    work
    `,
		"but does this\nwork",
	},
	{
		`
    partial indent:

      partial
        indent
      paragraph
    `,
		"partial indent:\n\npartial indent paragraph",
	},
	{
		`
        mkcmd simplecmd  
        mkcmd sk today tomorrow sessions schedule week reservations  
        mkcmd blog config set post edit delete rotate  

        This would create the following files:

        * main.go  
        * blog.go  
        * config.go  
        * set.go  
        * post.go  
        * edit.go  
        * delete.go  
        * rotate.go  

        If a cmd directory is detected the files will be created there instead:

        * cmd/blog/main.go  
        * cmd/blog/blog.go  
        * cmd/blog/config.go  
        * cmd/blog/set.go  
        * cmd/blog/post.go  
        * cmd/blog/edit.go  
        * cmd/blog/delete.go  
        * cmd/blog/rotate.go  

        These files contain the starter code for that subcommand. The main.go
        file contains the starter code for the built-in Main command. 
        `,
		"mkcmd simplecmd\nmkcmd sk today tomorrow sessions schedule week reservations\nmkcmd blog config set post edit delete rotate\n\nThis would create the following files:\n\n* main.go\n* blog.go\n* config.go\n* set.go\n* post.go\n* edit.go\n* delete.go\n* rotate.go\n\nIf a cmd directory is detected the files will be created there instead:\n\n* cmd/blog/main.go\n* cmd/blog/blog.go\n* cmd/blog/config.go\n* cmd/blog/set.go\n* cmd/blog/post.go\n* cmd/blog/edit.go\n* cmd/blog/delete.go\n* cmd/blog/rotate.go\n\nThese files contain the starter code for that subcommand. The main.go\nfile contains the starter code for the built-in Main command.",
	},
}

func TestFormat(t *testing.T) {
	for _, test := range tests {
		result := Format(test[0], 0, 80)
		//t.Logf("\n\nGOT....................\n\n%v\n\nWANT....................\n\n%v\n\n", result, test[1])
		if result != test[1] {
			t.Errorf("\n\nGOT....................\n\n%v\n\nWANT....................\n\n%v\n\n", result, test[1])
		}
	}
}

func TestEmphasize_italic(t *testing.T) {
	italic = "<italic>"
	bold = "<bold>"
	bolditalic = "<bolditalic>"
	reset = "<reset>"

	in := "here is *italic* text"
	out := Emphasize(in)
	t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	if out != "here is <italic>italic<reset> text" {
		t.Errorf("WANT:\n\n%v\n\nGOT:\n\n%v\n\n", in, out)
	}

	in = "here is **bold** text"
	out = Emphasize(in)
	t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	if out != "here is <bold>bold<reset> text" {
		t.Errorf("WANT:\n\n%v\n\nGOT:\n\n%v\n\n", in, out)
	}

	in = "here is ***bolditalic*** text"
	out = Emphasize(in)
	t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	if out != "here is <bolditalic>bolditalic<reset> text" {
		t.Errorf("WANT:\n\n%v\n\nGOT:\n\n%v\n\n", in, out)
	}

	in = "* one\n* two\n"
	out = Emphasize(in)
	t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	if out != "* one\n* two\n" {
		t.Errorf("WANT:\n\n%v\n\nGOT:\n\n%v\n\n", in, out)
	}

	in = "* *one*  \n* *two*  \nSomething else."
	out = Emphasize(in)
	//t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	if out != "* <italic>one<reset>  \n* <italic>two<reset>  \nSomething else." {
		t.Errorf("WANT:\n\n%v\n\nGOT:\n\n%v\n\n", in, out)
	}

	in = "how about <dem> apples <dere>"
	out = Emphasize(in)
	//t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	//t.Logf("'%v'", out)
	if out != "how about <<italic>dem<reset>> apples <<italic>dere<reset>>" {
		t.Errorf("IN:\n\n%v\n\nOUT:\n\n%v\n\n", in, out)
	}

	in = "Rob Muhlestein (Mr. Rob) <rwx@robs.io>"
	want := "Rob Muhlestein (Mr. Rob) <<italic>rwx@robs.io<reset>>"
	out = Emphasize(in)
	//t.Logf("\n\n%v\n\n --> \n\n%v\n\n", in, out)
	//t.Logf("'%v'", out)
	if out != want {
		t.Errorf("GOT:\n\n%v\n\nWANT:\n\n%v\n\n", out, want)
	}

}

func TestTermTitle(t *testing.T) {
	full := termHeading("left", "center", "right", 40)
	mid := termHeading("left", "center", "right", 12)
	small := termHeading("left", "center", "right", 6)
	none := termHeading("left", "center", "right", 4)

	if full != "left             center            right" {
		t.Log(full)
		t.Error("termHeading full failed")
	}
	if mid != "center right" {
		t.Log("'" + mid + "'")
		t.Error("termHeading mid failed")
	}
	if small != "center" {
		t.Error("termHeading small failed")
	}
	if none != "cent" {
		t.Error("termHeading none failed")
	}
}
