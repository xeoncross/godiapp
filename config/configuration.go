package config

import (
	"encoding/json"
	"os"
)

// Configuration read from .json file
type Configuration struct {
	HTTP struct {
		Address string `json:"address"`
	} `json:"http"`
	MySQL struct {
		DSN string `json:"dsn"`
	} `json:"mysql"`
}

// Load from file and return Configuration object
func Load(filename string) (config Configuration, err error) {

	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return
}
