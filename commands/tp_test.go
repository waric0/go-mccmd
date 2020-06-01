package commands

import "testing"

func TestWriteTp(t *testing.T) {
	var (
		str string
		tp  Tp
	)

	tp = Tp{}
	err := WriteTp(&tp, &str)
	if err == nil {
		t.Fatal("No error despite incorrect input")
	}

	str = "first-line\n"
	tp = Tp{Victim: TargetSelector{Variable: "@p"}}
	WriteTp(&tp, &str)
	if str != "first-line\ntp @p\n" {
		t.Fatal("Append processing is not successful")
	}

	inputs := []Tp{
		{Victim: TargetSelector{Variable: "@p"}},
		{Victim: TargetSelector{Variable: "@p"}, Destination: Position{Type: "~"}},
		{Victim: TargetSelector{Variable: "@p"}, Destination: Position{Type: "~"}, XYrotation: Rotation{Type: "~"}},
		{Victim: TargetSelector{Variable: "@p"}, XYrotation: Rotation{Type: "~"}},
		{Destination: Position{Type: "~"}, XYrotation: Rotation{Type: "~"}},
	}
	outputs := []string{
		"tp @p\n",
		"tp @p ~ ~ ~\n",
		"tp @p ~ ~ ~ ~ ~\n",
		"tp @p\n",
		"tp ~ ~ ~ ~ ~\n",
	}
	for i := range inputs {
		str = ""
		WriteTp(&inputs[i], &str)
		if str != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}
}
