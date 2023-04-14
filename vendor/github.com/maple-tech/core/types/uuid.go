package types

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
)

//UUIDReservedRFC4122 is the v4 variant byte needed
const UUIDReservedRFC4122 byte = 0x40

//UUIDHexPattern regular expression for matching a UUIDv4 in hex encoding
const UUIDHexPattern = `^{?([a-f0-9]{8})-([a-f0-9]{4})-([1-5][a-f0-9]{3})-([a-f0-9]{4})-([a-f0-9]{12})}?$`

var regexpUUIDHex = regexp.MustCompile(UUIDHexPattern)

const byteDash = byte('-')

//UUID represents a valid byte variation of a RFC4122 Universally Unique Identifier version 4.
//NOTE: We only support v4 UUID in Maple at this moment
type UUID [16]byte

func (id UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
}

//NewUUID generates and returns a new UUIDv4
func NewUUID() (*UUID, error) {
	//Fill with random
	var data [16]byte
	_, err := rand.Read(data[:])
	if err != nil {
		return nil, err
	}

	id := UUIDFromSeed(data)
	return &id, nil
}

//UUIDFromSeed takes the incoming data (byte array lengthed 16) and conforms it to the UUID pattern
func UUIDFromSeed(data [16]byte) UUID {
	id := UUID(data)

	//Conform to the variation (v4)
	id[8] = (id[8] | UUIDReservedRFC4122) & 0x7F

	//Set bits 12-15 of the version field to the 4 bit version number
	id[6] = (id[6] & 0xF) | (4 << 4)

	return id
}

//UUIDFromHex accepts a byte-slice UUID encoded as hex and returns the UUID object after parsing.
//Accepted formats include with brackets, and without. Should be lower-case.
func UUIDFromHex(str string) (*UUID, error) {
	match := regexpUUIDHex.FindStringSubmatch(str)
	if match == nil {
		return nil, errors.New("invalid UUID")
	}

	//Remove all the formatting and just compile the hex hash
	hash := match[1] + match[2] + match[3] + match[4] + match[5]
	val, err := hex.DecodeString(hash)
	if err != nil {
		return nil, err
	}

	id := new(UUID)
	copy(id[:], val)
	return id, nil
}

var (
	//UUIDUser is a dummy UUID for a known user value.
	//ff000000-0000-4000-4000-000000000000
	UUIDUser = UUIDFromSeed([16]byte{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	//UUIDCompany is a dummy UUID for a known company value.
	//00000000-0000-4000-4000-0000000000ff
	UUIDCompany = UUIDFromSeed([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255})
)
