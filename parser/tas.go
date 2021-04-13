package parser

// Tas - train data
type Tas struct {
	Type                string
	ActionCode          string
	TrainID             string
	Location            string
	AssociationType     string
	MainAssociatedTrain string
	TrainDates          string
	AssociatedTrainID   string
	MinTimeDuration     string
	MaxTimeDuration     string
	OffsetDays          string
	CancelledStatus     string
	Season              string
	NominalDays         string
	HolydayRelationship string
}

func NewTas(record []string) *Tas {
	tas := &Tas{}
	for i, value := range record {
		switch i {
		case 0:
			tas.Type = value
		case 1:
			tas.ActionCode = value
		case 2:
			tas.TrainID = value
		case 3:
			tas.Location = value
		case 4:
			tas.AssociationType = value
		case 5:
			tas.MainAssociatedTrain = value
		case 6:
			tas.TrainDates = value
		case 7:
			tas.AssociatedTrainID = value
		case 8:
			tas.MinTimeDuration = value
		case 9:
			tas.MaxTimeDuration = value
		case 10:
			tas.OffsetDays = value
		case 11:
			tas.CancelledStatus = value
		case 12:
			tas.Season = value
		case 13:
			tas.NominalDays = value
		case 14:
			tas.HolydayRelationship = value
		}
	}
	return tas
}
