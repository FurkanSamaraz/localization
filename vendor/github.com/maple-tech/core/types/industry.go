package types

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"encoding/json"
)

// Industry is an alias type of uint16 for enum purposes
type Industry uint16

// Industry index identifiers for use with the Industries map
const (
	IndNone Industry = iota
	IndOther
	IndAccommodations
	IndAccounting
	IndAdvertising
	IndAerospace
	IndAgriculture
	IndApparel
	IndAutomotive
	IndBanking
	IndBiotechnology
	IndBusiness
	IndChemical
	IndCommunications
	IndComputers
	IndConstruction
	IndConsulting
	IndConsumer
	IndCosmetics
	IndEducation
	IndElectronics
	IndEmployment
	IndEnergy
	IndEntertainment
	IndFashion
	IndFinancial
	IndFood
	IndGovernment
	IndHealthcare
	IndInformation
	IndInformationTechnology
	IndInsurance
	IndJournalism
	IndLegal
	IndManufacturing
	IndMarketing
	IndMedia
	IndMusic
	IndPharmaceutical
	IndPublicAdministration
	IndPublicTransportation
	IndPublicRelations
	IndPublishing
	IndRealEstate
	IndRetail
	IndService
	IndSports
	IndTechnology
	IndTelecommunications
	IndTourism
	IndTransportation
	IndUtilities
	IndVideoGames
	IndWebServices
)

// Industries maps the index from Constants Ind* to their string key
var Industries = []string{"NONE", "OTHER",
	"ACCOMMODATIONS", "ACCOUNTING", "ADVERTISING", "AEROSPACE", "AGRICULTURE", "APPAREL", "AUTOMOTIVE",
	"BANKING", "BIOTECHNOLOGY", "BUSINESS",
	"CHEMICAL", "COMMUNICATIONS", "COMPUTERS", "CONSTRUCTION", "CONSULTING", "CONSUMER", "COSMETICS",
	"EDUCATION", "ELECTRONICS", "EMPLOYMENT", "ENERGY", "ENTERTAINMENT",
	"FASHION", "FINANCIAL", "FOOD",
	"GOVERNMENT",
	"HEALTHCARE",
	"INFORMATION", "INFORMATION_TECHNOLOGY", "INSURANCE",
	"JOURNALISM",
	"LEGAL",
	"MANUFACTURING", "MARKETING", "MEDIA", "MUSIC",
	"PHARMACEUTICAL", "PUBLIC_ADMINISTRATION", "PUBLIC_TRANSPORTATION", "PUBLIC_RELATIONS", "PUBLISHING",
	"REAL_ESTATE", "RETAIL",
	"SERVICE", "SPORTS",
	"TECHNOLOGY", "TELECOMMUNICATIONS", "TOURISM", "TRANSPORTATION",
	"UTILITIES",
	"VIDEO_GAMES",
	"WEB_SERVICES",
}

// IndustryFromString returns a new industry (or error if not found) for a given string constant
func IndustryFromString(str string) (Industry, error) {
	for j, o := range Industries {
		if strings.EqualFold(str, o) {
			ind := Industry(j)
			return ind, nil
		}
	}

	return 0, fmt.Errorf("String value '%s' did not equate to a constant we accept", str)
}

// String returns the string representation of the industry by index
func (i Industry) String() string {
	if int(i) >= len(Industries) {
		panic("Industry value is out of range of acceptable industry types")
	}
	return Industries[i]
}

// Value implements driver.Value for converting this value to a SQL column type
func (i Industry) Value() (driver.Value, error) {
	if int(i) >= len(Industries) {
		return nil, errors.New("Industry value is out of range of acceptable industry types")
	}
	return i.String(), nil
}

// Scan implements sql.Scanner for converting a SQL column to an Industry type
func (i *Industry) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("Could not convert value to the expected string type")
	}

	ind, err := IndustryFromString(str)
	if err != nil {
		return err
	}
	i = &ind

	return nil
}

// MarshalJSON implements JSON marshaling interface for converting Industry -> String
func (i Industry) MarshalJSON() ([]byte, error) {
	if int(i) >= len(Industries) {
		return nil, errors.New("Industry value is out of range of acceptable industry types")
	}

	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(Industries[i])
	b.WriteByte('"')
	return b.Bytes(), nil
}

// UnmarshalJSON implements JSON marshaling interface for converting String -> Industry
func (i *Industry) UnmarshalJSON(data []byte) error {
	ind, err := IndustryFromString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	i = &ind
	return nil
}

