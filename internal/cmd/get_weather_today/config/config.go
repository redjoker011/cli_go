package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	ApiKey     string `json:"openWeatherApiKey"`
	ApiUrlBase string `json:"openWeatherApiUrlBase"`
}

func LoadConfiguration() (config Configuration, err error) {
	fileName := "config.json"

	if _, err := os.Stat(fileName); err != nil {
		log.Printf("Error Opening the file %v with error %v", fileName, err)
	}

	file, _ := os.Open(fileName)
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Printf("Error Reading the file %v with error %v", fileName, err)
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		log.Printf("Error Marshalling JSON file %v with error %v", fileName, err)
	}

	defer file.Close()

	return config, err
}
