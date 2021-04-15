package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tengla/tdef-parser/parser"
	"github.com/tengla/tdef-parser/record"
	"github.com/tengla/tdef-parser/serializer"
)

var (
	path = flag.String("path", "", "path to directory")
	id   = flag.String("id", "", "train id")
)

func ls(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range files {
		s := fmt.Sprintf("%s/%s", dir, f.Name())
		list = append(list, s)
	}
	return list
}

func main() {
	flag.Parse()
	if len(*path) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	var bid record.BidRecord
	for _, file := range ls(*path) {
		err := serializer.FromBin(file, &bid)
		if err != nil {
			log.Fatal(err)
		}
		for _, thd := range bid.ThdRecords {
			if len(*id) > 0 && thd.Thd.TrainID == *id {
				t, err := parser.TimeParse(thd.Thd.InitialDateOfOperation)
				if err != nil {
					fmt.Printf("%s\n", err.Error())
					continue
				}
				for _, d := range parser.CreateTrainDateSequence(t, thd.Thd.TrainDates) {
					if d.Run {
						fmt.Printf("%s %s %s %s %s\n", thd.Thd.TrainID,
							thd.Thd.OriginLocation, thd.Thd.DestinationLocation,
							thd.Thd.TrainOwner, d.Date)
					}
				}
			} else if len(*id) < 1 {
				fmt.Printf("%s %s\n", thd.Thd.TrainID, thd.Thd.InitialDateOfOperation)
			}
		}
	}
}
