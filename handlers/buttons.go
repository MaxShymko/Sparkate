package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/keybords"
	tele "gopkg.in/telebot.v4"
)

func ShowLlistOfGames(c tele.Context) error {
	return c.Edit("List of games available for you: ", kb.GameMenu)
}
