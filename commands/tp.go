package commands

import "errors"

// Tp : Tp command Structure Definitions
type Tp struct {
	Victim      TargetSelector
	Destination Position
	XYrotation  Rotation
}

// WriteTp : Write Tp command
func WriteTp(t *Tp, str *string) error {
	var tp string
	if (*t == Tp{}) {
		return errors.New("Victim or Destination must be specified")
	}
	tp = "tp"
	if (t.Victim != TargetSelector{}) {
		victim := GetTargetSelector(&t.Victim)
		tp += " " + victim
	}
	if (t.Destination != Position{}) {
		pos := GetPosStr(&t.Destination)
		tp += " " + pos
		if (t.XYrotation != Rotation{}) {
			rot := GetRotStr(&t.XYrotation)
			tp += " " + rot
		}
	}
	*str += tp + "\n"
	return nil
}
