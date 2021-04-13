package parser

// Bid - train data
type Bid struct {
	Type       string
	BusinessID string
	TrainDates string
}

func NewBid(record []string) *Bid {
	bid := &Bid{}
	for i, value := range record {
		switch i {
		case 0:
			bid.Type = value
		case 1:
			bid.BusinessID = value
		case 2:
			bid.TrainDates = value
		}
	}
	return bid
}
