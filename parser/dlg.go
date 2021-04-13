package parser

// Dlg - diagram data
type Dlg struct {
	Type            string
	LegType         string
	Ordinal         string
	ActionCode      string
	Orig            string
	Dest            string
	TimeStart       string
	TimeEnd         string
	OrigPlat        string
	DestPlat        string
	OrigDirection   string
	DestDirection   string
	BusinessID      string
	TrainAmended    string
	CoupledPosition string
	Distance        string
	Route           string
	Orientation     string
	ClosedVehicles  string
}

func NewDlg(record []string) *Dlg {
	dlg := &Dlg{}
	for i, value := range record {
		switch i {
		case 0:
			dlg.Type = value
		case 1:
			dlg.LegType = value
		}
	}
	return dlg
}
