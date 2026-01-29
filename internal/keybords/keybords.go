package keybords

import (
	tele "gopkg.in/telebot.v4"
)

var (
	MainMenu      *tele.ReplyMarkup
	ChooseGameBtn tele.Btn
	RulesOfBGBtn  tele.Btn
	// profile
	// rating

	GameMenu   *tele.ReplyMarkup
	StartBGBtn tele.Btn

	BackMenu *tele.ReplyMarkup
	BackBtn  tele.Btn
	// BackToMainMeuBtn

	LeaveQueueMenu *tele.ReplyMarkup
	LeaveQueueBtn  tele.Btn

	BGMenu            *tele.ReplyMarkup
	BGBtn1            tele.Btn
	BGBtn2            tele.Btn
	BGBtn3            tele.Btn
	BGBtn4            tele.Btn
	BGBtn5            tele.Btn
	LeaveBGSessionBtn tele.Btn
)

func InitKeybords() {
	initMainMenu()
	initGameMenu()
	initBackMenu()
	initLeaveQueueMenu()
	initBGMenu()
}

func initMainMenu() {
	MainMenu = &tele.ReplyMarkup{}

	ChooseGameBtn = MainMenu.Data("Choose game", "chooseGameBtn")
	RulesOfBGBtn = MainMenu.Data("Rules of Blind Game", "RulesOfBGBtn")

	MainMenu.Inline(
		MainMenu.Row(ChooseGameBtn),
		GameMenu.Row(RulesOfBGBtn),
	)
}

func initGameMenu() {
	GameMenu = &tele.ReplyMarkup{}

	StartBGBtn = GameMenu.Data("Blind Game", "startBGBtn")
	GameMenu.Inline(
		GameMenu.Row(StartBGBtn),
	)
}

func initBackMenu() {
	BackMenu = &tele.ReplyMarkup{}

	BackBtn = BackMenu.Data("Back", "BackBtn")

	BackMenu.Inline(
		BackMenu.Row(BackBtn),
	)
}

func initLeaveQueueMenu() {
	LeaveQueueMenu = &tele.ReplyMarkup{}

	LeaveQueueBtn = LeaveQueueMenu.Data("Leave Queue", "LeaveQueueBtn")

	LeaveQueueMenu.Inline(
		LeaveQueueMenu.Row(LeaveQueueBtn),
	)
}

func initBGMenu() {
	BGMenu = &tele.ReplyMarkup{}

	BGBtn1 = BGMenu.Data("1️⃣", "BGBtn1")
	BGBtn2 = BGMenu.Data("2️⃣", "BGBtn2")
	BGBtn3 = BGMenu.Data("3️⃣", "BGBtn3")
	BGBtn4 = BGMenu.Data("4️⃣", "BGBtn4")
	BGBtn5 = BGMenu.Data("5️⃣", "BGBtn5")
	LeaveBGSessionBtn = BGMenu.Data("Leave", "LeaveBGSessionBtn")

	BGMenu.Inline(
		BGMenu.Row(BGBtn1, BGBtn2, BGBtn3, BGBtn4, BGBtn5),
		BGMenu.Row(LeaveBGSessionBtn),
	)
}
