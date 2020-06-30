package commands

import (
	"testing"
)

func TestWriteExecute(t *testing.T) {
	var (
		str     string
		execute Execute
	)
	str = "first-line\n"
	tp := Tp{Victim: TargetSelector{Variable: "@p"}}
	execute = Execute{Origin: TargetSelector{Variable: "@p"}, Position: Position{Type: "~"}, Command: Command{Tp: tp}}
	WriteExecute(&execute, &str)
	if str != "first-line\nexecute @p ~ ~ ~ tp @p\n" {
		t.Fatal("Append processing is not successful")
	}

	inputs := []Execute{
		{},
		{Origin: TargetSelector{Variable: "@p"}},
		{Position: Position{Type: "~"}},
		{Command: Command{Tp: tp}},
	}
	for i := range inputs {
		str = ""
		err := WriteExecute(&inputs[i], &str)
		if err == nil {
			t.Fatal("No error despite incorrect input")
		}
	}

	inputs = []Execute{
		{Origin: TargetSelector{Variable: "@p"}, Position: Position{Type: "~"}, Command: Command{Tp: tp}},
	}
	outputs := []string{
		"execute @p ~ ~ ~ tp @p\n",
	}
	for i := range inputs {
		str = ""
		WriteExecute(&inputs[i], &str)
		if str != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}
}
