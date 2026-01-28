package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func initBot() *tele.Bot {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ BOT_TOKEN не установлен. Создайте файл .env или установите переменную окружения")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
