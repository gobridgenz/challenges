package main

import (
	"fmt"
)

func main() {
	printChristmasTree(5)
}

func printChristmasTree(height int) {
	if height < 1 {
		return
	} else if height == 1 {
		fmt.Println("o")
		fmt.Println("\"")
	}

	width := height*2 - 1

	for i := 1; i <= height; i++ {
		input := 0
		if i == 1 {
			input = 1

		} else {
			input = i*2 - 1

		}

		buildChristmasTree(width-input/2, input, "o")

	}

	buildChristmasTree(width, 1, "\"")

}

func buildChristmasTree(space int, char int, mychar string) {
	printSpace := ""
	printTree := ""

	for i := 0; i < space; i++ {
		printSpace += " "
	}
	for i := 0; i < char; i++ {
		printTree += mychar
	}

	fmt.Println(printSpace, printTree)
}
