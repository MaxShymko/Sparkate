package handlers

import (
	bg "github.com/ilya-shymko/Sparkate/internal/games/blindgame"
	kb "github.com/ilya-shymko/Sparkate/internal/keybords"
	mm "github.com/ilya-shymko/Sparkate/internal/matchmaking"
	txt "github.com/ilya-shymko/Sparkate/internal/texts"
	tele "gopkg.in/telebot.v4"
)

// Main Menu

func ShowLlistOfGames(c tele.Context) error {
	return mainHandler(c, "List of games available for you: ", kb.GameMenu)
}

func ShowRulesOfBG(c tele.Context) error {
	return mainHandler(c, txt.RulesOfBG, kb.BackMenu)
}

// Game menu

func AddUserInQueue(c tele.Context) error {
	mm.Queue.AddUserInQueue(c)
	if mm.Queue.IsReady() {
		c1 := mm.Queue.Players[0]
		c2 := mm.Queue.Players[1]

		BGsession := bg.NewBGSession(c1.Sender().ID, c1.Sender().FirstName, c2.Sender().ID, c2.Sender().FirstName)

		s := mm.NewMMBGSession(c1, c2, BGsession)

		mm.MapOfMMBGSession.Map[BGsession.GameID] = s

		if err := mainHandler(c1, "игра найдена", kb.BGMenu); err != nil {
			return err
		}
		if err := mainHandler(c2, "игра найдена", kb.BGMenu); err != nil {
			return err
		}

		return nil
	} else {
		return mainHandler(c, "вы в очереди...", kb.LeaveQueueMenu)
	}
}

// ...
//BlindGame menu

func HandleBGBtn1(c tele.Context) {
}
