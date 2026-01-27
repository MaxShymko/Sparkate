package main

import (
	"log"
	"os"
	"time"

	h "github.com/ilya-shymko/Sparkate/handlers"
	kb "github.com/ilya-shymko/Sparkate/keybords"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️  .env файл не найден: %v", err)
		log.Println("ℹ️  Использую переменные окружения системы")
	}

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
		return
	}

	kb.Init()

	h.RegistrHandlers(b)

	log.Println("Bot has started...")
	b.Start()
}
