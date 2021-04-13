package parser

// Dla - diagram data
type Dla struct {
	Type         string
	ActionCode   string
	LocationCode string
	StandardTime string
	Pay          string
	Platform     string
	Prohibited   string
}

func NewDla(record []string) *Dla {
	dla := &Dla{}
	for i, value := range record {
		switch i {
		case 0:
			dla.Type = value
		case 1:
			dla.ActionCode = value
		case 2:
			dla.LocationCode = value
		case 3:
			dla.StandardTime = value
		case 4:
			dla.Pay = value
		case 5:
			dla.Platform = value
		case 6:
			dla.Prohibited = value
		}
	}
	return dla
}
