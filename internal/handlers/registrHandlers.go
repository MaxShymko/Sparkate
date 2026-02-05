package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	tele "gopkg.in/telebot.v4"
)

func RegistrHandlers(b *tele.Bot) {
	b.Handle("/start", ShowMainMenu)

	b.Handle(&kb.ShowRulesOfGameBtn, ShowRulesOfGame)

	b.Handle("\fBackToMainMenuBtn", BackToMainMenu)

	b.Handle(&kb.StartGameBtn, SelectionOfRivals)

	b.Handle(&kb.GameBtn1, Helper1)
	b.Handle(&kb.GameBtn2, Helper2)
	b.Handle(&kb.GameBtn3, Helper3)
	b.Handle(&kb.GameBtn4, Helper4)
	b.Handle(&kb.GameBtn5, Helper5)
}
