package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tengla/tdef-parser/parser"
)

var (
	sourcefile = flag.String("file", "", "The tdef file")
)

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
	var pif *parser.Pif
	var tdt *parser.Tdt
	records := parser.Scanner(file)
	for record := range records {
		switch r := record.(type) {
		case *parser.Pif:
			pif = r
		case *parser.Tdt:
			tdt = r
		}
	}
	fmt.Println(*pif)
	fmt.Printf("%+v", *tdt)
}
