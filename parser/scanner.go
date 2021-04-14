package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NewRecord - dispatch to correct type
func NewRecord(record []string) interface{} {
	switch record[0] {
	case "PIF":
		return NewPif(record)
	case "BID":
		return NewBid(record)
	case "THD":
		return NewThd(record)
	case "TDT":
		return NewTdt(record)
	case "TSP":
		return NewTsp(record)
	case "TMV":
		return NewTmv(record)
	case "NTE":
		return NewNte(record)
	case "TAS":
		return NewTas(record)
	case "PIT":
		return NewPit(record)
	default:
		return nil
	}
}

func Scanner(file *os.File, dbg string) chan interface{} {

	ch := make(chan interface{})

	go func() {

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			text := scanner.Text()
			record := strings.Split(text, "\t")

			id := record[0]
			if len(id) != 3 {
				continue
			}

			if dbg == id {
				fmt.Printf("%s[%d] %+v\n", id, len(record), record[1:])
			}

			if len(record) < 2 {
				continue
			}

			desc := RecordType.Get(id)

			if len(desc) < 1 {
				continue
			}

			v := NewRecord(record)

			if dbg == id {
				fmt.Printf("%+v\n", v)
			}
			ch <- v
		}
		close(ch)
	}()
	return ch
}
