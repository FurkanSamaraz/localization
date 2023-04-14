package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
)

//DaysOfWeek provides a data type for representing weekly scheduled "days of the week".
//When a user needs to switch on/off certain days of the week this type provides a
//concise singular type to interact with them.
//It provides SQL & JSON marshaling.
//On the SQL side it is represented as a "bit string" for using binary operators to match/encode with.
//On the JSON side it will encode as an object of enum strings w/ boolean values.
type DaysOfWeek struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}

//ToInt converts the DaysOfWeek into a binary integer.
//This is the format that is kept in the database
func (d DaysOfWeek) ToInt() int {
	val := 0
	if d.Monday {
		val = val | 0b0000001
	}
	if d.Tuesday {
		val = val | 0b0000010
	}
	if d.Wednesday {
		val = val | 0b0000100
	}
	if d.Thursday {
		val = val | 0b0001000
	}
	if d.Friday {
		val = val | 0b0010000
	}
	if d.Saturday {
		val = val | 0b0100000
	}
	if d.Sunday {
		val = val | 0b1000000
	}
	return val
}

//FromInt replaces the data within the DaysOfWeek by
//running a bit-masking routine on the provided integer.
//This representation is the same as the DB uses.
func (d *DaysOfWeek) FromInt(src int) {
	d.Monday = (src & 0b0000001) > 0
	d.Tuesday = (src & 0b0000010) > 0
	d.Wednesday = (src & 0b0000100) > 0
	d.Thursday = (src & 0b0001000) > 0
	d.Friday = (src & 0b0010000) > 0
	d.Saturday = (src & 0b0100000) > 0
	d.Sunday = (src & 0b1000000) > 0
}

//Value satisfies the SQL scanning, converts DaysOfWeek into an integer
func (d DaysOfWeek) Value() (driver.Value, error) {
	return d.ToInt(), nil
}

//Scan satisfies the SQL scanning, converts the incoming number into a DaysOfWeek
func (d *DaysOfWeek) Scan(src interface{}) error {
	var val int
	var ok bool
	val, ok = src.(int)
	if !ok {
		//Figure out the type since something went wonky
		typ := reflect.TypeOf(src)
		if typ == nil {
			return errors.New("could not reflect the incoming type of the sql value for a DaysOfWeek")
		}

		switch typ.String() {
		case "uint8":
			tmp := src.(uint8)
			val = int(tmp)
		case "uint16":
			tmp := src.(uint16)
			val = int(tmp)
		case "uint32":
			tmp := src.(uint32)
			val = int(tmp)
		case "uint64":
			tmp := src.(uint64)
			val = int(tmp)
		case "uint":
			tmp := src.(uint)
			val = int(tmp)
		case "int8":
			tmp := src.(int8)
			val = int(tmp)
		case "int16":
			tmp := src.(int16)
			val = int(tmp)
		case "int32":
			tmp := src.(int32)
			val = int(tmp)
		case "int64":
			tmp := src.(int64)
			val = int(tmp)
		default:
			return fmt.Errorf("failed to scan type '%s' into a types.DaysOfWeek", typ.String())
		}
	}

	d.FromInt(val)
	return nil
}

//DaysOfWeekFromNumber returns a DaysOfWeek object by running a bit-masking routine
//on the provided integer.
func DaysOfWeekFromNumber(src int) DaysOfWeek {
	dow := DaysOfWeek{}
	dow.FromInt(src)
	return dow
}
