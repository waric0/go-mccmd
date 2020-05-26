package commands

import "fmt"

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
