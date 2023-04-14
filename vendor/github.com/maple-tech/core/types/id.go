package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ID represents a simple integer ID for a resource. Generally comes from PostgreSQL bigint/bigserial
// data types. For this reason, it's just a mask of a int64. The string representation is Hexidecimal.
type ID int64

// IDInvalid holds the simple representation of an invalid ID, or -1 as it is internally
const IDInvalid = ID(-1)

func (i ID) String() string {
	return strconv.FormatInt(int64(i), 16)
}

// Valid returns true if the internal value is valid
func (i ID) Valid() bool {
	return i > 0
}

// Value converts this value to SQL usable data type
func (i ID) Value() (driver.Value, error) {
	return int64(i), nil
}

// Scan converts from a SQL data type, into this type
func (i *ID) Scan(src interface{}) error {

	var val int64
	var ok bool
	val, ok = src.(int64)
	if !ok {
		//Figure out the type since something went wonky
		typ := reflect.TypeOf(src)
		if typ == nil {
			*i = IDInvalid
			return nil
		}

		switch typ.String() {
		case "int32":
			tmp := src.(int32)
			val = int64(tmp)
		case "int":
			tmp := src.(int)
			val = int64(tmp)
		default:
			return fmt.Errorf("failed to scan type '%s' into a types.ID", typ.String())
		}
	}

	*i = ID(val)

	return nil
}

// MarshalText converts this value into a string representation.
// This covers MarshalJSON as well.
// The simple ID.String() function is used, which returns it as Hex encoded
func (i ID) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalJSON converts from JSON data into this ID object.
// Hex string is the prefered type, but we will accept a base10 integer if we can coerse if validly
func (i *ID) UnmarshalJSON(src []byte) error {

	//Try to coerse it into a string first, since this is the default encoding type anyways
	if len(src) > 2 && src[0] == '"' && src[len(src)-1] == '"' {
		cont := src[1 : len(src)-1]
		if len(cont) > 0 {
			val, err := strconv.ParseInt(string(cont), 16, 64)
			if err != nil {
				return err
			}
			*i = ID(val)
			return nil
		}

		return errors.New("parsing JSON value for ID resulted in an empty string")
	} else if len(src) > 0 {
		//I guess we can accept integers as a whole if they can pass the test
		val, err := strconv.ParseInt(string(src), 16, 64)
		if err != nil {
			return err
		} else if val < -1 {
			return errors.New("failed to parse JSON integer as ID, resulted in a negative number")
		}

		*i = ID(val)

		return nil
	}

	return errors.New("cannot parse ID from empty JSON data")
}

// ParseID converts a hex string into an ID
func ParseID(str string) (ID, error) {
	val, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		return IDInvalid, err
	}

	return ID(val), nil
}

// NullID wraps an ID type with a valid flag for usage with
// PostgreSQl and PQ. When encoded, they ID will reflect normally,
// or be "null" in the json body if invalid.
type NullID struct {
	ID
	Valid bool
}

func (ni NullID) String() string {
	if ni.Valid {
		return ni.ID.String()
	}
	return ""
}

// Ptr returns the underlying ID as a pointer if available, otherwise nil
func (ni NullID) Ptr() *ID {
	if ni.Valid {
		return &(ni.ID)
	}
	return nil
}

// MarshalJSON encodes the value as JSON. If the value is nil then
// null is written, otherwise the ID is string encoded.
func (ni NullID) MarshalJSON() ([]byte, error) {
	if !ni.Valid || ni.ID == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, ni.ID.String())), nil
}

// UnmarshalJSON reads JSON in and converts to a NullID
func (ni *NullID) UnmarshalJSON(src []byte) error {

	str := string(src)
	if strings.EqualFold(str, "null") {
		ni.Valid = false
		ni.ID = 0
		return nil
	} else if str[0] == '"' && str[len(str)-2] == '"' {
		var err error
		ni.ID, err = ParseID(str[1 : len(str)-2])
		if err != nil {
			return err
		}
		ni.Valid = true
		return nil
	} else {
		return errors.New("failed to understand the ID you provided")
	}
}

// Value converts this ID to a proper PGSQL value.
func (ni NullID) Value() (driver.Value, error) {
	if !ni.Valid || ni.ID == 0 {
		return nil, nil
	}

	return ni.ID, nil
}

// Scan reads in the PGSQL value and converts it to a Nullable ID
func (ni *NullID) Scan(value interface{}) error {

	id, ok := value.(int64)
	if !ok {
		ni.ID = 0
		ni.Valid = false
		return nil
	}
	ni.ID = ID(id)
	ni.Valid = ni.ID > 0
	return nil
}

// NullIDFromValue converts an integer into a NullID type.
// If the ID value is 0 or below, the NullID is considered nil (null)
func NullIDFromValue(id int64) NullID {
	return NullID{
		ID:    ID(id),
		Valid: id > 0,
	}
}

// NullIDFromID converts an ID into a NullID.
// If the ID value is 0 or below, the NullID is considered nil (null)
func NullIDFromID(id ID) NullID {
	return NullID{
		ID:    id,
		Valid: id > 0,
	}
}

// IDFromNullID converts a NullID into an ID.
// If the NullID is invalid (null), then an ID of 0 is returned
func IDFromNullID(nullID NullID) ID {
	return nullID.ID
}
