package util

import (
	"encoding/json"
	"os"
)

type configuration struct {
	Token string `json:"token"`
}

// Config is the config.json file
var Config configuration

func init() {
	f, err := os.Open("config.json")
	if err != nil {
		panic("error opening config.json: " + err.Error())
	}
	err = json.NewDecoder(f).Decode(&Config)
	if err != nil {
		panic("error decoding config.json: " + err.Error())
	}
}