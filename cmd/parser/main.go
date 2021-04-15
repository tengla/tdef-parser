package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	tdef "github.com/tengla/tdef-parser/parser"
	"github.com/tengla/tdef-parser/serializer"
)

var (
	sourcefile = flag.String("file", "", "The tdef file")
	print      = flag.Bool("print", false, "print results")
	serialize  = flag.String("serialize", "", "dump binary records to directory")
	dbg        = flag.String("dbg", "", "debug record type")
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

func (b *BidRecord) AppendTdtRecord(t TdtRecord) {
	b.ThdRecords[len(b.ThdRecords)-1].TdtRecords =
		append(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords, t)
}

func (b *BidRecord) AppendTspRecord(t *tdef.Tsp) {
	b.ThdRecords[len(b.ThdRecords)-1].TdtRecords[len(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords)-1].Tsps =
		append(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords[len(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords)-1].Tsps, t)
}

func (b *BidRecord) AppendTmvRecord(t *tdef.Tmv) {
	b.ThdRecords[len(b.ThdRecords)-1].TdtRecords[len(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords)-1].Tmvs =
		append(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords[len(b.ThdRecords[len(b.ThdRecords)-1].TdtRecords)-1].Tmvs, t)
}

func main() {
	flag.Parse()
	if len(*sourcefile) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	file, err := os.Open(*sourcefile)
	if err != nil {
		log.Fatal(err.Error())
	}
	records := tdef.Scanner(file, *dbg)
	var bid BidRecord
	records_count := 0
	for record := range records {
		switch r := record.(type) {
		case *tdef.Bid:
			// we have a fully compiled record
			if bid.Bid != nil && *print {
				data, err := json.Marshal(bid)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(string(data))
				}
			}
			if bid.Bid != nil && len(*serialize) > 0 {
				dat := fmt.Sprintf("%s/dat%06d.dat", *serialize, records_count)
				serializer.ToBin(dat, bid)
			}
			records_count += 1
			// now start a new
			bid = BidRecord{}
			bid.Bid = r
		case *tdef.Thd:
			thd := ThdRecord{}
			thd.Thd = r
			bid.ThdRecords = append(bid.ThdRecords, thd)
		case *tdef.Tdt:
			tdt := TdtRecord{}
			tdt.Tdt = r
			bid.AppendTdtRecord(tdt)
		case *tdef.Tsp:
			bid.AppendTspRecord(r)
		case *tdef.Tmv:
			bid.AppendTmvRecord(r)
		case *tdef.Nte:
			bid.Ntes = append(bid.Ntes, r)
		}
	}
}
