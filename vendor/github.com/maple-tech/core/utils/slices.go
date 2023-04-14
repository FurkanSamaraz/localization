package utils

import (
	"reflect"
)

//InterfaceSlice Turns a slice provided by interface{} and returns it transformed by reflection
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

//ObjectEqualsOtherFunc Function signature for lamdas returning true if objects equate
type ObjectEqualsOtherFunc func(a, b int) bool

//FindSliceDifferences Sorts between 2 slices (old, and new) and finds which entries are additions, subtractions, or the same and returns them as slices
//Uses the equate paramter lambda to judge if they are the same
func FindSliceDifferences(oldSlice interface{}, newSlice interface{}, equate ObjectEqualsOtherFunc) (addition []interface{}, subtraction []interface{}, same []interface{}) {
	old := InterfaceSlice(oldSlice)
	new := InterfaceSlice(newSlice)

	for indA, entA := range old {
		found := false
		for indB := range new {
			if equate(indA, indB) == true {
				found = true
				break
			}
		}

		if !found {
			subtraction = append(subtraction, entA)
		}
	}

	for indB, entB := range new {
		found := false
		for indA := range old {
			if equate(indA, indB) == true {
				found = true
				break
			}
		}

		if !found {
			addition = append(addition, entB)
		} else {
			same = append(same, entB)
		}
	}

	return
}

//SliceItemFunc Callback function that is supplied the index of the slice specified by the function
//ie. ForSliceDifference calls this with an index of the newSlice paramter
//If an error is returned it should be propogated through the parent function
type SliceItemFunc func(i int) error

//ForSliceDifferences Is a magical function that takes a slice of existing objects (oldSlice) and a slice of new objects (newSlice)
//and finds the equivelency of the objects in them (equate).
//
//If something in the newSlice doesn't exist in the oldSlice
//then the onAdd callback is called with the index of the item in the newSlice. If something in the oldSlice doesn't
//exist in the newSlice then the onRemove callback is called with the index of the item in the oldSlice.
//
//The equate function always orders the index by old first, then new. So func(a, b) can be looked at as func(oldIndex, newIndex)
//See! Magic. Not really...
//
//If an error occures in the SliceItemFunc callback, it'll terminate this function and return that error
func ForSliceDifferences(oldSlice interface{}, newSlice interface{}, equate ObjectEqualsOtherFunc, onAdd SliceItemFunc, onRemove SliceItemFunc) error {
	old := InterfaceSlice(oldSlice)
	new := InterfaceSlice(newSlice)

	for indA := range old {
		found := false
		for indB := range new {
			if equate(indA, indB) == true {
				found = true
				break
			}
		}

		if !found {
			//Subtracted
			err := onRemove(indA)
			if err != nil {
				return err
			}
		}
	}

	for indB := range new {
		found := false
		for indA := range old {
			if equate(indA, indB) == true {
				found = true
				break
			}
		}

		if !found {
			err := onAdd(indB)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//ObjectEqualsFunc Callback function for checking if Index of slice is what you are looking for (equates)
type ObjectEqualsFunc func(a int) bool

//FindInSlice Takes a slice and an ObjectEqualsFunc and loops through the slice to find an object you want
func FindInSlice(slice interface{}, equate ObjectEqualsFunc) (obj interface{}, ok bool) {
	interSlice := InterfaceSlice(slice)
	for ind, obj := range interSlice {
		if equate(ind) {
			return obj, true
		}
	}
	return nil, false
}
