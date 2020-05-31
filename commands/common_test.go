package commands

import "testing"

func TestGetPosStr(t *testing.T) {
	inputs := []Position{
		{},
		{Type: ""},
		{Type: "abs"},
		{Magnitude: [3]float64{0, 0, 0}},
		{Type: "abs", Magnitude: [3]float64{0, 0, 0}},
		{Magnitude: [3]float64{0.1, 0, 0.1}},
		{Type: "abs", Magnitude: [3]float64{0.1, 0, 0.1}},
		{Type: "~"},
		{Type: "~", Magnitude: [3]float64{1, 1, 1}},
		{Type: "^", Magnitude: [3]float64{0, 0, 1}},
		{Type: "~", Magnitude: [3]float64{0.5}},
	}
	outputs := []string{
		"",
		"",
		"0 0 0",
		"",
		"0 0 0",
		"0.1 0 0.1",
		"0.1 0 0.1",
		"~ ~ ~",
		"~1 ~1 ~1",
		"^ ^ ^1",
		"~0.5 ~ ~",
	}
	for i := range inputs {
		result := GetPosStr(&inputs[i])
		if result != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}

	input := Position{Type: "abs", Magnitude: [3]float64{0, 0, 0}}
	output := "0 0 0"
	resultOne := GetPosStr(&input)
	resultTwo := GetPosStr(&input)
	if resultOne == output && resultOne != resultTwo {
		t.Fatal("Cannot handle reuse of the same instance")
	}
}

func TestGetRotStr(t *testing.T) {
	inputs := []Rotation{
		{},
		{Type: ""},
		{Type: "abs"},
		{Magnitude: [2]float64{0, 0}},
		{Type: "abs", Magnitude: [2]float64{0, 0}},
		{Magnitude: [2]float64{0.1, 0}},
		{Type: "abs", Magnitude: [2]float64{0.1, 0}},
		{Type: "~"},
		{Type: "~", Magnitude: [2]float64{1, 1}},
		{Type: "^", Magnitude: [2]float64{0, 1}},
		{Type: "~", Magnitude: [2]float64{0.5}},
	}
	outputs := []string{
		"",
		"",
		"0 0",
		"",
		"0 0",
		"0.1 0",
		"0.1 0",
		"~ ~",
		"~1 ~1",
		"^ ^1",
		"~0.5 ~",
	}
	for i := range inputs {
		result := GetRotStr(&inputs[i])
		if result != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}

	input := Rotation{Type: "abs", Magnitude: [2]float64{0, 0}}
	output := "0 0"
	resultOne := GetRotStr(&input)
	resultTwo := GetRotStr(&input)
	if resultOne == output && resultOne != resultTwo {
		t.Fatal("Cannot handle reuse of the same instance")
	}
}

func TestGetTargetSelector(t *testing.T) {
	inputs := []TargetSelector{
		{},
		{Variable: "@p"},
		{Variable: "@e", EntityType: "armor_stand", EntityName: "test"},
		{Variable: "@e", EntityType: "", EntityName: ""},
		{EntityName: "test"},
	}
	outputs := []string{
		"",
		"@p",
		"@e[name=test,type=armor_stand]",
		"@e",
		"",
	}
	for i := range inputs {
		result := GetTargetSelector(&inputs[i])
		if result != outputs[i] {
			t.Fatal("Write result is different than expected")
		}
	}
}
