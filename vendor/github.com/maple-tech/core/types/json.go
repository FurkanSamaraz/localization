package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"encoding/json"
)

// JSON tricks the marshalling service into writing
// raw json for a value instead of sterilizing it
// or returning a byte string. This is basically
// the same as json.RawMessage except it has SQL
// scanning added to it
type JSON []byte

// JSONNull is just "null" as a JSON object
var JSONNull = JSON("null")

// JSONEmptyObject is just "{}" as a JSON object
var JSONEmptyObject = JSON("{}")

// JSONEmptyArray is just "[]" as a JSON object
var JSONEmptyArray = JSON("[]")

// JSONEmptyString is just "" as a JSON object
var JSONEmptyString = JSON(`""`)

func (j JSON) String() string {
	return string(j)
}

// Equals returns true if the two JSON objects
// match. This is done by converting to a string
// and performing a string equates
func (j JSON) Equals(o JSON) bool {
	return string(j) == string(o)
}

// MarshalJSON returns back the JSON value exactly
// as it is
func (j JSON) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}

	return j, nil
}

// UnmarshalJSON reads the json bytes exactly as they
// are and sets the JSON bytes to be them
func (j *JSON) UnmarshalJSON(src []byte) error {
	*j = append((*j)[0:0], src...)
	return nil
}

// Value returns the JSON as a SQL value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 || j.Equals(JSONNull) {
		return nil, nil
	}
	return []byte(j), nil
}

// Scan reads the SQL value and puts it into
// this JSON object
func (j *JSON) Scan(src interface{}) error {
	var source []byte
	switch st := src.(type) {
	case []byte:
		if len(st) == 0 {
			source = []byte("null")
		} else {
			source = st
		}
	case string:
		source = []byte(st)
	case nil:
		source = []byte("null")
	default:
		return fmt.Errorf("invalid scan type %s for JSON", reflect.TypeOf(st).String())
	}

	*j = append((*j)[0:0], source...)
	return nil
}

// Unmarshal attempts to json.Unmarshal the underlying JSON
// into the object provided
func (j JSON) Unmarshal(obj interface{}) error {
	return json.Unmarshal(j, obj)
}

// MarshalJSON turns an object supplied into a JSON object
// or returns an error if unable to marshal
func MarshalJSON(obj interface{}) (JSON, error) {
	jbytes, err := json.Marshal(obj)
	if err != nil {
		return JSONNull, err
	}

	return JSON(jbytes), nil
}
