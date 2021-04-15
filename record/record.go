package record

import (
	tdef "github.com/tengla/tdef-parser/parser"
)

type TdtRecord struct {
	Tdt  *tdef.Tdt
	Tsps []*tdef.Tsp
	Tmvs []*tdef.Tmv
}

type ThdRecord struct {
	Thd        *tdef.Thd
	TdtRecords []TdtRecord `json:"Tdts"`
}

type BidRecord struct {
	Bid        *tdef.Bid
	ThdRecords []ThdRecord `json:"Thds"`
	Ntes       []*tdef.Nte
}
