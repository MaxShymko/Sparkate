package main

import (
	"log"

	cfg "github.com/ilya-shymko/Sparkate/internal/config"
	hand "github.com/ilya-shymko/Sparkate/internal/handlers"
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	ph "github.com/ilya-shymko/Sparkate/internal/photos"
	sys "github.com/ilya-shymko/Sparkate/internal/system"
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

	log.Println("Init User Map...")
	sys.InitUserMap()

	log.Println("Init Game Session Map...")
	sys.InitGameSessionMap()

	log.Println("Init Game Session Counter...")
	sys.InitGameSessionCounter()

	log.Println("Bot has started...")
	b.Start()
}
