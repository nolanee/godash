package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello!")
	arr := [8]int{}
	pred := predicator(0)
	pred2 := predicator(1)
	i := FindIndex(arr, pred, 0)
	fmt.Println(i)
	i = FindIndex(arr, pred2, 0)
	fmt.Println(i)
	arr[0] = 1
	i = FindIndex(arr, pred2, 2)
	fmt.Println(i)
	i = FindIndex(arr, pred2, 0)
	fmt.Println(i)
	arr[7] = 1
	i = FindIndex(arr, pred2, 2)
	fmt.Println(i)

}

func predicator(search interface{}) func(interface{}) bool {
	return func(value interface{}) bool {
		return search == value
	}
}

// FindIndex returns the first index at which predicate is true. Will panic if interface is not an
// array, slice, or string
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
