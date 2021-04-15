package parser

import (
	"strconv"
	"strings"
	"time"
)

// TimeParse - parse format 'dd-MM-yyyy hh:mm:ss'
func TimeParse(s string) (time.Time, error) {
	a := strings.Split(s, " ")
	b := strings.Split(a[0], "-")
	c := strings.Split(a[1], ":")
	r := append([]string{}, b...)
	r = append(r, c...)
	date := []int{}
	for _, d := range r {
		if n, err := strconv.Atoi(d); err != nil {
			return time.Now(), err
		} else {
			date = append(date, n)
		}
	}
	loc, err := time.LoadLocation("Europe/Oslo")
	if err != nil {
		return time.Now(), err
	}
	month := ItoM(date[1])
	t := time.Date(date[2], month, date[0],
		date[3], date[4], date[5], 0, loc)
	return t, nil
}

// ItoM - translate int to Month
func ItoM(m int) time.Month {
	var months = map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
	return months[m]
}
