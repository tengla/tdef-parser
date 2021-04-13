package parser

import "fmt"

// Tdt - Train data
type Tdt struct {
	Type                    string
	ActionCode              string
	CeROrdinal              string
	OriginJourneyLegOrdinal string
	FromLocation            string
	ToLocation              string
	TimingLoad              string
	GraphColour             string
	GraphStyle              string
	LocoType                string
	ServiceCodes            string
	Conditional             string // Behovstog
	Enbemannet              string
	Lokf                    string
	GreenCode               string
	GreenText               string
	Reservations            string
	CateringFacilities      string
	PreferredStockType      string `json:"preferred_stock_type"`
	ActualTrainLength       string `json:"actual_train_length"`
	ActualTrainWeight       string `json:"actual_train_weight"`
	SeatsRequired           string `json:"seat_required"`
	ProductCode             string `json:"product_code"`
}

// NewTdt - new tdt record
func NewTdt(record []string) *Tdt {
	tdt := &Tdt{}
	for i, value := range record {
		switch i {
		case 0:
			tdt.Type = value
		case 1:
			tdt.ActionCode = value
		case 2:
			tdt.CeROrdinal = value
		case 3:
			tdt.OriginJourneyLegOrdinal = value
		case 4:
			tdt.FromLocation = value
		case 5:
			tdt.ToLocation = value
		case 6:
			tdt.TimingLoad = value
		case 7:
			tdt.GraphColour = value
		case 8:
			tdt.GraphStyle = value
		case 9:
			tdt.LocoType = value
		case 10:
			tdt.ServiceCodes = value
		case 11:
			tdt.Conditional = value
		case 12:
			tdt.Enbemannet = value
		case 13:
			tdt.Lokf = value
		case 14:
			tdt.GreenCode = value
		case 15:
			tdt.GreenText = value
		case 16:
			tdt.Reservations = value
		case 17:
			tdt.CateringFacilities = value
		case 18:
			tdt.PreferredStockType = value
		case 19:
			tdt.ActualTrainLength = value
		case 20:
			tdt.ActualTrainWeight = value
		case 21:
			tdt.SeatsRequired = value
		case 22:
			tdt.ProductCode = value
		}
	}
	fmt.Println(len(record))
	return tdt
}
