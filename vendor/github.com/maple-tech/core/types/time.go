package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"
)

//TimeFormat format string for time.Time parsing
const TimeFormat = "15:04:05Z07:00"

//TimeFormatShort format string for time.Time parsing, only including HH:MM
const TimeFormatShort = "15:04"

//Time enacpsulates a time.Time object for only referring to the time portion (date truncated)
//swagger:type time
type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.Format(TimeFormat)
}

//Compare Returns 0 if the same, other wise -1 if other (parameter) is before this, or 1 if other (parameter) is above this
func (t Time) Compare(other Time) int {
	if t.Hour() == other.Hour() {
		if t.Minute() == other.Minute() {
			if t.Second() == other.Second() {
				return 0
			} else if t.Second() < other.Second() {
				return 1
			}
			return -1
		} else if t.Minute() < other.Minute() {
			return 1
		}
		return -1
	} else if t.Hour() < other.Hour() {
		return 1
	}
	return -1
}

//MarshalJSON Implements JSON marshalling
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(TimeFormat) + `"`), nil
}

//UnmarshalJSON Implements JSON marshalling
func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)

	var val time.Time
	var err error
	if len(str) == 10 {
		val, err = time.Parse(TimeFormat, str[1:9]) //Trim quotes
		if err != nil {
			return errors.New("error parsing time-of-day, incorrect format (expected '15:04:05')")
		}
	} else if len(str) == 7 {
		val, err = time.Parse(TimeFormatShort, str[1:6]) //Trim quotes
		if err != nil {
			return errors.New("error parsing time-of-day, incorrect format (expected '15:04')")
		}
	} else {
		return errors.New("error parsing time-of-day string, expected length should be 10")
	}

	t.Time = val
	return nil
}

//Value Implements driver.Value
func (t Time) Value() (driver.Value, error) {
	return t.Format(TimeFormat), nil
}

//Scan Implements sql.Scanner
func (t *Time) Scan(src interface{}) error {
	time, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("expected Time to be time.Time, received %s instead", reflect.TypeOf(src).String())
	}

	t.Time = time
	return nil
}

//TimeDifference Returns a time.Duration comparing a local time and the starting Time
//Does so by converting to time.Time before, so may include dates
func TimeDifference(dt time.Time, start Time) time.Duration {
	return time.Duration(
		(dt.Hour()-start.Hour())*int(time.Hour) +
			(dt.Minute()-start.Minute())*int(time.Minute) +
			(dt.Second()-start.Second())*int(time.Second))
}

//NewTimeFromString Converts a Time formatted string to an object, or returns an error
func NewTimeFromString(str string) (Time, error) {
	tod := Time{}

	val, err := time.Parse(TimeFormat, str)
	if err != nil {
		return tod, err
	}

	tod.Time = val
	return tod, nil
}

//GetTimeFromTime Returns a Time object from a time.Time parameter
//Essentially just a cast
func GetTimeFromTime(dt time.Time) Time {
	t := time.Date(0, 0, 0, dt.Hour(), dt.Minute(), dt.Second(), dt.Nanosecond(), dt.Location()) //Clear out the date
	return Time{t}
}
