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
