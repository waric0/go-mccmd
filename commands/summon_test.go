package commands

import "testing"

func TestWriteSummon(t *testing.T) {
	var (
		str    string
		summon Summon
	)

	summon = Summon{}
	err := WriteSummon(&summon, &str)
	if err == nil {
		t.Fatal("No error despite incorrect input")
	}

	str = "first-line\n"
	summon = Summon{EntityType: "lightning_bolt"}
	WriteSummon(&summon, &str)
	if str != "first-line\nsummon lightning_bolt\n" {
		t.Fatal("Append processing is not successful")
	}

	inputs := []Summon{
		{EntityType: "lightning_bolt"},
		{EntityType: "lightning_bolt", SpawnPos: Position{Type: "~"}},
		{EntityType: "chicken", SpawnPos: Position{Type: "~"}, SpawnEvent: "from_egg"},
		{EntityType: "chicken", SpawnEvent: "from_egg"},
	}
	outputs := []string{
		"summon lightning_bolt\n",
		"summon lightning_bolt ~ ~ ~\n",
		"summon chicken ~ ~ ~ from_egg\n",
		"summon chicken\n",
	}
	for i := range inputs {
		str = ""
		WriteSummon(&inputs[i], &str)
		if str != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}
}
