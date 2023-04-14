package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

//ContactType [enum] flags the type that a Contact object contains
type ContactType byte

const (
	//ContactTypeEmail flags the ContactType as an email address
	ContactTypeEmail ContactType = iota

	//ContactTypePhone flags the ContactType as a phone number
	ContactTypePhone
)

func (ct ContactType) String() string {
	switch ct {
	case ContactTypeEmail:
		return "EMAIL"
	case ContactTypePhone:
		return "PHONE"
	}
	return "UNKNOWN"
}

var regexpPhoneClean = regexp.MustCompile(`[^\+\d]`)
var regexpPhoneTest = regexp.MustCompile(`^\+\d{12}$`)

//Contact type-agnostic contact information. Allows for either email/phone to be entered
//and parsed.
type Contact struct {
	Type ContactType
	Data string
}

func (c Contact) String() string {
	return c.Data
}

//Valid returns true if the underlining data is meets type validation
func (c Contact) Valid() bool {
	switch c.Type {
	case ContactTypeEmail:
		return len(c.Data) >= 3 && strings.ContainsRune(c.Data, '@')
	case ContactTypePhone:
		return len(c.Data) == 13 && regexpPhoneTest.MatchString(c.Data)
	}
	return false
}

//Set replaces the data in the contact with the provided value.
//Attempts to find the type and clean the value as well.
//This is prefered instead of just setting the fields as it cleans and type checks it.
func (c *Contact) Set(value string) {
	if strings.ContainsRune(value, '@') { //Email
		c.Type = ContactTypeEmail
		value = strings.TrimSpace(value)
	} else {
		c.Type = ContactTypePhone
		value = regexpPhoneClean.ReplaceAllString(value, "")
	}

	c.Data = value
}

//Value used by SQL to turn this into a SQL value
func (c Contact) Value() (driver.Value, error) {
	return c.String(), nil
}

//Scan used by SQL to turn a column into a Contact
func (c *Contact) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		typ := reflect.TypeOf(src)
		switch typ.String() {
		case "[]uint8", "[]byte":
			barr, ok := src.([]byte)
			if !ok {
				return fmt.Errorf("failed to coerce value during scan, expected string found %s instead", typ.String())
			}
			str = string(barr)
		default:
			return fmt.Errorf("failed to coerce value during scan, expected string found %s instead", typ.String())
		}
	}

	c.Set(str)
	return nil
}

//MarshalText turns the Contact into a string representation
func (c Contact) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

//UnmarshalText converts incoming text into a Contact object
func (c *Contact) UnmarshalText(data []byte) error {
	str := string(data)
	c.Set(str)
	return nil
}

//NewContact creates a new Contact object with the specified value.
//Will type check as well so Contact.Type should be valid
func NewContact(value string) Contact {
	var c Contact
	c.Set(value)
	return c
}
