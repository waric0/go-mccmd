package commands

import (
	"errors"
	"fmt"
)

// Execute : Execute command Structure Definitions
type Execute struct {
	Origin   TargetSelector
	Position Position
	Command  Command
}

// WriteExecute : Write Execute command
func WriteExecute(e *Execute, str *string) error {
	var execute string
	if (e.Origin == TargetSelector{}) || (e.Position == Position{}) {
		return errors.New("Origin and Position must be specified")
	}
	origin := GetTargetSelector(&e.Origin)
	pos := GetPosStr(&e.Position)
	com := GetComStr(&e.Command)
	execute = fmt.Sprintf("execute %s %s %s\n", origin, pos, com)
	*str += execute
	return nil
}
