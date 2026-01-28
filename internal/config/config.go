package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️  .env файл не найден: %v", err)
		log.Println("ℹ️  Использую переменные окружения системы")
	}
}
