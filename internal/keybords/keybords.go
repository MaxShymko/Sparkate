package keybords

import (
	tele "gopkg.in/telebot.v4"
)

var (
	MainMenu           *tele.ReplyMarkup
	StartGameBtn       tele.Btn
	ShowRulesOfGameBtn tele.Btn
	// profile
	// rating

	BackToMainMenuMenu *tele.ReplyMarkup
	BackToMainMenuBtn  tele.Btn

	GameMenu           *tele.ReplyMarkup
	GameBtn1           tele.Btn
	GameBtn2           tele.Btn
	GameBtn3           tele.Btn
	GameBtn4           tele.Btn
	GameBtn5           tele.Btn
	BackToMainMenuBtn2 tele.Btn
)

func InitKeybords() {
	initMainMenu()
	initBackToMainMenuMenu()
	initGameMenu()
}

func initMainMenu() {
	MainMenu = &tele.ReplyMarkup{}

	StartGameBtn = MainMenu.Data("Start Blind Game", "StartGameBtn")
	ShowRulesOfGameBtn = MainMenu.Data("Rules of Blind Game", "RulesOfGameBtn")

	MainMenu.Inline(
		MainMenu.Row(StartGameBtn),
		MainMenu.Row(ShowRulesOfGameBtn),
	)
}

func initBackToMainMenuMenu() {
	BackToMainMenuMenu = &tele.ReplyMarkup{}

	BackToMainMenuBtn = BackToMainMenuMenu.Data("Leave to main menu", "BackToMainMenuBtn")

	BackToMainMenuMenu.Inline(
		BackToMainMenuMenu.Row(BackToMainMenuBtn),
	)
}

func initGameMenu() {
	GameMenu = &tele.ReplyMarkup{}

	GameBtn1 = GameMenu.Data("1️⃣", "GameBtn10")
	GameBtn2 = GameMenu.Data("2️⃣", "GameBtn20")
	GameBtn3 = GameMenu.Data("3️⃣", "GameBtn30")
	GameBtn4 = GameMenu.Data("4️⃣", "GameBtn40")
	GameBtn5 = GameMenu.Data("5️⃣", "GameBtn50")
	BackToMainMenuBtn2 = GameMenu.Data("Leave to main menu", "BackToMainMenuBtn")

	GameMenu.Inline(
		GameMenu.Row(GameBtn1, GameBtn2, GameBtn3, GameBtn4, GameBtn5),
		GameMenu.Row(BackToMainMenuBtn2),
	)
}
