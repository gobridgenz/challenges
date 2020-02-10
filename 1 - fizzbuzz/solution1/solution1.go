package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if fizz(i) && buzz(i) {
			fmt.Println("fizzbuzz")
		} else if fizz(i) {
			fmt.Println("fizz")
		} else if buzz(i) {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizz(value int) bool {
	return value%3 == 0
}

func buzz(value int) bool {
	return value%5 == 0
}
