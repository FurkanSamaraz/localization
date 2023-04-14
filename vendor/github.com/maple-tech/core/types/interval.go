package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//TimeInterval Holds duration/interval information received from PostgreSQL,
//supports each format except verbose. JSON marshals to a clock format "HH:MM:SS".
//Uses a Valid property so it can be nullable
type TimeInterval struct {
	Valid   bool
	Hours   uint8
	Minutes uint8
	Seconds uint8
}

//Value Implements driver.Value
func (t TimeInterval) Value() (driver.Value, error) {
	if t.Hours >= 24 {
		return nil, errors.New("TimeInterval does not support intervals over 1 day (24 hours)")
	}

	return fmt.Sprintf("P0Y0M0DT%dH%dM%dS", t.Hours, t.Minutes, t.Seconds), nil
}

//Scan Implements sql.Scanner
func (t *TimeInterval) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		//Probably nil
		t.Valid = false
		t.Hours = 0
		t.Minutes = 0
		t.Seconds = 0
		return nil
	}

	str := strings.ToUpper(string(bytes))
	if len(str) == 0 {
		return errors.New("received bytes for TimeInterval but string ended up empty")
	}

	if str[0] == 'P' { //ISO 8601
		timePortion := str[strings.Index(str, "T")+1:]
		indHr := strings.Index(timePortion, "H")
		if indHr > 0 {
			hours, err := strconv.ParseUint(timePortion[0:indHr], 10, 32)
			if err != nil {
				return err
			}

			t.Hours = uint8(hours)
		}

		indMin := strings.Index(timePortion, "M")
		if indMin > 0 {
			startInd := indHr + 1
			minutes, err := strconv.ParseUint(timePortion[startInd:indMin], 10, 32)
			if err != nil {
				return err
			}

			t.Minutes = uint8(minutes)
		}

		indSec := strings.Index(timePortion, "S")
		if indSec > 0 {
			startInd := indMin + 1
			seconds, err := strconv.ParseUint(timePortion[startInd:indSec], 10, 32)
			if err != nil {
				return err
			}

			t.Minutes = uint8(seconds)
		}
	} else { //SQL Standard (I think), and Postgres (since only time)
		indHr := strings.Index(str, ":")
		hours, err := strconv.ParseUint(str[0:indHr], 10, 32)
		if err != nil {
			return err
		}
		t.Hours = uint8(hours)

		minutes, err := strconv.ParseUint(str[indHr+1:indHr+3], 10, 32)
		if err != nil {
			return err
		}
		t.Minutes = uint8(minutes)

		seconds, err := strconv.ParseUint(str[len(str)-2:], 10, 32)
		if err != nil {
			return err
		}
		t.Seconds = uint8(seconds)
	}
	t.Valid = true

	return nil
}

//MarshalJSON Marshals JSON
func (t TimeInterval) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return []byte(fmt.Sprintf("\"%02d:%02d:%02d\"", t.Hours, t.Minutes, t.Seconds)), nil
	}

	return []byte("null"), nil
}

//ToSeconds Returns the culminative seconds of this interval
func (t TimeInterval) ToSeconds() uint {
	return uint(t.Hours*60*60) + uint(t.Minutes*60) + uint(t.Seconds)
}

//UnmarshalJSON Implements JSON marshalling
func (t *TimeInterval) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str[0] != '"' {
		if str == "null" {
			t.Valid = false
			return nil
		}

		return fmt.Errorf("expected TimeInterval to be a string, received %s instead", str)
	}

	if len(str) != 10 {
		return fmt.Errorf("expected TimeInterval to be '00:00:00', received %s instead", str)
	}

	str = str[1 : len(str)-1]
	if len(str) == 0 {
		t.Hours = 0
		t.Minutes = 0
		t.Seconds = 0
		t.Valid = false
		return nil
	}

	indHr := strings.Index(str, ":")
	hours, err := strconv.ParseUint(str[:indHr], 10, 32)
	if err != nil {
		return errors.New("failed to parse the hours section of the time interval")
	}
	t.Hours = uint8(hours)

	minutes, err := strconv.ParseUint(str[indHr+1:indHr+3], 10, 32)
	if err != nil {
		return errors.New("failed to parse the minutes section of the time interval")
	}
	t.Minutes = uint8(minutes)

	seconds, err := strconv.ParseUint(str[len(str)-2:], 10, 32)
	if err != nil {
		return errors.New("failed to parse the seconds section of the time interval")
	}
	t.Seconds = uint8(seconds)

	t.Valid = true

	return nil
}

//NewTimeIntervalFromDuration Converts a duration, into a TimeInterval object
func NewTimeIntervalFromDuration(d time.Duration) TimeInterval {
	hrs := uint8(math.Floor(d.Hours()))
	mins := uint8(math.Floor(d.Minutes()) - float64(hrs*60))
	secs := uint8(math.Floor(d.Seconds()) - float64(mins*60) - float64(hrs*60*60))
	return TimeInterval{
		Valid:   true,
		Hours:   hrs,
		Minutes: mins,
		Seconds: secs,
	}
}
