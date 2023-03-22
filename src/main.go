package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			noCall()
		} else {
			call()
		}
	}
}

func call() {
	fmt.Println("hello")
}

func noCall() {
	fmt.Println("hello2")
}