// IndustryArray wrapper of a slice of Industry objects for marshalling
type IndustryArray []Industry

// ToStringSlice converts the IndustryArray slice into a string slice
func (a IndustryArray) ToStringSlice() []string {
	strs := make([]string, len(a))
	for j := range a {
		strs[j] = a[j].String()
	}
	return strs
}

// MarshalJSON converts an IndustryArray into a JSON string array
func (a IndustryArray) MarshalJSON() ([]byte, error) {
	strArr := a.ToStringSlice()
	return json.Marshal(strArr)
}

// UnmarshalJSON converts a JSON string array into a IndustryArray and validates the industries
func (a *IndustryArray) UnmarshalJSON(data []byte) error {
	strArr := make([]string, 0)
	err := json.Unmarshal(data, &strArr)
	if err != nil {
		return fmt.Errorf("failed to pre-parse types.IndustryArray into string array")
	}

	*a = make(IndustryArray, len(strArr))
	for i := range strArr {
		ind, e := IndustryFromString(strArr[i])
		if e != nil {
			return e
		}
		(*a)[i] = ind
	}

	return nil
}

// Scan implements the sql.Scanner interface.
func (a *IndustryArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to IndustryArray", src)
}

func (a *IndustryArray) scanBytes(src []byte) error {
	elems, err := scanLinearArray(src, []byte{','}, "StringArray")
	if err != nil {
		return err
	}
	if *a != nil && len(elems) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(IndustryArray, len(elems))
		var ind Industry
		for i, v := range elems {
			if v == nil {
				return fmt.Errorf("pq: parsing array element index %d: cannot convert nil to string", i)
			}
			ind, err = IndustryFromString(string(v))
			if err != nil {
				return fmt.Errorf("pq: parsing array element index %d: cannot convert to industry string - %v", i, err)
			}
			b[i] = ind
		}
		*a = b
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (a IndustryArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(a[0].String()))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(a[i].String()))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

func appendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}

func parseArray(src, del []byte) (dims []int, elems [][]byte, err error) {
	var depth, i int

	if len(src) < 1 || src[0] != '{' {
		return nil, nil, fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '{', 0)
	}

Open:
	for i < len(src) {
		switch src[i] {
		case '{':
			depth++
			i++
		case '}':
			elems = make([][]byte, 0)
			goto Close
		default:
			break Open
		}
	}
	dims = make([]int, i)

Element:
	for i < len(src) {
		switch src[i] {
		case '{':
			if depth == len(dims) {
				break Element
			}
			depth++
			dims[depth-1] = 0
			i++
		case '"':
			var elem = []byte{}
			var escape bool
			for i++; i < len(src); i++ {
				if escape {
					elem = append(elem, src[i])
					escape = false
				} else {
					switch src[i] {
					default:
						elem = append(elem, src[i])
					case '\\':
						escape = true
					case '"':
						elems = append(elems, elem)
						i++
						break Element
					}
				}
			}
		default:
			for start := i; i < len(src); i++ {
				if bytes.HasPrefix(src[i:], del) || src[i] == '}' {
					elem := src[start:i]
					if len(elem) == 0 {
						return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
					}
					if bytes.Equal(elem, []byte("NULL")) {
						elem = nil
					}
					elems = append(elems, elem)
					break Element
				}
			}
		}
	}

	for i < len(src) {
		if bytes.HasPrefix(src[i:], del) && depth > 0 {
			dims[depth-1]++
			i += len(del)
			goto Element
		} else if src[i] == '}' && depth > 0 {
			dims[depth-1]++
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}

Close:
	for i < len(src) {
		if src[i] == '}' && depth > 0 {
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}
	if depth > 0 {
		err = fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '}', i)
	}
	if err == nil {
		for _, d := range dims {
			if (len(elems) % d) != 0 {
				err = fmt.Errorf("pq: multidimensional arrays must have elements with matching dimensions")
			}
		}
	}
	return
}

func scanLinearArray(src, del []byte, typ string) (elems [][]byte, err error) {
	dims, elems, err := parseArray(src, del)
	if err != nil {
		return nil, err
	}
	if len(dims) > 1 {
		return nil, fmt.Errorf("pq: cannot convert ARRAY%s to %s", strings.Replace(fmt.Sprint(dims), " ", "][", -1), typ)
	}
	return elems, err
}
