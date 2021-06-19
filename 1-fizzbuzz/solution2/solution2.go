package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		s := ""
		if i%3 == 0 {
			s = "fizz"
		}
		if i%5 == 0 {
			s += "buzz"
		}

		if s == "" {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}
	}
}
