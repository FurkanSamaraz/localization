package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/maple-tech/core/utils"
)

// LoadEnv takes an incoming options pointer and applies settings values based on environment variables.
// The environment variables are expected to override the file based options. The name of the keys is derived
// from the JSON tag name. Essentially the reflection packages are used to get all the values that could
// have been present in the file, and conjoin the object keys into a environment variable style key.
// For instance, the JSON lookup line of `database.host` would result in the Env key `MAPLE_DATABASE_HOST` being used.
// The type exchange is based on the type within the config object.
func LoadEnv(opts *Options) error {
	//Ensure it isn't nil
	if opts == nil {
		opts = new(Options)
	}

	typ := reflect.TypeOf(opts)
	val := reflect.ValueOf(opts)
	return loadEnvReflect(typ, val, "MAPLE")
}

func loadEnvReflect(typ reflect.Type, val reflect.Value, ns string) error {
	//If it's a pointer, get what it points to instead
	//NOTE: There's no nil check here since config shouldn't use it
	if typ.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
		typ = val.Type()
	}

	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		fldVal := val.Field(i)

		//Just use the JSON tag form
		jsonTag := fld.Tag.Get("json")
		if jsonTag == "" {
			continue //Empty, or non-encoded tag
		}

		name := strings.ToUpper(strings.TrimSpace(jsonTag))
		newNameSpace := fmt.Sprintf("%s_%s", ns, name)

		switch fld.Type.Kind() {
		case reflect.Ptr, reflect.Invalid, reflect.UnsafePointer, reflect.Func, reflect.Chan:
			//NOTE: No current support for pointers!
			//The rest of the matches can't be encoded eitherway
			continue
		case reflect.Struct:
			//Break down the struct by recursively running this function
			if err := loadEnvReflect(fld.Type, fldVal, newNameSpace); err != nil {
				return err
			}
		default: //Attempt to decode the value by looking for the environment variable
			if !fldVal.CanSet() {
				continue //Cannot set the value anyways
			}

			//Check the environment variable
			// fmt.Printf("checking environment variable %s... ", newNameSpace)
			env := os.Getenv(newNameSpace)
			if env == "" {
				// fmt.Print("nothing\n")
				continue //No environment variable
			}
			fmt.Printf("'%s'\n", env)
			//We have a string for the environment variable lookup, attempt to convert it

			//Double switch since we know that it is assignable, but not which method we want yet
			switch fld.Type.Kind() {
			case reflect.String:
				fldVal.SetString(env)
				fmt.Printf("\t(set string value '%s' for %s)\n", env, name)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				enc, err := strconv.ParseInt(env, 10, 64)
				if err != nil {
					return fmt.Errorf("parsing integer type %s from environment variable %s, failed because of: %s", fld.Type.Kind().String(), newNameSpace, err.Error())
				}
				fldVal.SetInt(enc)
				fmt.Printf("\t(set integer value %d for %s)\n", enc, name)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				enc, err := strconv.ParseUint(env, 10, 64)
				if err != nil {
					return fmt.Errorf("parsing unsigned-integer type %s from environment variable %s, failed because of: %s", fld.Type.Kind().String(), newNameSpace, err.Error())
				}
				fldVal.SetUint(enc)
				fmt.Printf("\t(set unsigned-integer value %d for %s)\n", enc, name)
			case reflect.Float32, reflect.Float64:
				enc, err := strconv.ParseFloat(env, 64)
				if err != nil {
					return fmt.Errorf("parsing float type %s from environment variable %s, failed because of: %s", fld.Type.Kind().String(), newNameSpace, err.Error())
				}
				fldVal.SetFloat(enc)
			case reflect.Complex64, reflect.Complex128:
				continue //No support for now
			case reflect.Bool:
				//Util package provides a way to parse booleans for us.
				//Basically accepts the values in any variety of:
				//  true, TRUE, t, 1
				enc := utils.ParseBool(env)
				fldVal.SetBool(enc)
				fmt.Printf("\t(set boolean value %v for %s)\n", enc, name)
			case reflect.Array, reflect.Map, reflect.Slice:
				//Use the JSON format for these
				err := json.Unmarshal([]byte(env), fldVal.Interface())
				if err != nil {
					return fmt.Errorf("parsing %s from environment variable %s using JSON, failed because of: %s", fld.Type.Kind().String(), newNameSpace, err.Error())
				}
			}
		}
	}

	return nil
}
