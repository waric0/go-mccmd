package commands

import (
	"errors"
	"fmt"
)

// Fill : Fill command Structure Definitions
type Fill struct {
	From     Position
	To       Position
	TileName string
	TileData int
}

// WriteFill : Write Fill command
func WriteFill(f *Fill, str *string) error {
	var fill string
	if f.TileName == "" {
		return errors.New("TileName must be specified")
	}
	if (f.From == Position{} || f.To == Position{}) {
		return errors.New("From/To Position must be specified")
	}
	from := GetPosStr(&f.From)
	to := GetPosStr(&f.To)
	fill = fmt.Sprintf("fill %s %s %s %d\n", from, to, f.TileName, f.TileData)
	*str += fill
	return nil
}
