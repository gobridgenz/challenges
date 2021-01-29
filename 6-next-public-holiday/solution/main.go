package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// PublicHoliday contains the name and date of a public holiday.
type PublicHoliday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func main() {
	// Get all public holiday for the current year.
	publicHolidays, err := getNZPublicHolidays()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Find the next public holiday.
	ph, err := getNextPublicHoliday(publicHolidays)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Next public holiday is '%s' the %s.\n", ph.Name, ph.Date)
}

// getNZPublicHolidays send a request to retrieve all the public holidays of
// the current year.
func getNZPublicHolidays() ([]PublicHoliday, error) {
	const publicHolidayAPI = "https://date.nager.at/api/v2/publicholidays"
	year := time.Now().Year()
	path := fmt.Sprintf("%s/%d/NZ", publicHolidayAPI, year)

	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("could not fetch public holidays from API %s: %w", path, err)
	}

	publicHolidays := []PublicHoliday{}
	if err := json.NewDecoder(resp.Body).Decode(&publicHolidays); err != nil {
		return nil, fmt.Errorf("could not decode request body: %w", err)
	}

	return publicHolidays, nil
}

// getNextPublicHoliday returns the next public holiday.
func getNextPublicHoliday(publicHolidays []PublicHoliday) (*PublicHoliday, error) {
	now := time.Now()

	for _, ph := range publicHolidays {
		date, err := time.Parse("2006-01-02", ph.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time %s: %w", ph.Date, err)
		}

		if date.After(now) {
			return &ph, nil
		}
	}

	return nil, fmt.Errorf("the next public holiday will be in %d", now.Year()+1)
}
