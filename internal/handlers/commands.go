package handlers

import (
	"fmt"
	"log"

	"github.com/ilya-shymko/Sparkate/internal/game"
	sys "github.com/ilya-shymko/Sparkate/internal/system"

	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	ph "github.com/ilya-shymko/Sparkate/internal/photos"
	tele "gopkg.in/telebot.v4"
)

func ShowMainMenu(c tele.Context) error {
	if err := sys.UserMap.AddUser(c); err != nil {
		log.Printf("произошла ошибка в функции (ShowMainMenu): %v", err)
	}

	photo := &tele.Photo{
		File:    ph.MainPhoto.File,
		Caption: "Main menu: ",
	}

	// Параметры отправки
	options := &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: kb.MainMenu,
	}

	// Отправляем
	return c.Send(photo, options)
}

func mainHandler(c tele.Context, caption string, menu *tele.ReplyMarkup) error {
	// Создаем фото с текущей подписью
	photo := &tele.Photo{
		File:    ph.MainPhoto.File,
		Caption: caption,
	}

	// Параметры отправки
	options := &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: menu,
	}

	// Отправляем
	return c.Edit(photo, options)
}

func StartGame(ctx2 tele.Context, id1, id2 int64) error {
	ctx1, _ := sys.UserMap.GetUserLastCtx(id1)

	counter := sys.GameSessionCounter.Add(1)

	sys.UserMap.UpdateUserGameID(id1, counter)
	sys.UserMap.UpdateUserGameID(id2, counter)

	sys.GameSessionMap.CreateGameSession(counter, id1, id2)
	sys.UserMap.UpdateUserState(id1, sys.StateChoiceDigit)
	sys.UserMap.UpdateUserState(id2, sys.StateChoiceDigit)

	str1, _ := sys.UserMap.GetUserNickName(id1)
	str2, _ := sys.UserMap.GetUserNickName(id2)

	capt1 := fmt.Sprintf("Your rival is - %s", str2)
	capt2 := fmt.Sprintf("Your rival is - %s", str1)

	mainHandler(ctx1, capt1, kb.GameMenu)
	mainHandler(ctx2, capt2, kb.GameMenu)
	return nil
}

func StartRound(ctx1, ctx2 tele.Context, id1, id2, sessionId int64, player1choice, player2choice int8) error {
	sys.UserMap.UpdateUserState(id1, sys.StateChoiceDigit)
	sys.UserMap.UpdateUserState(id2, sys.StateChoiceDigit)

	result, _ := game.CalculateResults(player1choice, player2choice)

	oldpts1, _ := sys.GameSessionMap.GetPlayerPoints(sessionId, id1)
	oldpts2, _ := sys.GameSessionMap.GetPlayerPoints(sessionId, id2)

	totalpts1 := oldpts1 + result.Player1Points
	totalpts2 := oldpts2 + result.Player2Points

	nick1, _ := sys.UserMap.GetUserNickName(id1)
	nick2, _ := sys.UserMap.GetUserNickName(id2)

	if totalpts1 >= 21 {

		sys.UserMap.UpdateUserState(id1, sys.StateEndOfGame)
		sys.UserMap.UpdateUserState(id2, sys.StateEndOfGame)

		sys.GameSessionMap.DeleteGameSession(sessionId)
		sys.UserMap.UpdateUserGameID(id1, 0)
		sys.UserMap.UpdateUserGameID(id2, 0)

		capt := fmt.Sprintf("winner - %s\n%s - %d points\n%s - %d points", nick1, nick1, totalpts1, nick2, totalpts2)

		mainHandler(ctx1, capt, kb.BackToMainMenuMenu)
		mainHandler(ctx2, capt, kb.BackToMainMenuMenu)
		return nil
	} else if totalpts2 >= 21 {
		sys.UserMap.UpdateUserState(id1, sys.StateEndOfGame)
		sys.UserMap.UpdateUserState(id2, sys.StateEndOfGame)

		sys.GameSessionMap.DeleteGameSession(sessionId)
		sys.UserMap.UpdateUserGameID(id1, 0)
		sys.UserMap.UpdateUserGameID(id2, 0)

		capt := fmt.Sprintf("winner - %s\n%s - %d points\n%s - %d points", nick2, nick1, totalpts1, nick2, totalpts2)

		mainHandler(ctx1, capt, kb.BackToMainMenuMenu)
		mainHandler(ctx2, capt, kb.BackToMainMenuMenu)
		return nil
	} else {

		sys.GameSessionMap.UdatePlayerPoints(sessionId, id1, totalpts1)
		sys.GameSessionMap.UdatePlayerPoints(sessionId, id2, totalpts2)

		capt1 := fmt.Sprintf("Your rival is - %s\nYour choice is - %d\nRival's choice is - %d\nYour points is - %d\nRival's points is - %d", nick2, player1choice, player2choice, totalpts1, totalpts2)
		capt2 := fmt.Sprintf("Your rival is - %s\nYour choice is - %d\nRival's choice is - %d\nYour points is - %d\nRival's points is - %d", nick1, player2choice, player1choice, totalpts2, totalpts1)

		mainHandler(ctx1, capt1, kb.GameMenu)
		mainHandler(ctx2, capt2, kb.GameMenu)
		return nil
	}
}
