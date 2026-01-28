package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	tele "gopkg.in/telebot.v4"
)

func RegistrHandlers(b *tele.Bot) {
	b.Handle(tele.OnText, ShowMainMenu)

	b.Handle(&kb.ChooseGameBtn, ShowLlistOfGames)
}
