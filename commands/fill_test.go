package commands

import "testing"

func TestWriteFill(t *testing.T) {
	var (
		str  string
		fill Fill
	)

	str = "first-line\n"
	fill = Fill{From: Position{Type: "~"}, To: Position{Type: "~"}, TileName: "stone"}
	WriteFill(&fill, &str)
	if str != "first-line\nfill ~ ~ ~ ~ ~ ~ stone 0\n" {
		t.Fatal("Append processing is not successful")
	}

	inputs := []Fill{
		{From: Position{Type: "~"}, To: Position{Type: "~"}},
		{TileName: "stone"},
		{},
	}
	for i := range inputs {
		str = ""
		err := WriteFill(&inputs[i], &str)
		if err == nil {
			t.Fatal("No error despite incorrect input")
		}
	}

	inputs = []Fill{
		{From: Position{Type: "~"}, To: Position{Type: "~"}, TileName: "stone"},
		{From: Position{Type: "~"}, To: Position{Type: "~"}, TileName: "stone", TileData: 1},
		{From: Position{Type: "abs"}, To: Position{Type: "abs"}, TileName: "stone"},
	}
	outputs := []string{
		"fill ~ ~ ~ ~ ~ ~ stone 0\n",
		"fill ~ ~ ~ ~ ~ ~ stone 1\n",
		"fill 0 0 0 0 0 0 stone 0\n",
	}
	for i := range inputs {
		str = ""
		WriteFill(&inputs[i], &str)
		if str != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}
}
