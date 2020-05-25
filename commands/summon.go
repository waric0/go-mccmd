package commands

import (
	"errors"
	"fmt"
)

// Summon : Summon command Structure Definitions
type Summon struct {
	EntityType string
	SpawnPos   Position
	SpawnEvent string
	NameTag    string
}

// WriteSummon : Write Summon command
func WriteSummon(s *Summon, str *string) error {
	var summon string
	if s.EntityType == "" {
		return errors.New("EntityType must be specified")
	}
	summon = fmt.Sprintf("summon %s", s.EntityType)
	if (s.SpawnPos != Position{}) {
		pos := GetPosStr(&s.SpawnPos)
		summon += " " + pos
		if s.SpawnEvent != "" {
			summon += " " + s.SpawnEvent
			if s.NameTag != "" {
				summon += " " + s.NameTag
			}
		}
	}
	*str += summon + "\n"
	return nil
}
