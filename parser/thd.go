package parser

// Thd - Train data
type Thd struct {
	Type                   string
	ActionCode             string
	TrainOwner             string
	TrainID                string
	InitialDateOfOperation string
	TrainDates             string
	OriginLocation         string
	OriginTime             string
	DestinationLocation    string
	DestinationTime        string
	CancelledStatus        string
	ProductCode            string
	Klasse                 string
	PriceCategory          string
	BookingClass           string
	DataOwner              string
	SeasonCode             string
	NominalDays            string
	HolydayRelationship    string
}

func (thd *Thd) GetTrainDates() ([]TrainDate, error) {
	initial, err := TimeParse(thd.InitialDateOfOperation)
	if err != nil {
		return nil, err
	}
	dates := CreateTrainDateSequence(initial, thd.TrainDates)
	return dates, nil
}

func NewThd(record []string) *Thd {
	thd := &Thd{}
	for i, value := range record {
		switch i {
		case 0:
			thd.Type = value
		case 1:
			thd.ActionCode = value
		case 2:
			thd.TrainOwner = value
		case 3:
			thd.TrainID = value
		case 4:
			thd.InitialDateOfOperation = value
		case 5:
			thd.TrainDates = value
		case 6:
			thd.OriginLocation = value
		case 7:
			thd.OriginTime = value
		case 8:
			thd.DestinationLocation = value
		case 9:
			thd.DestinationTime = value
		case 10:
			thd.CancelledStatus = value
		case 11:
			thd.ProductCode = value
		case 12:
			thd.Klasse = value
		case 13:
			thd.PriceCategory = value
		case 14:
			thd.BookingClass = value
		case 15:
			thd.DataOwner = value
		case 16:
			thd.SeasonCode = value
		case 17:
			thd.NominalDays = value
		case 18:
			thd.HolydayRelationship = value
		}
	}
	return thd
}
