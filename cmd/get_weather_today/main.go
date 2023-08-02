package main

import (
	"cli_go/internal/cmd/get_weather_today/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type GeoCodingResponse struct {
	Name    string
	Lat     float64
	Lon     float64
	Country string
	State   string
}

type WeatherResponse struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	Current   WeatherDetail `json:"current_weather"`
}

type WeatherDetail struct {
	Time        string
	Temperature float32
	WeatherCode int `json:"weathercode"`
}

func main() {
	config, err := config.LoadConfiguration()

	args := os.Args[1:]

	city := strings.ToUpper(args[0])
	geoCodeResp, err := getGeoCodeCity(city, config)

	if err != nil {
		return
	}

	resp, _ := getWeatherByCity(geoCodeResp.Lat, geoCodeResp.Lon)

	weatherMap := map[int]string{0: "Clear Sky", 1: "Mainly Clear", 2: "Partly Cloudy", 3: "Overcast", 61: "Slight Rain", 63: "Moderate Rain", 65: "Intense Rain", 80: "Slight Rain Showers", 81: "Moderate Rain Showers", 82: "Violent Rain Showers"}

	weather, ok := weatherMap[resp.Current.WeatherCode]

	if ok == true {
		fmt.Printf("Based on Open-Meteo, The weather today in `%v` at %v is %v with temperature %v degree celcius", city, resp.Current.Time, weather, resp.Current.Temperature)
	} else {
		log.Printf("Weather code for %v cannot be decoded", resp.Current.WeatherCode)
	}
}

func getGeoCodeCity(city string, config config.Configuration) (resp GeoCodingResponse, err error) {
	geocodingUrl := fmt.Sprintf("%v/geo/1.0/direct?limit=1&q=%v&appid=%v", config.ApiUrlBase, city, config.ApiKey)

	r, err := http.Get(geocodingUrl)

	if err != nil {
		log.Printf("Error to direct geo code location %v with status code %v", city, r.StatusCode)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	var response []GeoCodingResponse

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println("Error marshalling json response", err)
	}

	return response[0], err
}

func getWeatherByCity(lat float64, lon float64) (resp WeatherResponse, err error) {
	weatherUrl := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current_weather=true", lat, lon)

	r, err := http.Get(weatherUrl)

	if err != nil {
		log.Printf("Error to fetch weather data for location lat: %v | long: %v with status code %v", lat, lon, r.StatusCode)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &resp)

	if err != nil {
		log.Println("Error marshalling json response", err)
	}

	return resp, err
}
