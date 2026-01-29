package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	tele "gopkg.in/telebot.v4"
)

func RegistrHandlers(b *tele.Bot) {
	b.Handle("/start", ShowMainMenu)

	// Main menu
	b.Handle(&kb.ChooseGameBtn, ShowLlistOfGames)
	b.Handle(&kb.RulesOfBGBtn, ShowRulesOfBG)
	b.Handle(&kb.StartBGBtn, AddUserInQueue)

	b.Handle(&kb.BGBtn1, HandleBGBtn1)
}
