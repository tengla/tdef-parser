package parser

// Tmv - Train data move
type Tmv struct {
	Type                 string
	ActionCode           string
	JourneyLegOrdinal    string
	Origin               string
	Destination          string
	RunningLineCode      string
	TimeStart            string
	TimeEnd              string
	EntrySpeed           string
	ExitSpeed            string
	EngineeringAllowance string
	PathingAllowance     string
	PerformanceAllowance string
	Adjustment           string
	StartPivot           string
	EndPivot             string
}

func NewTmv(record []string) *Tmv {
	tmv := &Tmv{}
	for i, value := range record {
		switch i {
		case 0:
			tmv.Type = value
		case 1:
			tmv.ActionCode = value
		case 2:
			tmv.JourneyLegOrdinal = value
		case 3:
			tmv.Origin = value
		case 4:
			tmv.Destination = value
		case 5:
			tmv.RunningLineCode = value
		case 6:
			tmv.TimeStart = value
		case 7:
			tmv.TimeEnd = value
		case 8:
			tmv.EntrySpeed = value
		case 9:
			tmv.ExitSpeed = value
		case 10:
			tmv.EngineeringAllowance = value
		case 11:
			tmv.PathingAllowance = value
		case 12:
			tmv.PerformanceAllowance = value
		case 13:
			tmv.Adjustment = value
		case 14:
			tmv.StartPivot = value
		case 15:
			tmv.EndPivot = value
		}
	}
	return tmv
}
