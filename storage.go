package main

import (
	_ "github.com/lib/pq"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Chattable interface {
    SendMessage()    
}

func GetChats() (*tgbotapi.Chat) {
    return nil
}
