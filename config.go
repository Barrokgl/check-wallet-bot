package main

import (
    "os"
    "errors"
    "log"
)

type Config struct {
	Token string `json:"token"`
	Url string `json:"url"`
	Port string `json:"port"`
}

func GetConfig() Config {
    token := os.Getenv("TOKEN")
    checkEnvVar(token)

    url := os.Getenv("URL")
    checkEnvVar(url)

    port := os.Getenv("PORT")
    checkEnvVar(port)

	return Config{token, url, port}
}

func checkEnvVar(variable string) {
    if variable == "" {
        log.Fatal(errors.New("To start bot required 3 env vars: TOKEN, PORT, URL"))
        return
    }

    return 
}
