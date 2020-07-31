package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var pics = []string{
	`
               __
              \  ,\
              " =__)
`,
	`
             ,\
             \\\,_
              \  ,\
              " =__)
`,
	`
             ,\
             \\\,_
              \  ,\
         __,.-" =__)
       ."        )
      /         /
      |______--'
`,
	`
             ,\
             \\\,_
              \  ,\
         __,.-" =__)
       ."        )
      /        \/
      |______-\ \_
               '--'
`,
	`
             ,\
             \\\,_
              \  ,\
         __,.-" =__)
       ."        )
      /        \/\_
      |_____-\ \_-'
              '--'
`,
	`
             ,\
             \\\,_
              \  ,\
         __,.-" =__)
       ."        )
      /   ,    \/\_
      |    )_-\ \_-'
      '-----' '--'
`,
	`
             ,\
             \\\,_
              \  ,\
         __,.-" =__)
       ."        )
    ,_/   ,    \/\_
    \_|    )_-\ \_-'
      '-----' '--'
`}

const numberOfWords = 858

func main() {
	rand.Seed(time.Now().UnixNano())

	word, err := chooseWord()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// currentWord contains the current letters that have been guessed.
	hiddenWord := make([]string, 0, len(word))
	for range word {
		hiddenWord = append(hiddenWord, "_")
	}

	bad := -1

	for bad < len(pics)-1 && strings.Join(hiddenWord, "") != word {
		// Display hidden word
		fmt.Println(strings.Join(hiddenWord, ""))

		// Ask for letter
		var playerLetter string
		fmt.Print("Enter a letter: ")
		fmt.Scanf("%s", &playerLetter)

		// Check if the letter exists & replace the occurrences.
		exists := replaceLetter(playerLetter, word, hiddenWord)

		if !exists {
			bad++
			fmt.Println(pics[bad])
		}
	}

	if bad == len(pics)-1 {
		fmt.Printf("🔥 You loose! 🔥\nThe word was: %s\n", word)
	} else {
		fmt.Println(word)
		fmt.Println("✨ You win! ✨")

	}
}

// chooseWord selects a random word in words.txt.
func chooseWord() (string, error) {
	wordNum := rand.Intn(numberOfWords)

	// open the file: could error if the file is not found for example.
	file, err := os.Open("words.txt")
	if err != nil {
		return "", err
	}

	// make sure to close the file when you are done reading it.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read line by line
	for i := 0; i <= wordNum; i++ {
		// read a line and make it available in the `Text()` method.
		// any error is available in the `Err()` method.
		scanner.Scan()
	}

	return scanner.Text(), scanner.Err()
}

// replaceLetter replaces the letter in the current word and return true if the letter is part of the word.
func replaceLetter(letter, word string, hiddenWord []string) bool {
	exists := false

	for i, l := range word {
		if string(l) == letter {
			hiddenWord[i] = string(l)
			exists = true
		}
	}

	return exists
}
