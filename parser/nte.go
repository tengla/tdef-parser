package parser

// Nte - train data
type Nte struct {
	Type                string
	ActionCode          string
	TrainJournalOrdinal string
	NoteOrdinal         string
	NoteType            string
	NoteBody            string
}

func NewNte(records []string) *Nte {
	nte := &Nte{}
	for i, value := range records {
		switch i {
		case 0:
			nte.Type = value
		case 1:
			nte.ActionCode = value
		case 2:
			nte.TrainJournalOrdinal = value
		case 3:
			nte.NoteOrdinal = value
		case 4:
			nte.NoteType = value
		case 5:
			nte.NoteBody = value
		}
	}
	return nte
}
