package blindgame

import (
	"fmt"
	"sync"
)

type ResultOfBGRound struct {
	Player1Points int8
	Player2Points int8
	IsDraw        bool
}

type BGPlayer struct {
	UserID   int64
	Username string
	Score    int8
	LastMove int  // Последнее выбранное число (1-5)
	IsReady  bool // Готов ли к следующему раунду
}

type BGSession struct {
	mu      sync.Mutex
	Player1 *BGPlayer
	Player2 *BGPlayer
	GameID  string
}

func NewBGSession(player1ID int64, player1Name string, player2ID int64, player2Name string) *BGSession {
	return &BGSession{
		Player1: &BGPlayer{UserID: player1ID, Username: player1Name, Score: 0},
		Player2: &BGPlayer{UserID: player2ID, Username: player2Name, Score: 0},
		GameID:  fmt.Sprintf("%d_%d", player1ID, player2ID),
	}
}
