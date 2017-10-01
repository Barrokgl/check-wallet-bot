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

		text := update.Message.Text
		parsedText := strings.Fields(text)
		var response string
		var money float64
		var err error
		var category string

		account := strconv.FormatInt(int64(update.Message.From.ID), 10)
		if len(parsedText) > 3 {
			money, err = strconv.ParseFloat(parsedText[1], 64)
			if err != nil {
				log.Println("[ERROR]: ", err)
			}

			category = parsedText[2]
		} else {
			money = 0
			category = "default"
		}


		switch true {
		case strings.Contains(text, "/start"):
			response = startMessage()
		case strings.HasPrefix(text, "+"):
			response = addMoney(money, account, category)
		case strings.HasPrefix(text, "-"):
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

func startMessage() string {
	return `This is check wallet bot, usage:
		   + <money> <category(optional)> - add money to wallet
		   - <money> <category(optional)> - remove money from wallet`
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
