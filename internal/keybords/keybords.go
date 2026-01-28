package keybords

import (
	tele "gopkg.in/telebot.v4"
)

var (
	MainMenu      *tele.ReplyMarkup
	ChooseGameBtn tele.Btn

	GameMenu          *tele.ReplyMarkup
	StartBlindGameBtn tele.Btn
)

func InitKeybords() {
	// 1
	MainMenu = &tele.ReplyMarkup{}

	ChooseGameBtn = MainMenu.Data("Choose game", "chooseGameBtn")

	MainMenu.Inline(
		MainMenu.Row(ChooseGameBtn),
	)

	// 2
	GameMenu = &tele.ReplyMarkup{}

	StartBlindGameBtn = GameMenu.Data("Start Blind Game", "startBlindGameBtn")

	GameMenu.Inline(GameMenu.Row(StartBlindGameBtn))
}
