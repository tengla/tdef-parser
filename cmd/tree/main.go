package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tdef "github.com/tengla/tdef-parser/parser"
)

// Alternative data structure. Create tree

var (
	file = flag.String("file", "", "open file")
	//typ  = flag.String("type", "", "record type")
)

type RecordType []string

var ValidTypes = RecordType{
	"PIF", "BID", "THD", "TDT", "TSP", "TMV", "NTE", "TAS", "PIT",
}

func (rt RecordType) Has(key string) bool {
	for _, t := range rt {
		if t == key {
			return true
		}
	}
	return false
}

type Record struct {
	Fields   []string
	Parent   *Record
	Children []*Record
}

func (r *Record) GetInstance() interface{} {
	return tdef.NewRecord(r.Fields)
}

func (r *Record) GetType() string {
	return r.Fields[0]
}

func NewRecord(cells []string) *Record {
	return &Record{
		Fields: cells,
	}
}

func (el *Record) AddRecord(record *Record) *Record {

	if record.GetType() == "PIF" || record.GetType() == "PIT" {
		record.Parent = el
		el.Children = append(el.Children, record)
		return el
	}
	if record.GetType() == "BID" {
		record.Parent = el.Children[0]
		el.Children[0].Children = append(el.Children[0].Children, record)
		return el
	}
	if record.GetType() == "THD" {
		lastBid := el.Children[0].Children[len(el.Children[0].Children)-1]
		record.Parent = lastBid
		lastBid.Children = append(lastBid.Children, record)
		return el
	}
	if record.GetType() == "TDT" {
		lastBid := el.Children[0].Children[len(el.Children[0].Children)-1]
		lastThd := lastBid.Children[len(lastBid.Children)-1]
		record.Parent = lastThd
		lastThd.Children = append(lastThd.Children, record)
		return el
	}
	if record.GetType() == "TSP" || record.GetType() == "TMV" {
		lastBid := el.Children[0].Children[len(el.Children[0].Children)-1]
		lastThd := lastBid.Children[len(lastBid.Children)-1]
		lastTdt := lastThd.Children[len(lastThd.Children)-1]
		record.Parent = lastTdt
		lastTdt.Children = append(lastTdt.Children, record)
	}
	return el
}

func (el *Record) Traverse(depth int, f func(int, *Record)) {
	f(depth, el)
	for _, child := range el.Children {
		child.Traverse(depth+1, f)
	}
}

func main() {
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	scanner := bufio.NewScanner(f)
	root := &Record{Fields: []string{"ROOT"}}
	for scanner.Scan() {
		text := scanner.Text()
		cells := strings.Split(text, "\t")
		if !ValidTypes.Has(cells[0]) {
			continue
		}
		// persist record in tree
		root.AddRecord(&Record{
			Fields: cells,
		})
	}
	tabs := func(n int) string {
		return strings.Join(make([]string, n), "\t")
	}
	root.Traverse(0, func(depth int, r *Record) {
		data := r.GetInstance()
		switch t := data.(type) {
		case *tdef.Thd:
			if len(t.TrainOwner) < 1 {
				fmt.Printf("%s%s %s\n", tabs(depth), t.TrainID, t.TrainOwner)
			}
		case *tdef.Pit:
			fmt.Print("Changes: ")
			for _, change := range t.Changes {
				if n, err := strconv.Atoi(change.AdditionCount); err != nil {
					fmt.Println(err)
				} else if n > 0 {
					fmt.Printf("%+v ", *change)
				}
			}
			fmt.Println("")
		}
	})
}
