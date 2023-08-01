package main

import (
	"cli_go/internal/cmd/get_weather_today/config"
	"fmt"
)

func main() {
	config, err := config.LoadConfiguration()

	if err != nil {
		return
	}

	fmt.Printf("Weather CLI Config %+v", config)
}
