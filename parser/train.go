package parser

import (
	"time"
)

// TrainDate
type TrainDate struct {
	Date time.Time
	Run  bool
}

// CreateDateSequence
func CreateTrainDateSequence(initial time.Time, seq string) []TrainDate {
	var dates []TrainDate
	for i, yes_or_no := range seq {
		date := initial.AddDate(0, 0, i)
		run := false
		if yes_or_no == 'Y' {
			run = true
		}
		ttd := TrainDate{
			Date: date,
			Run:  run,
		}
		dates = append(dates, ttd)
	}
	return dates
}
