package euler

import (
	"time"
)

/**
* Project Euler #19 :
* 
* How many Sundays fell on the first of the month during
* the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

func Euler19() int {
	startDate := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC)
	sundays := 0
	for !startDate.Equal(endDate) {
		day := startDate.Weekday()
		if day == time.Sunday {
			sundays++
		}
		startDate = startDate.AddDate(0, 1, 0)
	}
	return sundays
}
