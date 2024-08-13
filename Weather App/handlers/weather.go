package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/benM4k/go_projects/utils"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := queryWeather(city)

	if err != nil {
		log.Printf("Error fetching weather data: %v", err)
		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func queryWeather(city string) (WeatherData, error) {
	apiKey := os.Getenv("API_KEY")
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey)
	if err != nil {
		return WeatherData{}, err
	}

	defer response.Body.Close()

	var weatherData WeatherData

	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return WeatherData{}, err
	}

	weatherData.Main.Kelvin = float64(utils.Kelvin2Celsius(utils.Kelvin(weatherData.Main.Kelvin)))
	return weatherData, nil
}
