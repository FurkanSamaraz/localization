package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

//NullString wraps a string value with a valid flag for when a pgsql
//column can be a string, but can also be null.
type NullString struct {
	string
	Valid bool
}

func (ns NullString) String() string {
	if ns.Valid {
		return ns.string
	}
	return ""
}

//StringPtr returns a pointer to the current value, or nil if one is not set
func (ns NullString) StringPtr() *string {
	if ns.Valid {
		return &(ns.string)
	}
	return nil
}

//MarshalJSON allows us to write this value to JSON.
//NOTE: We interpret empty strings as null as well, this may not be true later
//as is an important distinction
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid || len(ns.string) == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, ns.string)), nil
}

//UnmarshalJSON allows us to read from JSON
func (ns *NullString) UnmarshalJSON(src []byte) error {
	str := string(src)

	if strings.EqualFold(str, "null") {
		ns.Valid = false
		return nil
	} else if str[0] == '"' && str[len(str)-2] == '"' {
		ns.string = str[1 : len(str)-2]
		ns.Valid = true
	} else {
		return errors.New("failed to understand the string you provided")
	}
	return nil
}

//Value converts this value to a database usable value
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.string, nil
}

//Scan reads the database value and converts it to one we can use
func (ns *NullString) Scan(value interface{}) error {
	ns.string, ns.Valid = value.(string)
	if !ns.Valid {
		bytes, ok := value.([]byte)
		if ok {
			ns.string = string(bytes)
			ns.Valid = true
			return nil
		}
	}
	return nil
}

//NewNullString creates a new NullString from the provided value,
//will consider the string null if the value is empty
func NewNullString(str string) NullString {
	return NullString{
		string: str,
		Valid:  (len(str) > 0),
	}
}

//NewNullStringFromPtr creates a new NullString from the provided pointer to another string
func NewNullStringFromPtr(str *string) NullString {
	if str == nil {
		return NullString{
			Valid: false,
		}
	}

	return NullString{
		string: *str,
		Valid:  (len(*str) > 0),
	}
}
