package godash

import (
	"errors"
	"reflect"
)

const badParameterMsg = "Godash error. Function takes slice as first parameter."

// Find returns the first value of slice for which predicte returns true within interface.
// Returns empty interface if not found.
func Find(arrayInterface interface{}, predicate func(interface{}) bool, fromIndex int) (found interface{}, err error) {
	i, err := FindIndex(arrayInterface, predicate, fromIndex)
	if err != nil {
		return
	}
	if i == -1 {
		return
	}
	found = reflect.ValueOf(arrayInterface).Index(i).Interface()
	return
}

// FindIndex returns the first index of a slice for which predicate returns true. If not found, returns -1.
func FindIndex(sliceInterface interface{}, predicate func(interface{}) bool, fromIndex int) (index int, err error) {
	value, err := readSlice(sliceInterface, badParameterMsg)
	if err != nil {
		index = -1
		return
	}
	l := value.Len()
	for i := fromIndex; i < l; i++ {
		if predicate(value.Index(i).Interface()) {
			index = i
			return
		}
	}
	index = -1
	return
}

func readSlice(slice interface{}, errorMessage string) (value reflect.Value, err error) {
	value = reflect.ValueOf(slice)
	if value.Type().Kind() != reflect.Slice {
		err = errors.New(errorMessage)
		return
	}
	return
}
