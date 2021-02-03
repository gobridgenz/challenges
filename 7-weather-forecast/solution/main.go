package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	loc, err := userLocation()
	if err != nil {
		log.Fatalf("error getting coordinates: %v", err)
	}

	weather, err := weatherForecast(loc)
	if err != nil {
		log.Fatalf("error getting weather: %v", err)
	}

	displayWeatherForecast(loc, weather)
}

type location struct {
	City        string
	CountryName string `json:"country_name"`
	Latitude    float64
	Longitude   float64
}

// userLocation fetches the GeoIP location from freegeoip.app
func userLocation() (*location, error) {
	resp, err := http.Get("https://freegeoip.app/json/")
	if err != nil {
		return nil, fmt.Errorf("error querying geoip API: %w", err)
	}
	defer resp.Body.Close()

	var coords location
	err = json.NewDecoder(resp.Body).Decode(&coords)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &coords, nil
}

type weatherData struct {
	// number of hours from now
	Timepoint           int    `json:"timepoint"`
	CloudCover          int    `json:"cloudcover"`
	PrecipitationType   string `json:"prec_type"`
	PrecipitationAmount int    `json:"prec_amount"`
	Temperature         int    `json:"temp2m"`
	Weather             string `json:"weather"`
	Wind                struct {
		Direction string
		Speed     int
	} `json:"wind10m"`
}

type weatherResponse struct {
	Dataseries []weatherData
}

// weatherForecast fetches the weather from the 7timer API, see
// http://www.7timer.info/doc.php?lang=en#civil
func weatherForecast(coords *location) ([]weatherData, error) {
	q := make(url.Values)
	q.Add("lat", fmt.Sprintf("%f", coords.Latitude))
	q.Add("lon", fmt.Sprintf("%f", coords.Longitude))
	q.Add("unit", "metric")
	q.Add("output", "json")

	u, err := url.Parse("http://www.7timer.info/bin/civil.php")
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error querying weather API: %w", err)
	}
	defer resp.Body.Close()

	var weather weatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return weather.Dataseries, nil
}

func displayWeatherForecast(loc *location, weather []weatherData) {
	fmt.Printf("Weather forecast for \033[1m%s, %s\033[0m\n\n", loc.City, loc.CountryName)

	today := weather[0]
	var tomorrow weatherData
	// find forecast for tomorrow
	for _, d := range weather {
		if d.Timepoint >= 24 {
			tomorrow = d
			break
		}
	}

	fmt.Printf("\033[4mToday\033[0m\n")
	displayWeatherDatapoint(today)

	fmt.Printf("\n\033[4mTomorrow\033[0m\n")
	displayWeatherDatapoint(tomorrow)
}

func displayWeatherDatapoint(d weatherData) {
	var weatherType string
	switch d.Weather {
	case "clearday", "clearnight":
		weatherType = "Clear"
	case "mcloudyday", "mcloudynight":
		weatherType = "Mostly Cloudy"
	case "pcloudyday", "pcloudynight":
		weatherType = "Partly Cloudy"
	case "cloudyday", "cloudynight":
		weatherType = "Cloudy"
	case "humidday", "humidnight":
		weatherType = "Foggy"
	case "lightrainday", "lightrainnight":
		weatherType = "Light rain"
	case "oshowerday", "oshowernight":
		weatherType = "Occasional showers"
	case "ishowerday", "ishowernight":
		weatherType = "Isolated showers"
	case "lightsnowday", "lightsnownight":
		weatherType = "Light snow"
	case "rainday", "rainnight":
		weatherType = "Rain"
	case "snowday", "snownight":
		weatherType = "Snow"
	case "rainsnowday", "rainsnownight":
		weatherType = "Rain and snow"
	default:
		weatherType = "Unknown"
	}

	fmt.Printf("\033[1m%d°C, %s\033[0m\n", d.Temperature, weatherType)
	fmt.Printf("Precipitation [\033[44m%s\033[0m%s]\n", spaces(d.PrecipitationAmount), spaces(9-d.PrecipitationAmount))
	// wind scale should be shorter but it's prettier that way
	fmt.Printf("Wind          [\033[47m%s\033[0m%s]\n", spaces(d.Wind.Speed-1), spaces(10-d.Wind.Speed))
}

func spaces(n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += " "
	}
	return res
}
