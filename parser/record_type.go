package parser

var ActionCodes = map[string]string{
	"A": "Add/Update train action",
	"D": "Delete train action",
}

type RecordTypeMap map[string]string

// Get - by key
func (rt RecordTypeMap) Get(k string) string {
	return rt[k]
}

func (rt RecordTypeMap) Has(k string) bool {
	return rt[k] != ""
}

var RecordType = RecordTypeMap{
	"PIF": "TDEF Header",
	"LOC": "Location",
	"PLT": "Platform",
	"RTL": "Route Link",
	"NWK": "Network Link",
	"TLD": "Timing Load",
	"TLK": "Timing Link",
	"REF": "Lookup",
	"PIT": "TDEF Trailer",
	"COD": "General Reference Lookup Combos",
	"VMV": "Valid moves, i.e. connectivitybetween network links and platforms",
	"CMV": "Conflicting Moves",
	"BID": "A business ID",
	"THD": "A train header for each variant of that business ID",
	"TDT": "Each THD record is followed by one or more change-en-route records",
	"TSP": "A stop record",
	"TMV": "A move record",
	"NTE": "A train note",
	"TAS": "Train association",
	"DHR": "A diagram header",
	"DLE": "One or more of these to represent each leg in the diagram",
	"DLG": "A leg within a vehicle diagram",
	"DNE": "Diagram note",
	"D_X": "A diagram trailer record",
	"VST": "Vehicle Set Types",
	"VFM": "Vehicle Formations",
	"VFD": "Each vehicle formation is followed by one or more VFD records to definethe vehicle set type in a coupled position",
	"V_X": "Vehicle Formation Trailer record",
	"PMV": "Permitted moves",
	"DAC": "Diagram activities",
	"DLA": "Location specific diagram activities",
	"DRA": "Resource specific diagram activities",
	"DLR": "Location and resource specific activities",
}
