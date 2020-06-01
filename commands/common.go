package commands

import (
	"fmt"
	"strings"
)

// Position : Position Structure Definitions
////// Memo : You must explicitly set Type to 'abs' if you want to point to the origin of the world //////
type Position struct {
	Type      string
	Magnitude [3]float64
}

// GetPosStr : Returns position in command format
func GetPosStr(p *Position) string {
	var (
		str     string
		posType string
	)
	if (*p == Position{}) {
		return ""
	}
	if p.Type != "abs" {
		posType = p.Type
	}
	for i, v := range p.Magnitude {
		if posType != "" && v == 0 {
			str += fmt.Sprintf("%s", posType)
		} else {
			str += fmt.Sprintf("%s%g", posType, v)
		}
		if i+1 != len(p.Magnitude) && str != "" {
			str += " "
		}
	}
	return str
}

// Rotation : Rotation Structure Definitions
////// Memo : You must explicitly set Type to 'abs' if you want to specify an absolute angle //////
type Rotation struct {
	Type      string
	Magnitude [2]float64
}

// GetRotStr : Returns rotation in command format
func GetRotStr(r *Rotation) string {
	var (
		str     string
		rotType string
	)
	if (*r == Rotation{}) {
		return ""
	}
	if r.Type != "abs" {
		rotType = r.Type
	}
	for i, v := range r.Magnitude {
		if rotType != "" && v == 0 {
			str += fmt.Sprintf("%s", rotType)
		} else {
			str += fmt.Sprintf("%s%g", rotType, v)
		}
		if i+1 != len(r.Magnitude) && str != "" {
			str += " "
		}
	}
	return str
}

// TargetSelector : Target selector Structure Definitions
type TargetSelector struct {
	Variable   string
	EntityName string
	EntityType string
}

// GetTargetSelector : Returns target selector in command format
func GetTargetSelector(t *TargetSelector) string {
	var (
		str string
		arg string
	)
	if (*t == TargetSelector{}) || t.Variable == "" {
		return ""
	}
	str = fmt.Sprintf("%s", t.Variable)
	if t.EntityName != "" {
		if arg != "" {
			arg += ","
		}
		arg += fmt.Sprintf("name=%s", t.EntityName)
	}
	if t.EntityType != "" {
		if arg != "" {
			arg += ","
		}
		arg += fmt.Sprintf("type=%s", t.EntityType)
	}
	if arg != "" {
		str += "[" + arg + "]"
	}
	return str
}

// Command : structure definition of a collection of commands
type Command struct {
	Fill
	Summon
	Tp
}

// GetComStr : Returns the command string for indirect use
func GetComStr(c *Command) string {
	var str string
	if (*c == Command{}) {
		return ""
	}
	if (c.Fill != Fill{}) {
		WriteFill(&c.Fill, &str)
	} else if (c.Summon != Summon{}) {
		WriteSummon(&c.Summon, &str)
	} else if (c.Tp != Tp{}) {
		WriteTp(&c.Tp, &str)
	}
	return strings.TrimRight(str, "\n")
}
