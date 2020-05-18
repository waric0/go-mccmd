package output

import "testing"

func TestGetMCPackName(t *testing.T) {
	cases := []string{"", "/", "\\", ".test", " test "}
	for _, c := range cases {
		man := &Manifest{Header: Header{Name: c}}
		_, err := GetMCPackName(man)
		if err == nil {
			t.Fatal("No error despite incorrect input")
		}
	}

	cases = []string{"test", "test-test", "test.mcpack"}
	for _, c := range cases {
		man := &Manifest{Header: Header{Name: c}}
		_, err := GetMCPackName(man)
		if err != nil {
			t.Fatal("There is an error even though the input is correct")
		}
	}
}
