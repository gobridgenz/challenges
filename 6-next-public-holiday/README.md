# Next public holiday

## Goal

The goal of this challenge is to display the next public holiday in New Zealand.
To do so, you can request this public API:

https://date.nager.at/api/v2/publicholidays/2021/NZ


If you want your program to work next year, find a way to replace `2021` from
the URL with the current year!

<details>
<summary>Click here to display an extract of the response.</summary>

```json
[
	{
		"date": "2021-01-01",
		"localName": "New Year's Day",
		"name": "New Year's Day",
		"countryCode": "NZ",
		"fixed": false,
		"global": true,
		"counties": null,
		"launchYear": null,
		"type": "Public"
	},
	{
		"date": "2021-01-04",
		"localName": "Day after New Year's Day",
		"name": "Day after New Year's Day",
		"countryCode": "NZ",
		"fixed": false,
		"global": true,
		"counties": null,
		"launchYear": null,
		"type": "Public"
	},
	{
		"date": "2021-02-08",
		"localName": "Waitangi Day",
		"name": "Waitangi Day",
		"countryCode": "NZ",
		"fixed": false,
		"global": true,
		"counties": null,
		"launchYear": null,
		"type": "Public"
	}
]
```
</details>

## Help

Remember that it's a good idea to search online at any time!

If you are stuck, check the steps to follow below. For each of them, you can have a
look at the corresponding tip and solution.

<details>
<summary>Steps to follow</summary>

1. Call the URL with the current year

1. Define the type that will contain the API response

1. Call the API and save the response in a variable

1. Find out which public holiday is next based on the current time


</details>

<details>
<summary>Tip 1</summary>

To call the URL with the current year, you can use the following functions:

- [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf): to insert a variable in a string

- [`time.Now`](https://golang.org/pkg/time/#Now): to get the current time

</details>

<details>
<summary>Solution 1</summary>

```go
// Define the base URL
const publicHolidayAPI = "https://date.nager.at/api/v2/publicholidays"

// Retrieve the current date and time
now := time.Now()

// Extract the year of the current time
year := now.Year()

// Construct the url with the base URL + current year + country code
path := fmt.Sprintf("%s/%d/NZ", publicHolidayAPI, year)

// Note that you could also make the country code variable :)
```

</details>

<details>
<summary>Tip 2</summary>

You can enter the url in your favorite web browser and look how the response
looks like.

Focus on what you are trying to achieve: display the name of the next public
holiday.

The only information you need are:
- `name`: display the name of the public holiday,
- `date`: date of the public holiday.

Click [here](https://gobyexample.com/structs) to see an example of how to
declare a struct in Go.

And [here](https://gobyexample.com/json) to see how to give it the possibility
to retrieve some json fields.

</details>

<details>
<summary>Solution 2</summary>

```go
// PublicHoliday contains the name and date of a public holiday.
type PublicHoliday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
```

</details>

<details>
<summary>Tip 3</summary>

To call the API and retrieve the response, have a look at the following
functions:

- [`http.Get`](https://golang.org/pkg/net/http/#example_Get): Send a GET
request to the specified URL

- [`json.NewDecoder`](https://golang.org/pkg/encoding/json/#NewDecoder) and
[`Decode`](https://golang.org/pkg/encoding/json/#Decoder.Decode): To decode the
response body (formatted in JSON) into a given struct.

</details>


<details>
<summary>Solution 3</summary>

```go
	// Send a GET request to the public holiday API (see tip 1 to make the year
	// always valid!)
	const path = "https://date.nager.at/api/v2/publicholidays/2021/NZ"
	resp, err := http.Get(path)

	// Remember to always handle the error! You can print it or return it if you
	// are in a function.
	if err != nil {
		fmt.Printf("Could not fetch public holidays from API %s: %w\n", path, err)
	}

	// Because the API's response is an array of public holidays, we need to store
	// it in a slice.
	// Let's start by declaring a slice of our struct (see declaration in
	// solution 2).
	publicHolidays := []PublicHoliday{}

	// We decode the response body in our struct and handle the error directly.
	if err := json.NewDecoder(resp.Body).Decode(&publicHolidays); err != nil {
		fmt.Printf("could not decode request body: %w\n", err)
	}
```

</details>


<details>
<summary>Tip 4</summary>	

The last step is to compare the public holidays date with our current time.


You will need to loop through each public holiday (from the api response) to:

1. [`Parse`](https://golang.org/pkg/time/#Parse) the date (see example
[here](https://gobyexample.com/time-formatting-parsing)): It will transform the
`string` into a `time`. The `time` type allows you to do operations with the
time (like comparing which time is after another one for example!)

1. Compare the current time with the date of the public holiday: if you have a
look at the API response, you can see that the dates are ordered (from Jan. to
Dec.). The first one to be greater than the current time is the correct one!
Have a look at the [`After`](https://golang.org/pkg/time/#Time.After) function.

</details>


<details>
<summary>Solution 4</summary>

```go
// getNextPublicHoliday takes in parameter all the public holidays or the current year
// and return the next incoming public holiday.
func getNextPublicHoliday(publicHolidays []PublicHoliday) (*PublicHoliday, error) {
	// Get the current date and time.
	now := time.Now()

	// Loop through all the public holidays this year.
	for _, ph := range publicHolidays {
		// Convert the public holiday date from a `string` into a `time`.
		// Thanks to this, we can use functions specific to time (otherwise, Go
		// doesn't know what our `string` contains).
		date, err := time.Parse("2006-01-02", ph.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time %s: %w", ph.Date, err)
		}

		// We can compare the public holiday's date with the current date.
		// The first public holiday to be after "now" is the good one (the api
		// response is ordered).
		if date.After(now) {
			return &ph, nil
		}
	}

	// Do not forget to handle the end of the year where the next public holiday
	// is the following year...
	return nil, fmt.Errorf("the next public holiday will be in %d!", now.Year() + 1)
}
```

</details>
