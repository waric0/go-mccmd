package commands

// Tp : Tp command Structure Definitions
type Tp struct {
	Victim      TargetSelector
	Destination Position
	XYrotation  Rotation
}

// WriteTp : Write Tp command
func WriteTp(t *Tp, str *string) {
	tp := "tp"
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
}
