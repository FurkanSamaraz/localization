package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"

	"encoding/json"
)

// ErrNoMatchingSchemas error when none of the matching schemas can be unmarshalled without an error
var ErrNoMatchingSchemas = errors.New("no matching schemas")

// UnmarshalJSONWithSchemas EXPERIMENTAL. An attempt to find and parse JSON bodies into
// optional schemas using a variable amount of target objects.
// Returns the index of the schema that matched best, or an error if something failed.
// Accuracy is based on percentage of matching field keys
func UnmarshalJSONWithSchemas(r io.Reader, schemas ...interface{}) (int, error) {
	if len(schemas) == 0 {
		return -1, errors.New("no schemas supplied")
	}

	//Absorb the body for JSON unmarshalling
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return -1, err
	} else if !json.Valid(body) {
		return -1, errors.New("json body is invalid")
	}

	accuracy := make(map[int]float32)
	for i := range schemas {
		//Grab the json field names to use as a best guess of match
		jsonTags := make([]string, 0)

		typ := reflect.TypeOf(schemas[i])
		for j := 0; j < typ.NumField(); j++ {
			fld := typ.Field(j)

			jtag, ok := fld.Tag.Lookup("json")
			if ok {
				if strings.Contains(jtag, ",") {
					jtag = jtag[0:strings.IndexRune(jtag, ',')]
				}

				jsonTags = append(jsonTags, jtag)
			}
			//No json tag, since we ALWAYS use tag names, the field will be ignored
		}

		numTags := len(jsonTags)

		matches := 0
		//Gonna go with string based searching I suppose
		for _, t := range jsonTags {
			//If the tag exists, call it a match I suppose.
			//We disregard whether the TYPE of the value is even a match, because that
			//would just be doing to much
			tag := fmt.Sprintf(`"%s":`, t)
			if bytes.Contains(body, []byte(tag)) {
				matches++
			}
		}

		accuracy[i] = float32(matches) / float32(numTags)
	}

	//Sort the map based on accuracy
	type kv struct {
		Key   int
		Value float32
	}
	var srt []kv
	for k, v := range accuracy {
		srt = append(srt, kv{k, v})
	}

	sort.Slice(srt, func(i, j int) bool {
		return srt[i].Value > srt[j].Value
	})

	//Go through and get the first one that doesn't error.
	//Most cases, this should be the first one
	for i := range schemas {
		err := json.Unmarshal(body, &(schemas[i]))
		if err == nil {
			return i, nil
		}
	}

	return srt[0].Key, ErrNoMatchingSchemas
}
