package main

import (
	_ "github.com/lib/pq"
	"database/sql"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

var db *sql.DB

type Chattable interface {
    SendMessage()
}

func initDatabse() error {
	db, err := sql.Open("postgres", "postgres://gorod@localhost?port=5433")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func GetChats() ([]*tgbotapi.Chat, error) {
    rows, err := db.Query("SELECT * FROM chat")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	chats := make([]*tgbotapi.Chat, 0)
	for rows.Next() {
		chat := new(tgbotapi.Chat)
		err := rows.Scan(&chat.ID, &chat.UserName, &chat.Title)
		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}
