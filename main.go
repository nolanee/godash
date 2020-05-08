package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello!")
	arr := [8]int{}

	pred := predicator(0)
	i := FindIndex(arr, pred, 0)
	fmt.Println(i)

	pred2 := predicator(1)
	i = FindIndex(arr, pred2, 0)
	fmt.Println(i)
	arr[0] = 1
	i = FindIndex(arr, pred2, 2)
	fmt.Println(i)
	arr[7] = 1
	i = FindIndex(arr, pred2, 0)
	fmt.Println(i)
	i = FindIndex(arr, pred2, 2)
	fmt.Println(i)

	pred3 := predicator(7)
	arr[6] = 7
	v := Find(arr, pred3, 2).(int)
	fmt.Println(v)
	e := Find(arr, pred3, 7)
	fmt.Println(e)

}

func predicator(search interface{}) func(interface{}) bool {
	return func(value interface{}) bool {
		return search == value
	}
}

// FindIndex returns the first index at which predicate is true. Will panic if interface is not an
// array, slice, or string. If not found, returns -1.
func FindIndex(arrayInterface interface{}, predicate func(interface{}) bool, fromIndex int) int {
	arr := reflect.ValueOf(arrayInterface)
	l := arr.Len()
	for i := fromIndex; i < l; i++ {
		if predicate(arr.Index(i).Interface()) {
			return i
		}
	}
	return -1
}

// Find returns the first value at which is true within interface. Will panic if interface is not an
// array, slice, or string. If not found returns an empty interface.
func Find(arrayInterface interface{}, predicate func(interface{}) bool, fromIndex int) interface{} {
	i := FindIndex(arrayInterface, predicate, fromIndex)
	if i == -1 {
		var empty interface{}
		return empty
	}
	arrV := reflect.ValueOf(arrayInterface).Index(i).Interface()
	return arrV
}
