# Find the number

## Goal of the game

The game is going to choose a number between 0 and 1000.
The goal of the game if to find that number!

Each turn, you can make a guess and the game will tell you if the number is smaller, greater or if you won!

### Output example

```
Enter a number: 500
It's less!
Enter a number: 200
It's more!
Enter a number: 400
It's more!
Enter a number: 450
It's more!
Enter a number: 470
It's more!
Enter a number: 480
It's less!
Enter a number: 479
✨ You win! ✨
```

## Environment set up

This time, you will need to have your environment set up:
- A tool to write your code (we recommend [VSCode](https://code.visualstudio.com/download))
- [Go](https://golang.org/dl/) installed on you computer;

Don't hesitate to ask for help during a workshop if you have some trouble to install everything!

## Where should I start?

Try to "cut" the game in smaller chunks in your head. If you have no idea where to start, don't hesitate to look at the first tip! They are here to help you to build that game little by little.

If you don't know how to do something, try to Google it to find help online. If you are still stuck, look at the solution and try to understand it.

## Help

<details>
    <summary>Tip 1</summary>

    How to generate a random number between 0 and 1000?

    Try to search this information online: Go By Example is a good source!
</details>

<details>
    <summary>Solution 1</summary>

```go
rand.Seed(time.Now().UnixNano())
number := rand.Intn(1000) + 1
fmt.Println(number)
```
</details>

<details>
<summary>Tip 2</summary>

    How to read the number that the player suggested?

    Try to search for: `golang read input command line` in Google.
</details>

<details>
    <summary>Solution 2</summary>

```go
var playerNumber int
fmt.Scanf("%d", &playerNumber)
fmt.Println(playerNumber)
```
</details>

<details>
    <summary>Tip 3</summary>

    Create an infinite loop in which you need to:
    - Read the player number;
    - Check if the number is equal to the "generated number"
        - if yes, print something to show the player won and quit the game.
</details>

<details>
    <summary>Solution 3</summary>

```go
for {
	var playerNumber int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &playerNumber)

	if playerNumber == number {
		fmt.Println("You win!")
		return
	}
}
```
</details>

<details>
    <summary>Tip 4</summary>

    In the game loop, if the player number is smaller or greater than the number to find: display some help!
</details>

<details>
    <summary>Solution 4</summary>

```go
if playerNumber > number {
	fmt.Println("It's less!")
} else if playerNumber < number {
	fmt.Println("It's more!")
} else {
	fmt.Println("You win!")
	return
}
```
</details>

## Bonus

Make the game yours!
- Add a fancy title when the game start;
- Limit the number of tries to 10;
- Add a 2 players option;
- Add a little animation when the player win;
- Anything you want!