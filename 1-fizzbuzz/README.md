# FizzBuzz

The program should print the numbers from 1 to 100.

Exceptions: 

- If the number is divisible by 3 print ‘fizz’,
- If the number is divisible by 5 print ‘buzz’ 
- If the number is divisible by both print ‘fizzbuzz’.

Output should look like this:

```
1
2
fizz
4
buzz
fizz
7
8
fizz
buzz
11
fizz
13
14
fizzbuzz
16
17
fizz
19
buzz
fizz
22
23
fizz
buzz
26
fizz
28
29
fizzbuzz
...
```

## Help

Achieving everything at once can be challenging. Always try to see if you can
divide a problem in smaller problems easier to tackle!

<details>
<summary>Tip 1</summary>

Try to display all the numbers from 1 to 100.
</details>

<details>
<summary>Solution 1</summary>

```go
// We initialized i at 1.
// If i is lower or equal to 100 execute what is in the brackets, increment i by 1 and read this line again!
for i := 1; i <= 100; i++ {
	fmt.Println(i)
}
```
</details>

<details>
<summary>Tip 2</summary>

Now, how can we display `fizz` if the number is divisible by 3?

Have you heard about the `modulo (%)` operator? It allows to find the rest of a division.

A number X is divisible by a number Y if the rest of the division is 0!

```go
func main() {
	fmt.Println(1 % 2) // 1 ~~ 1 / 2 = 0; rest 1
	fmt.Println(4 % 2) // 0 ~~ 4 / 2 = 2; rest 0
	fmt.Println(5 % 3) // 2 ~~ 5 / 3 = 1; rest 2
}
```

</details>

<details>
<summary>Solution 2</summary>

```go
for i := 1; i <= 100; i++ {
	if i%3 == 0 { // if i is divisible by 3
		fmt.Println("fizz") // print "fizz"
	} else { // otherwise
		fmt.Println(i) // print the number
	}
}
```
</details>

<details>
<summary>Tip 3</summary>

Now, how can we display `buzz` is the number is divisible by 5?

Have a look at a number like `15` (divisible by 3 and 5); does it looks correct?
If not, what do you think the issue is?

</details>

<details>
<summary>Solution 3</summary>

The following solution *DOES NOT WORK*:
```go
for i := 1; i <= 100; i++ {
	if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
}	
```

For a number divisible by 3 and 5 (like 15), the code will do:
- is the number divisible by 3?
- Yes! Display "fizz" and exit the condition.
- Oh no! It doesn't display "fizzbuzz"... 

This is why you need to handle this case separately.
As often in computer science, there are multiple way to achieve this! Have a look at the solutions folders for different way to do!

Here is one way to do it:

```go
for i := 1; i <= 100; i++ {
	if i%3 == 0 && i%5 == 0{ // note that this condition needs to be the first one! 
		fmt.Println("fizzbuzz")
	} else if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
}
```


</details>
