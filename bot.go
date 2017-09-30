package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var store = make(map[string]map[string]float64)

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Println("Account: ", bot.Self.UserName)

	return bot, nil
}

func InitWebHook(bot *tgbotapi.BotAPI, config Config) (tgbotapi.UpdatesChannel, error) {
	URL, err := url.Parse(config.Url)
	if err != nil {
		return nil, err
	}

	bot.SetWebhook(tgbotapi.WebhookConfig{URL: URL})

	updates := bot.ListenForWebhook(URL.Path)

	go http.ListenAndServe("localhost:"+config.Port, nil)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func ProcessUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Println(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text)
		text := update.Message.Text
		var response string

		account := strconv.FormatInt(int64(update.Message.From.ID), 10)
		money, err := strconv.ParseFloat(strings.Split(text, " ")[0], 64)
		if err != nil {
			log.Println("[ERROR]: ", err)
		}

		switch true {
		case strings.Contains(text, "/add"):
			response = addMoney(money, account, "default")
		case strings.Contains(text, "/rem"):
			response = "removing money"
		case strings.Contains(text, "/status"):
			response = "getting status"
		default:
			response = "i don't get what you want from me."
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func addMoney(money float64, account, category string) string {
	val, ok := store[account]
	if !ok {
		log.Println("store value: ", val)
		store[account] = make(map[string]float64)
		log.Println("store value: ", val)
	}
	log.Println("category", category)
	return "add: " + strconv.FormatFloat(money, 'f', -1, 64)
}
