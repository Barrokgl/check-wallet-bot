package main

import "log"

func main() {
	config := GetConfig()

	bot, err := InitBot(config.Token)
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Bot started on port: ", config.Port)
    updates, err := InitWebHook(bot, config)

	ProcessUpdates(updates, bot)
}
