package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	file = flag.String("file", "", "The filename")
)

func main() {
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	var b = make([]byte, 120)
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	spl := strings.Split(s, "\t")
	fmt.Println(strings.Join(spl, ","))
}
