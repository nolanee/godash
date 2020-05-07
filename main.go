package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello!")
	arr := [8]int{}
	i := FindIndex(arr, Identity, 0)
	fmt.Println(i)

}

// Identity is the default predicate. It checks whether two values are the same
func Identity(x interface{}, y interface{}) bool {
	return x == y
}

// FindIndex returns the first index at which predicate is true. TODO add defaault values
func FindIndex(arr interface{}, predicate func(interface{}, interface{}) bool, fromIndex int) int {
	t := reflect.TypeOf(arr)
	fmt.Println(t)
	return 0
}
