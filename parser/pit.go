package parser

// Pit - Final single record
type Pit struct {
	Type    string
	Changes []*PitChange
}

type PitChange struct {
	RecordTypeCode string
	AdditionCount  string
	ChangeCount    string
	DeletionCount  string
}

func NewPit(record []string) *Pit {
	pit := &Pit{}
	for i, value := range record {
		switch i {
		case 0:
			pit.Type = value
		default:
			j := (i - 1) % 4
			if j == 0 {
				pit.Changes = append(pit.Changes, &PitChange{
					RecordTypeCode: value,
					AdditionCount:  record[i+1],
					ChangeCount:    record[i+2],
					DeletionCount:  record[i+3],
				})
			}
		}
	}
	return pit
}
