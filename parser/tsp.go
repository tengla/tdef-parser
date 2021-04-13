package parser

// Tsp - Train data
type Tsp struct {
	Type                string
	ActionCode          string
	JourneyLegOrdinal   string
	Location            string
	TimeStart           string
	TimeEnd             string
	Platform            string
	AdvertisedArrival   string
	AdvertisedDeparture string
	StartPivot          string
	EndPivot            string
	Activities          string
	NonCommercialStop   string
}

func NewTsp(record []string) *Tsp {
	tsp := &Tsp{}
	for i, value := range record {
		switch i {
		case 0:
			tsp.Type = value
		case 1:
			tsp.ActionCode = value
		case 2:
			tsp.JourneyLegOrdinal = value
		case 3:
			tsp.Location = value
		case 4:
			tsp.TimeStart = value
		case 5:
			tsp.TimeEnd = value
		case 6:
			tsp.Platform = value
		case 7:
			tsp.AdvertisedArrival = value
		case 8:
			tsp.AdvertisedDeparture = value
		case 9:
			tsp.StartPivot = value
		case 10:
			tsp.EndPivot = value
		case 11:
			tsp.Activities = value
		case 12:
			tsp.NonCommercialStop = value
		}
	}
	return tsp
}
