# Weather forecast

## What you'll learn

- Use query parameters
- Combine results
- Console text formatting

## Goal

Let's find out what the weather forecast will be where you are!

This challenge involves two APIs and to combining their results.

_Start with the challenge 6 if you don't have experience with calling APIs!_

### IP Geolocation API

The first step is dertermining _where_ you are (what weather forecast you
need to get).

[Freegeoip](https://freegeoip.app/) gives you your location based on your IP
address.

Here is the URL to request:
[`https://freegeoip.app/json/`](https://freegeoip.app/json/).

<details>
<summary>API response</summary>

```json
{
  "ip": "xxx.xxx.xxx.xxx",
  "country_code": "NZ",
  "country_name": "New Zealand",
  "region_code": "AUK",
  "region_name": "Auckland",
  "city": "Auckland",
  "zip_code": "1010",
  "time_zone": "Pacific/Auckland",
  "latitude": -36.8506,
  "longitude": 174.7679,
  "metro_code": 0
}
```

</details>
<br>

Thanks to the latitude and longitude in the response, you can now request the
weather forecast API.

### Weather forecast API

The weather forecast URL is the following:

`http://www.7timer.info/bin/civil.php?lon=<FIXME>&lat=<FIXME>&unit=metric&output=json`

Of course, you will need to replace the two `<FIXME>` in the parameters by the latitude and
longitude returned by the ip geolocation API.

Click [here](http://www.7timer.info/bin/civil.php?lon=174.7679&lat=-36.8506&unit=metric&output=json) to see a response example in your browser.

<details>
<summary>Or click here to see an extract of the response </summary>

```json
{
  "product": "civil",
  "init": "2021012318",
  "dataseries": [
    {
      "timepoint": 3,
      "cloudcover": 9,
      "lifted_index": 6,
      "prec_type": "none",
      "prec_amount": 1,
      "temp2m": 20,
      "rh2m": "78%",
      "wind10m": {
        "direction": "W",
        "speed": 3
      },
      "weather": "cloudyday"
    },
    {
      "timepoint": 6,
      "cloudcover": 9,
      "lifted_index": 2,
      "prec_type": "rain",
      "prec_amount": 1,
      "temp2m": 24,
      "rh2m": "69%",
      "wind10m": {
        "direction": "SW",
        "speed": 3
      },
      "weather": "lightrainday"
    },
    {
      "timepoint": 9,
      "cloudcover": 9,
      "lifted_index": 6,
      "prec_type": "rain",
      "prec_amount": 2,
      "temp2m": 24,
      "rh2m": "77%",
      "wind10m": {
        "direction": "SW",
        "speed": 3
      },
      "weather": "lightrainday"
    }
  ]
}
```

</details>
<br>

#### How to interpret the result?

The documentation explaining the meaning of all the fields is available
[here](http://www.7timer.info/doc.php?lang=en#civil).

Here are a few examples of interesting fields:

- `timepoint`: defines in how many hours this forecast is.
- `cloudcover`: integer that vary from 1 to 9 (1 = "0%-6%"; 9 = "94%-100%")
- `temp2m`: integer that varies from -76 to 60 (temperature in C°)
- `prec_amount`: integer that vary from 0 to 9 (0 = "None"; 9 = "Over 75mm/hr")
- `prec_type`: string indicating the precipitation type (`snow`, `rain`, `frzr`
  (freezing rain), `icep` (ice pellets), `none`)

### Result Example

Have fun displaying the information you want, the way you want!
Here is an example of what you could achieve:

```txt
$> go run main.go
Weather forecast for Auckland, New Zealand
* Today *
20°C, Cloudy
Precipitation [XX       ]
Wind          [XXX      ]

* Tomorrow *
17°C, Partly Cloudy
Precipitation [X        ]
Wind          [XXXX     ]
```

Good luck and have fun!

## Help

Remember that you can search online at any time!

If you are stuck, check the steps to follow below. For each of them, you can
have a look at the corresponding tip and solution.

<details>
<summary>Steps to follow</summary>

1. Send a GET request to the IP Geolocation API and save the response in a
   struct that you have defined.

1. Send a GET request to the Weather forecast API using the lat and long from
   the first API call response to build the URL. Save the response in a struct
   that you have defined.

1. Display the result! Start to display the "raw" data from the API and try to
   make it prettier once you know that your values are correct!

</details>

<details>
<summary>Tip 1</summary>

You can enter the url in your favorite web browser and look how the response
looks like:
[`https://freegeoip.app/json/`](https://freegeoip.app/json/).

Focus on what you are trying to achieve: you only need the latitude and
longitude to pass to the next API request. Maybe you can also retrieve the name
of the country and city to display them?

Click [here](https://gobyexample.com/structs) to see an example of how to
declare a struct in Go.

And [here](https://gobyexample.com/json) to see how to give it the possibility
to retrieve some json fields.

To call the API and retrieve the response, have a look at the following
functions:

- [`http.Get`](https://golang.org/pkg/net/http/#example_Get): Send a GET
  request to the specified URL

- [`json.NewDecoder`](https://golang.org/pkg/encoding/json/#NewDecoder) and
  [`Decode`](https://golang.org/pkg/encoding/json/#Decoder.Decode): To decode the
  response body (formatted in JSON) into a given struct.

</details>

<details>
<summary>Solution 1</summary>

```go
// Location defines location information.
//
// It contains the response from freegeoip API.
type Location struct {
	Country string  `json:"country_name"`
	City    string  `json:"city"`
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
}

// getLocation returns the location based on the ip.
//
// It sends an HTTP GET request to freegeoip.
func getLocation() (*Location, error) {
	resp, err := http.Get("https://freegeoip.app/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	loc := Location{}
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return nil, err
	}

	return &loc, nil
}
```

</details>

<details>
<summary>Tip 2</summary>

To call the URL with the latitude and longitude, you can use the
[`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function.

It will help you to replace the two `FIXME` in
`http://www.7timer.info/bin/civil.php?lon=<FIXME>&lat=<FIXME>&unit=metric&output=json`.

You will need to define two structs to store the API response.

Why two? Have a closer look at the API response in your browser:
It's a json containing only 3 fields. Between them, only `dataseries` interests
us. It contains an array with the weather forecast for different point in time
(+3 hours, +6 hours +9 hours, +12 hours, etc...).

So you need to define a struct to get the dataseries and another one to get the
dataseries content.

_Note that [`nested structs`](https://play.golang.org/p/GHaY0uSbtkv) or
[`anonymous structs`](https://play.golang.org/p/eeTKEGAipqA) are also a thing
if you don't want to declare two separated structs. But declaring two structs
is a common practice!_

See `Tip 1` for a tip on how to make the GET request or define the structs.

</details>

<details>
<summary>Solution 2</summary>

```go
// WeatherForecast is a weather forecast for a location and time.
//
// It contains the response from the 7timer API.
type WeatherForecast struct {
	TimePoint     int `json:"timepoint"`
	Clouds        int `json:"cloudcover"`
	Precipitation int `json:"prec_amount"`
	Temperature   int `json:"temp2m"`
}

// getWeatherForecast get tomorrow's weather forecast for the given location.
//
// It sends an HTTP GET request to 7timer.
func getWeatherForecast(lon, lat float64) (*WeatherForecast, error) {
	// Send a get request to the 7timer API.
	path := fmt.Sprintf("http://www.7timer.info/bin/civil.php?lon=%f&lat=%f&unit=metric&output=json", lon, lat)
	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("failed to GET 7timer: %w", err)
	}
	defer resp.Body.Close()

	// Declare the struct that will receive the API response.
	//
	// Note that an anonymous struct is defined here as this code only display
	// tomorrow's weather forecast. So the complete dataseries isn't useful
	// outside of this function.
	//
	// You can also define another type (like "NextsWeatherForecast") outside of
	// this function if you want to display multiple series!
	data := struct {
		Series []WeatherForecast `json:"dataseries"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	// We loop through all the data series to find the "timepoint" for tomorrow
	// (in 24 hours).
	for _, wf := range data.Series {
		if wf.TimePoint == 24 {
			return &wf, nil
		}
	}

	// We return an error if we couldn't find this data.
	return nil, fmt.Errorf("no data for tomorrow")
}
```

</details>

<details>
<summary>Tip 3</summary>
Just have fun!
You can start by displaying some text.

If you want a nice graphical output, have a look at the scale in the
documentation.

</details>

<details>
<summary>Solution 3</summary>

This is the most basic solution possible. Have fun improving it ;)

```go
fmt.Printf("\nTomorrow:\nClouds: %d/9\nTemperature (°C): %d\nPrecipitation: %d/9\n",
	wf.Clouds,
	wf.Temperature,
	wf.Precipitation,
)
```

```txt
$> go run main.go
Tomorrow:
Clouds: 3/9
Temperature (°C): 17
Precipitation: 2/9
```

</details>

## Useful links

- [`http` package](https://golang.org/pkg/net/http/)
- [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code)
