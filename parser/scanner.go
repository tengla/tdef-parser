package parser

import (
	"bufio"
	"os"
	"strings"
)

func Scanner(file *os.File) chan interface{} {

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

			if len(record) < 2 {
				continue
			}

			desc := RecordType.Get(id)

			if len(desc) < 1 {
				continue
			}

			switch id {
			case "PIF":
				ch <- NewPif(record)
			case "BID":
				ch <- NewBid(record)
			case "THD":
				ch <- NewThd(record)
			case "TDT":
				ch <- NewTdt(record)
			case "TSP":
				ch <- NewTsp(record)
			case "TMV":
				ch <- NewTmv(record)
			case "NTE":
				ch <- NewNte(record)
			case "TAS":
				ch <- NewTas(record)
			case "PIT":
				ch <- NewPit(record)
			}
		}
		close(ch)
	}()
	return ch
}
