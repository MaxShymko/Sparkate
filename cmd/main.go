package main

import (
	"log"

	cfg "github.com/ilya-shymko/Sparkate/internal/config"
	hand "github.com/ilya-shymko/Sparkate/internal/handlers"
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	mm "github.com/ilya-shymko/Sparkate/internal/matchmaking"
	ph "github.com/ilya-shymko/Sparkate/internal/photos"
)

func main() {
	log.Println("Init Config...")
	cfg.InitConfig()

	log.Println("Init Bot...")
	b := initBot()

	log.Println("Init Keybords...")
	kb.InitKeybords()

	log.Println("Registr Handlers...")
	hand.RegistrHandlers(b)

	log.Println("Init Photos...")
	ph.InitPhotos()

	log.Println("Init MM...")
	mm.InitMM()

	log.Println("Bot has started...")
	b.Start()
}
