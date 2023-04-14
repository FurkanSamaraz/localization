package utils

import (
	"math"
	"time"
)

//DateEqual returns true if the supplied time.Time parameters are on the same day
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

//TimeToDate rounds a time.Time object to just be the midnight of that day
func TimeToDate(toRound time.Time) time.Time {
	return time.Date(toRound.Year(), toRound.Month(), toRound.Day(), 0, 0, 0, 0, toRound.Location())
}

//DaysBetweenDates returns the number of days betwen two specific dates
//Note: Does not min/max the dates, so date2 should be the higher of the two for positive results
func DaysBetweenDates(date1, date2 time.Time) int {
	t1 := TimeToDate(date1)
	t2 := TimeToDate(date2)
	days := math.Ceil(t2.Sub(t1).Hours()/24) + 1
	return int(days)
}
