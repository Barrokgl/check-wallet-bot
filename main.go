package main

import "log"

func main() {
	log.Println("Starting bot")

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Get config token: ", config.Token, " and webhook url: ", config.Url)

	updates, err := InitBot(config)
	if err != nil {
		log.Fatal(err)
	}

	ProcessUpdates(updates)
}
