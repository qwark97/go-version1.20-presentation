package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	myList := [2]int{1, 2}
	mySlice := []int{1, 2}

	fmt.Printf("list: %v(%T)\n\n", myList, myList)
	fmt.Printf("slice: %v(%T)\n\n", mySlice, mySlice)

	// list to slice
	myNewSlice := myList[:]
	fmt.Printf("new slice: %v(%T)\n\n", myNewSlice, myNewSlice)

	// slice to list (hack)
	myNewListHack := *(*[2]int)(mySlice)
	fmt.Printf("new list (hack): %v(%T)\n\n", myNewListHack, myNewListHack)

	// slice to list (NEW)
	myNewList := [2]int(mySlice)
	fmt.Printf("new list: %v(%T)\n\n", myNewList, myNewList)

}
