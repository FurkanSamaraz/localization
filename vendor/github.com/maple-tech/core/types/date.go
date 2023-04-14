package types

import (
	"time"
	"database/sql/driver"
	"fmt"
	"reflect"
)

//DateFormat is the time.Time format string for retieving only the date portion
const DateFormat = "2006-01-02"

//Date wraps a time.Time object with marshaling and scanning for date only conversion.
//This essentially eliminates the time portion.
type Date struct {
	t time.Time
}

//DateInvalid represents an invalid date to test against. It is computed using the largest Unix timestamp possible (1<<63-1)
var DateInvalid = Date{time.Unix(1<<63-1, 0)}

func (d Date) String() string {
	return d.t.Format(DateFormat)
}

//Time returns/exposes the underlining time.Time object
func (d Date) Time() time.Time {
	return d.t
}

//Unix returns the unix timestamp (int64) of the underlining date
func (d Date) Unix() int64 {
	return d.t.Unix()
}

//Day returns the day portion of the date
func (d Date) Day() int {
	return d.t.Day()
}

//Month returns the time.Month portion of the date
func (d Date) Month() time.Month {
	return d.t.Month()
}

//Year returns the year portion of the date
func (d Date) Year() int {
	return d.t.Year()
}

//Set applies a new time.Time object to this Date, will truncate to 0 hour
func (d *Date) Set(t time.Time) {
	//Truncate the value for sanity
	d.t = t.Truncate(24 * time.Hour)
}

//MarshalText turns this into a textual representation. Covers JSON marshalling as well
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

//UnmarshalText reads a date string and converts it
func (d *Date) UnmarshalText(src []byte) error {
	dt, err := ParseDate(string(src))
	if err != nil {
		return err
	}

	*d = dt
	return nil
}

//Value converts this into a SQL value. For PostgreSQL the string representation is fine
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

//Scan converts a SQL value into this
func (d *Date) Scan(src interface{}) error {
	//PQ and the PGSQL drivers take dates as time objects
	time, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("expected Date to be received as time.Time for scanning, received %s instead", reflect.TypeOf(src).String())
	}

	d.t = time
	return nil
}

//NewDate creates a new Date object (set to now as default) and returns it
func NewDate() Date {
	return DateFromTime(time.Now())
}

//ParseDate accepts a date string (YYYY-MM-DD) and returns a Date object
func ParseDate(src string) (Date, error) {
	t, err := time.Parse(DateFormat, src)
	if err != nil {
		return DateInvalid, err
	}
	return Date{t}, nil
}

//DateFromTime takes an incoming time.Time value and returns the Date object for it (truncated)
func DateFromTime(t time.Time) Date {
	d := Date{}
	d.Set(t)
	return d
}