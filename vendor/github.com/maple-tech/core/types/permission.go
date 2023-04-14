package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"

	"encoding/json"
)

// Permission represents a pseudo-ENUM encapsulation of available permission types
// that is extendable by module's standards
type Permission string

// Constants of used Permission keys in the core system
const (
	PermOwner      Permission = "OWNER"
	PermSuperAdmin Permission = "SUPER_ADMIN"
)

// PermissionArray is a map that get's transformed into a string array for pgsql serialization and JSON.
// It's in map form for convienience of bool checking
type PermissionArray map[Permission]struct{}

// Scan implements the sql.Scanner interface for converting from DB
func (pl *PermissionArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return pl.scanBytes(src)
	case string:
		return pl.scanBytes([]byte(src))
	case nil:
		*pl = nil
		return nil
	}
	return nil
}

// Utility for use with Scan to translate StringArray from pgsql into a map type
func (pl *PermissionArray) scanBytes(src []byte) error {
	elems, err := scanLinearArray(src, []byte{','}, "StringArray")
	if err != nil {
		return err
	}

	b := make(PermissionArray)

	if len(elems) != 0 {
		var prm Permission
		for i, el := range elems {
			if el == nil {
				return fmt.Errorf("pq: parsing array element index %d: cannot convert nil to string", i)
			}
			prm = Permission(string(el))
			b[prm] = struct{}{}
		}
	}

	*pl = b

	return nil
}

// Value implements driver.Value for converting to pgsql StringArray
func (pl PermissionArray) Value() (driver.Value, error) {
	if pl == nil {
		return nil, nil
	}

	if n := len(pl); n > 0 {
		//Not empty so make array for pgsql (formatted in sql '{item,item}')
		b := make([]byte, 1, 1+3*n)
		b[0] = '{' //Start char

		i := 0
		for k := range pl {
			if i == 0 {
				b = appendArrayQuotedBytes(b, []byte(k))
			} else {
				b = append(b, ',')
				b = appendArrayQuotedBytes(b, []byte(k))
			}
			i++
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// MarshalJSON implements json.Marshalling to convert this to JSON string array
func (pl PermissionArray) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteByte('[')

	opening := true
	for k := range pl {
		if !opening {
			b.WriteByte(',')
		}

		b.WriteByte('"')
		b.WriteString(string(k))
		b.WriteByte('"')

		opening = false
	}

	b.WriteByte(']')
	return b.Bytes(), nil
}

// UnmarshalJSON implements json.Marshalling to convert this from JSON string array into a PermissionArray
func (pl *PermissionArray) UnmarshalJSON(data []byte) error {
	b := make(PermissionArray)

	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}

	for _, el := range arr {
		prm := Permission(el)

		b[prm] = struct{}{}
	}

	*pl = b

	return nil
}

//Validate implements our validator.Validate interface for form checking
/*
func (pl PermissionArray) Validate(opts validator.Options) error {
	if opts.Required && len(pl) == 0 {
		return errors.New("permission list is empty")
	}

	for p := range pl {
		if err := p.Validate(opts); err != nil {
			return err
		}
	}

	return nil
}
*/

// Has is a convienience function for checking if permission exists, returns true if so
func (pl PermissionArray) Has(perm Permission) bool {
	_, ok := pl[perm]
	return ok
}

// Add wraps a adding method onto the PermissionArray type
func (pl PermissionArray) Add(perm Permission) {
	pl[perm] = struct{}{}
}

// AddAll adds permissions by variable amount of arguments
func (pl PermissionArray) AddAll(perms ...Permission) {
	for _, prm := range perms {
		pl.Add(prm)
	}
}

// Remove just wraps the delete func for this list
func (pl PermissionArray) Remove(perm Permission) {
	delete(pl, perm)
}

// IsNeutral checks that other is either equal or below our privileges
func (pl PermissionArray) IsNeutral(other PermissionArray) bool {
	for p := range other {
		if !pl.Has(p) {
			return false
		}
	}
	return true
}

// NewPermissionArray creates a new PermissionArray with (optionaly) permissions
func NewPermissionArray(perms ...Permission) PermissionArray {
	pl := make(PermissionArray)

	for _, el := range perms {
		pl.Add(el)
	}

	return pl
}

// PermissionModal holds the record information for a group of permissions given a name/id
// that belongs to a company
type PermissionModal struct {
	ID          ID              `json:"id" db:"perm_id"`
	Name        string          `json:"name" db:"perm_name"`
	Permissions PermissionArray `json:"permissions" db:"perm_perms"`
	TimeCreated time.Time       `json:"timeCreated" db:"perm_created"`
	TimeUpdated NullTime        `json:"timeUpdated,omitempty" db:"perm_updated"`
}

// PermissionModelComplete wraps the PermissionModal and includes the members
// of the model (role).
type PermissionModelComplete struct {
	PermissionModal

	Members []User `json:"members"`
}
