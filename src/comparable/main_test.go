package main

import (
	"fmt"
	"testing"
)

func TestFirstCase(t *testing.T) {

	myMap := make(map[any]string)
	myMap[0] = "a"
	myMap[1] = "b"

	myFunc(myMap)
}

func myFunc[T comparable](x map[T]string) {
	for key, val := range x {
		fmt.Printf("key: %v(%T), val: %+v(%T)\n", key, key, val, val)
	}
}
