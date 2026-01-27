package handlers

import (
	kb "github.com/ilya-shymko/Sparkate/keybords"
	tele "gopkg.in/telebot.v4"
)

func ShowMainMenu(c tele.Context) error {
	// проверить в каком состоянии пользователь а потом уже выводить для него сообщение
	return c.Send("Main menu: ", kb.MainMenu)
}
