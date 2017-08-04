package main

import "log"

func main() {
	log.Println("Starting bot")

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Get config token: ", config.Token, " and webhook url: ", config.Url)

	bot, err := InitBot(config.Token)
	if err != nil {
		log.Fatal(err)
	}

    updates, err := InitWebHook(bot, config)

	ProcessUpdates(updates, bot)
}
