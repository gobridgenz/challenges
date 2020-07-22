package main

import (
	"fmt"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		s := ""
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}

		if s == "" {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}

	}
}
