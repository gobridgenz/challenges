package main

import (
	"fmt"
)

func main() {
	printTree(5)
}

func printTree(size int) {
	for i := 1; i <= size; i++ {
		for j := 0; j < size-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < size-2; j++ {
			fmt.Print(" ")
		}
		fmt.Println("III")
	}
}
