package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	apiKey  = "eb544f912c613bb4f0b7d9d5da2b189e" //(change your key here)//
	baseURL = "http://api.openweathermap.org/data/2.5/weather"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// Weather handler function to fetch weather data
func getWeather(w http.ResponseWriter, r *http.Request) {
	// Get city name from query parameter
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Please provide a city name", http.StatusBadRequest)
		return
	}

	// Build URL with the API key and city
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)

	// Create a Resty client
	client := resty.New()

	// Send GET request to OpenWeatherMap API
	resp, err := client.R().Get(url)
	if err != nil {
		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON response into the WeatherResponse struct
	var weather WeatherResponse
	err = json.Unmarshal(resp.Body(), &weather)
	if err != nil {
		http.Error(w, "Error unmarshaling response", http.StatusInternalServerError)
		return
	}

	// Return weather data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func main() {
	// Set up HTTP routes
	http.HandleFunc("/weather", getWeather)

	// Start HTTP server
	fmt.Println("Weather app is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
