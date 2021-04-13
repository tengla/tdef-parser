package parser

import (
	"log"
	"strconv"
)

// Pif - Control record (header, file info)
type Pif struct {
	Type               string
	Version            int
	SourceSystem       string
	Operator           string
	TimeTableStartDate string
	TimeTableEndDate   string
	FileCreationDate   string
	FileNumberSequence int
}

// NewPif
func NewPif(record []string) *Pif {
	pif := &Pif{
		Type: record[0],
	}
	if version, err := strconv.Atoi(record[1]); err != nil {
		log.Println(err)
	} else {
		pif.Version = version
	}
	pif.SourceSystem = record[2]
	pif.Operator = record[3]
	pif.TimeTableStartDate = record[4]
	pif.TimeTableEndDate = record[5]
	pif.FileCreationDate = record[6]
	if seq, err := strconv.Atoi(record[7]); err != nil {
		log.Println(err)
	} else {
		pif.FileNumberSequence = seq
	}
	return pif
}
