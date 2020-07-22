# The bunny (Hangman game)

## Goal of the game

The goal of the game is to find a word by giving the correct letters.
After 7 false attempts, you loose the game.


## Where should I start?

With this challenge, we are going to learn how to read in a file!
So the first thing to do, is to download the `words.txt` file and to place it in your new Go project.

## Help

<details>
<summary>Tip 1</summary>

How to choose a random word from `words.txt`?

- select a random number between 0 and the maximum number of lines (858 lines);
- open the file;
- read the file line by line until you reached that line.

Try to search this information online!
</details>

<details>
<summary>Solution 1</summary>

```go
// chooseWord select a random word in words.txt.
func chooseWord() (string, error) {
	rand.Seed(time.Now().UnixNano())
    wordNum := rand.Intn(858)

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
```
</details>

<details>
<summary>Tip 2</summary>

You will need another variable `hiddenWord` to store the progress of the player.
It should be the same length than your chosen word, but it only contains: `_`.
</details>


<details>
<summary>Tip 3</summary>

Store the user input in a `letter` variable.
</details>

<details>
    <summary>Solution 3</summary>

```go
	var playerLetter string
	fmt.Print("Enter a letter: ")
	fmt.Scanf("%s", &playerLetter)
```
</details>

<details>
<summary>Tip 4</summary>

Write a function `replaceLetter` which replaces all occurrences a given letter in the hidden word.
</details>

<details>
<summary>Tip 5</summary>

If the player entered the wrong letter (the letter does not exists in `word`), display the hangman in the next "stage".
</details>

<details>
<summary>Solution 5</summary>

You can store the score of the player in a variable (for example: `bad`).
Each time the player give a wrong answer, increment this value and display the associated picture:

```go
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

bad := 1
fmt.Println(hangmanPics[bad])

```
</details>