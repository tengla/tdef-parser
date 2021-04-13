package parser

// Dhr - diagram data
type Dhr struct {
	Type              string
	DiagramID         string
	DiagramClass      string
	DateFrom          string
	DateUntil         string
	Days              string
	Dsetcode          string
	DateRef           string
	Layer             string
	DiagramCalender   string
	Origin            string
	Destination       string
	StartTime         string
	EndTime           string
	PublicationStatus string
	Owner             string
	Status            string
	Depot             string
	SetType           string
	Distance          string
	Duration          string
	FirstMove         string
	LastMove          string
	MaxLegOrdinal     string
}

func NewDhr(record []string) *Dhr {
	dhr := &Dhr{}
	for i, value := range record {
		switch i {
		case 0:
			dhr.Type = value
		case 1:
			dhr.DiagramID = value
		case 2:
			dhr.DiagramClass = value
		case 3:
			dhr.DateFrom = value
		case 4:
			dhr.DateUntil = value
		case 5:
			dhr.Days = value
		case 6:
			dhr.Dsetcode = value
		case 7:
			dhr.DateRef = value
		case 8:
			dhr.Layer = value
		case 9:
			dhr.DiagramCalender = value
		case 10:
			dhr.Origin = value
		case 11:
			dhr.Destination = value
		case 12:
			dhr.StartTime = value
		case 13:
			dhr.EndTime = value
		case 14:
			dhr.PublicationStatus = value
		case 15:
			dhr.Owner = value
		case 16:
			dhr.Status = value
		case 17:
			dhr.Depot = value
		case 20:
			dhr.SetType = value
		case 21:
			dhr.Distance = value
		case 22:
			dhr.Duration = value
		case 23:
			dhr.FirstMove = value
		case 24:
			dhr.LastMove = value
		case 25:
			dhr.MaxLegOrdinal = value
		}
	}
	return dhr
}
