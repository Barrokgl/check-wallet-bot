package main

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Token string `json:"token"`
	Url string `json:"url"`
	Port string `json:"port"`
}

func GetConfig() (Config, error){
	data, err := ioutil.ReadFile("./config/bot.json")
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

