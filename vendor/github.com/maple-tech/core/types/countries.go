package types

import (
	"errors"
	"strings"

	"encoding/json"
)

// CountryCode attempts to encode the ISO standard
// 2 letter country code
type CountryCode string

// Pre-defined country codes. More will be added later
var (
	CountryCodeTurkey        = CountryCode("TR")
	CountryCodeUnitedStates  = CountryCode("US")
	CountryCodeUnitedKingdom = CountryCode("UK")
)

func (c CountryCode) String() string {
	return string(c)
}

// Valid returns true if the country code is 2 characters long
func (c CountryCode) Valid() bool {
	//TODO: Check against actual list of country codes
	return len(c) == 2
}

// MarshalJSON implements JSON marshaling for returning the
// country code string
func (c CountryCode) MarshalJSON() ([]byte, error) {
	if !c.Valid() {
		return nil, errors.New("invalid country code")
	}
	return []byte(`"` + strings.ToUpper(c.String()) + `"`), nil
}

// UnmarshalJSON implements JSON unmarshalling for reading
// incoming JSON strings as country codes
func (c *CountryCode) UnmarshalJSON(src []byte) error {
	if src[0] != '"' || src[len(src)-1] != '"' {
		return errors.New("CountryCode requires a valid string")
	}

	str := string(src[1 : len(src)-2])
	if len(str) != 2 {
		return errors.New("CountryCode is required to be exactly two characters long")
	}

	*c = CountryCode(str)
	return nil
}

// Country holds the information for a global country
// that is retrieved from the database.
//
// This is JSON encoded as an array instead of an object
//
//	["CC", "Name"]
//
// At the moment, I don't care about the flag, or fence
// information.
type Country struct {
	Code CountryCode `db:"cnt_code"`
	Name string      `db:"cnt_name"`
}

// MarshalJSON converts the country into a JSON array
// ["CC", "Name"]
func (c Country) MarshalJSON() ([]byte, error) {
	arr := []string{c.Code.String(), c.Name}
	return json.Marshal(arr)
}

// UnmarshalJSON reads in the JSON array and applies
// it to the Country object
func (c *Country) UnmarshalJSON(src []byte) error {
	arr := make([]string, 2)
	err := json.Unmarshal(src, &arr)
	if err != nil {
		return err
	}

	c.Code = CountryCode(arr[0])
	if c.Code.Valid() == false {
		return errors.New("invalid country code")
	}

	c.Name = arr[1]

	return nil
}
