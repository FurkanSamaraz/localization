package types

import (
	"database/sql/driver"
	"reflect"
)

//NullValue is a dummy object for implementng the sql driver interfaces.
//It is intentially blank, this will always be null.
type NullValue struct{}

//Value implements the sql driver.Value interface, will always return nil
func (n NullValue) Value() (driver.Value, error) {
	return nil, nil
}

//Nullable returns a database usable NullValue object if the object
//provided is a pointer, and is nil. If the object is valid then it
//will just pass-through as usual.
//This is experimental abstraction for making anything nullable/optional
//in queries.
func Nullable(obj interface{}) interface{} {
	if reflect.TypeOf(obj).Kind() == reflect.Ptr && obj == nil {
		return NullValue{}
	}
	return obj
}