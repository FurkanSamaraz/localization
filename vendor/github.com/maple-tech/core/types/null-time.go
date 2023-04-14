package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// NullTime wraps a time.Time object with a valid flag for when a pgsql
// column type is temporal, but nullable.
// This is needed since Go doesn't seem to have native nil time.Time support,
// and it's kinda good for specifying that a column is expected/allowed to be null.
//
//swagger:type datetime
type NullTime struct {
	time.Time
	Valid bool
}

func (nt NullTime) String() string {
	return nt.Time.String()
}

// MarshalJSON allows for JSON marshalling (setting values to JSON)
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, nt.Time.Format(time.RFC3339))), nil
}

// UnmarshalJSON allows for JSON unmarshalling (getting values from JSON)
func (nt *NullTime) UnmarshalJSON(src []byte) error {
	str := string(src)

	if strings.EqualFold(str, "null") {
		nt.Valid = false
	} else if str[0] == '"' && str[len(str)-2] == '"' {
		timeStr := str[1 : len(str)-2]
		t, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			return errors.New("failed to parse date/time provided, must be a valid RFC3339 date/time")
		}

		nt.Time = t
		nt.Valid = true
	} else if input := strings.Replace(str, "\"", "", -1); len(input) > 1 {
		t, err := time.Parse(time.RFC3339, input)
		if err != nil {
			return errors.New("failed to parse date/time provided, must be a valid RFC3339 date/time")
		}

		nt.Time = t
		nt.Valid = true
	} else {
		return errors.New("could not understand the date/time value provided")
	}

	return nil
}

// Value allows for conversion to the SQL data type
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	//time.Time already has SQL interfaces for it thanks to PQ and the like
	return nt.Time, nil
}

// Scan allows SQL to take a database column and convert it to this type
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// NewEmptyNullTime returns an empty NullTime value that represents a null in the eyes of the database
func NewEmptyNullTime() NullTime {
	return NullTime{}
}

// NewNullTime returns a new NullTime object from a given time.Time object, this will be valid to the database
func NewNullTime(t time.Time) NullTime {
	return NullTime{
		Time:  t,
		Valid: true,
	}
}
