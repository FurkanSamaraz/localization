package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/base32"
	"encoding/hex"
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/maple-tech/core/log"
)

// UUIDRegex regular expression for matching (and verifying) a UUID
var UUIDRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`)

// URID wraps a UUID into a solid string with encoders set to format it to-from a more condensed format.
// General input/output format is Base32 string, and a UUID naturally for PostgreSQL
// The internal storage encoding for this is UUID, which itself is a [16]byte representation
//
//swagger:type string
type URID string

// URIDInvalid constant for invalid URIDs, just an empty string at this moment
var URIDInvalid = URID("")

// String returns the UUID as a Base32 encoded string (w/ dashes removed)
func (u URID) String() string {
	if len(u) == 0 {
		//panic("URID is empty, or invalid")
		return ""
	}

	raw := string(u)
	clean := strings.Replace(raw, "-", "", -1)
	enc, err := hex.DecodeString(clean)
	if err != nil {
		panic(err)
	}

	b32 := base32.StdEncoding.WithPadding(base32.NoPadding)
	return b32.EncodeToString(enc)
}

// Bytes convenience wrapper for casting the String() results into a byte slice
func (u URID) Bytes() []byte {
	return []byte(u.String())
}

// IsNull returns true if this value is null/empty
func (u URID) IsNull() bool {
	return len(u) == 0
}

// Equals returns true if the supplied other URID equals this one
func (u URID) Equals(other URID) bool {
	return u == other
}

// Valid returns true if the URID is valid.
func (u URID) Valid() bool {
	return len(u) == 36 && UUIDRegex.MatchString(string(u))
}

// ToUUID returns the URID as a types.UUID pointer, or error if parse failed
func (u URID) ToUUID() (*UUID, error) {
	return UUIDFromHex(string(u))
}

// ToUUIDString returns the URID as a UUID string, or error if one occured
func (u URID) ToUUIDString() (string, error) {
	uid, err := u.ToUUID()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

// FromUUID takes an incoming UUID string and reads it's data into this
func (u *URID) FromUUID(str string) error {
	if !UUIDRegex.MatchString(str) {
		log.Error("URID", "Malformed UUID for URID.FromUUID: ", str)
		return errors.New("URID malformed, input UUID is invalid")
	}
	*u = URID(str)
	return nil
}

// ToBase32 returns the URID as a Base32 encoded string (w/dashes removed).
// Nearly the same as String(), but does not panic, and instead returns the error.
// Makes NO checks on if there's anything in it
func (u URID) ToBase32() (string, error) {
	clean := strings.Replace(string(u), "-", "", -1)
	enc, err := hex.DecodeString(clean)
	if err != nil {
		return "", err
	}

	b32 := base32.StdEncoding.WithPadding(base32.NoPadding)
	return b32.EncodeToString(enc), nil
}

// FromBase32 converts a Base32 string representation of a URID and replaces the value
// of this to be it's decoded counter-part
func (u *URID) FromBase32(str string) error {
	//Decode base32 into hex
	b32 := base32.StdEncoding.WithPadding(base32.NoPadding)
	dec32, err := b32.DecodeString(str)
	if err != nil {
		log.Error("URID", "Failed to decode base32 URID: ", err)
		return err
	}

	//Decode hex into string
	dec16 := hex.EncodeToString(dec32)
	if len(dec16) != 32 {
		log.Error("URID", "Failed to HEX decode base32 URID: ", err)
		return errors.New("URID is malformed, during hex portion of decoding it returned an invalid length")
	}

	var strBuf bytes.Buffer
	for i, char := range dec16 {
		if i == 8 || i == 12 || i == 16 || i == 20 {
			strBuf.WriteByte('-')
		}
		strBuf.WriteRune(char)
	}

	if !UUIDRegex.Match(strBuf.Bytes()) {
		log.Error("URID", "Decoded URID failed Regex match for valid UUID: ", strBuf.String())
		return errors.New("URID is malformed, not a valid UUID input")
	}

	*u = URID(strBuf.String())

	return nil
}

// Clone does a deep-copy of this value and returns it
func (u URID) Clone() URID {
	var n URID
	n = u + ""
	return n
}

// Copy takes another URID and deep-copies it's data into this pointer
func (u *URID) Copy(other URID) {
	clone := other.Clone()
	*u = clone
}

// Generate replaces the data in this object with a new randomly generated UUID value
func (u *URID) Generate() error {
	uid, err := NewUUID()
	if err != nil {
		return err
	}

	*u = URID(uid.String())

	return nil
}

// MarshalText implements Text marshaling
func (u URID) MarshalText() ([]byte, error) {
	return u.Bytes(), nil
}

// UnmarshalText implements Text unmarshaling
func (u *URID) UnmarshalText(data []byte) error {
	return u.FromBase32(string(data))
}

// MarshalJSON implements JSON marshaling.
// This differs from MarshalText in that it returns null if the URID has a length of 0
func (u URID) MarshalJSON() ([]byte, error) {

	if len(u) == 0 {
		return []byte("null"), nil
	}

	if u.Valid() == false {
		return nil, errors.New("URID value was invalid (for JSON marshalling)")
	}

	var buf bytes.Buffer
	buf.WriteByte('"')
	buf.WriteString(u.String())
	buf.WriteByte('"')
	return buf.Bytes(), nil
}

// UnmarshalJSON implements JSON unmarshaling
func (u *URID) UnmarshalJSON(data []byte) error {

	if bytes.Equal(data, []byte("null")) {
		log.Debug("URID", "Encountered null URID value")
		*u = URID("")
		return nil
	}
	if len(data) == 38 {
		*u = URID(string(data[1:37]))
		return nil
	}

	if data[0] != '"' || data[len(data)-1] != '"' {
		log.Error("URID", "Invalid JSON object for URID, expected string, found: ", string(data))
		return errors.New("URID type expected to be a JSON string")
	}

	str := string(data[1 : len(data)-1])
	log.Debug("URID", "Decoding JSON base32 URID: ", str)
	return u.FromBase32(str)
}

// Value returns this URID into a SQL driver.Value
func (u URID) Value() (driver.Value, error) {
	if !u.Valid() {
		log.Error("URID", "Invalid URID for SQL Value encoding: "+string(u))
		return nil, errors.New("URID value was invalid for SQL encoding")
	}

	if len(u) == 0 {
		return nil, nil
	}

	return u.ToUUIDString()
}

// Scan takes an incoming SQL column and converts it into a URID
func (u *URID) Scan(src interface{}) error {
	//Allow nulls I suppose
	if reflect.TypeOf(src) == nil {
		*u = ""
		return nil
	}

	var value string
	switch src.(type) {
	case []byte:
		b := src.([]byte)
		value = string(b)
	case string:
		value = src.(string)
	default:
		return errors.New("expected URID field to be either a []byte, or string. Encountered " + reflect.TypeOf(src).String() + " instead")
	}

	return u.FromUUID(value)
}

/*
//Validate implements validator.Validatable interface
func (u URID) Validate(opts validator.Options) error {
	if !u.Valid() {
		return errors.New("URID value is invalid")
	}

	if opts.Required && u.IsNull() {
		return errors.New("URID is required")
	}

	return nil
}
*/

// NewURID Generates a new URID value by filling it with a randomly generated UUID
func NewURID() (URID, error) {
	var id URID
	err := id.Generate()
	return id, err
}

// ParseURID takes the Base32 string version and parses it into a URID
func ParseURID(str string) (URID, error) {
	var id URID
	err := id.FromBase32(str)
	return id, err
}

// URIDFromUUIDString returns a new URID object filled with a UUID string from input
func URIDFromUUIDString(str string) (URID, error) {
	var id URID
	err := id.FromUUID(str)
	return id, err
}

func unsafeURIDFromUUIDString(str string) URID {
	var id URID
	id.FromUUID(str)
	return id
}

// URIDFromBase32String returns a new URID object filled from a Base32 string.
// NOTE: This is for backwards compatibility and is identical to ParseURID
func URIDFromBase32String(str string) (URID, error) {
	var id URID
	err := id.FromBase32(str)
	return id, err
}

// URIDSliceToSQLArray returns a formatted SQL string for an array of URID's (UUIDs)
func URIDSliceToSQLArray(ids []URID) string {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, id := range ids {
		buf.WriteByte('"')
		str, err := id.ToUUIDString()
		if err != nil {
			panic(err)
		}
		buf.WriteString(str)
		buf.WriteByte('"')

		if i < len(ids)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("}")

	return buf.String()
}

// CloneURID returns a new URID object cloned from another URID object
func CloneURID(other URID) URID {
	return other.Clone()
}

// EmptyURID returns a constant empty URID
var EmptyURID = URID("")

var (
	//URIDUser known URID from the known UUID for an example user
	//74AAAAAAABAAAQAAAAAAAAAAAA
	URIDUser = unsafeURIDFromUUIDString(UUIDUser.String())

	//URIDCompany known URID from the known UUID for an example company
	//AAAAAAAAABAAAQAAAAAAAAAA74
	URIDCompany = unsafeURIDFromUUIDString(UUIDCompany.String())
)
