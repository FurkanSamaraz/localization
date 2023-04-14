package types

import (
	"errors"
	"strings"
)

//RegionCode wraps the ISO standard regional code
//(5 characters long).
//The format is [CC]-[ID] with CC being the country code
type RegionCode string

func (r RegionCode) String() string {
	return string(r)
}

//Valid returns true if the RegionCode
//is in a valid format
func (r RegionCode) Valid() bool {
	return len(r) == 5
}

//GetCountryCode extracts the 2 leter country code
//from the RegionCode. This will panic if the
//value is invalid.
func (r RegionCode) GetCountryCode() CountryCode {
	if !r.Valid() {
		panic("invalid region code")
	}

	return CountryCode(string(r[:2]))
}

//GetRegionID extracts the 2 letter region code
//from the RegionCode. This will panic if the
//value is invalid.
func (r RegionCode) GetRegionID() string {
	if !r.Valid() {
		panic("invalid region code")
	}

	return string(r[4:])
}

//MarshalJSON converts the region code into a JSON string
func (r RegionCode) MarshalJSON() ([]byte, error) {
	if !r.Valid() {
		return nil, errors.New("invalid region code")
	}

	return []byte(`"` + strings.ToUpper(r.String()) + `"`), nil
}

//UnmarshalJSON reads in a JSON strings and converts it to
//the RegionCode
func (r *RegionCode) UnmarshalJSON(src []byte) error {
	if src[0] != '"' || src[len(src)-1] != '"' {
		return errors.New("RegionCode requires a valid string")
	}

	str := string(src[1 : len(src)-2])
	if len(str) != 5 {
		return errors.New("RegionCode is required to be exactly five characters long")
	} else if str[3] != '-' {
		return errors.New("RegionCode is in the wrong format")
	}

	*r = RegionCode(str)
	return nil
}

//Region contains the information about a region within a country.
//Represented by it's ISO code, and the given name.
//
//For the moment I am not loading in the gps fence. Also since
//the country is extractable from the code, I don't carry the
//database column that joins the country, just use GetCountryCode()
type Region struct {
	Code RegionCode `db:"rgn_code"`
	Name string     `db:"rgn_name"`
}
