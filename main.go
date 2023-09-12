package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Localtime string `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Current struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
	Condition        Condition
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load .env file: %v\n", err)
		return
	}
	apiToken := os.Getenv("WEATHER_API_TOKEN")

	city := "nova-lima"
	state := "mg"
	country := "br"
	reqURL := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s,%s&aqi=no", apiToken, url.QueryEscape(city), url.QueryEscape(state), url.QueryEscape(country))

	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Printf("Failed to request weather data: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP error: %v\n", resp.Status)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Failed to close body: %v\n", err)
			return
		}
	}(resp.Body)

	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Printf("Failed to decode weather data: %v\n", err)
		return
	}

	// print table with data using columns and rows
	fmt.Printf("Weather in %s, %s, %s\n", weather.Location.Name, weather.Location.Region, weather.Location.Country)
	fmt.Printf("Temperature: %.2f°C\n", weather.Current.TempC)
	fmt.Printf("Condition: %s\n", weather.Current.Condition.Text)
	fmt.Printf("Last updated: %s\n", weather.Current.LastUpdated)
	fmt.Printf("Local time: %s\n", weather.Location.Localtime)

}
