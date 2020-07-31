package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(1000) + 1

	for {
		var playerNumber int
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &playerNumber)

		if playerNumber > number {
			fmt.Println("It's less!")
		} else if playerNumber < number {
			fmt.Println("It's more!")
		} else {
			fmt.Println("✨ You win! ✨")
			return
		}
	}
}
