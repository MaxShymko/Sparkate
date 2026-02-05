package system

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

const (
	StateEndOfGame    = "StateEndOfGame"
	StateChoiceDigit  = "StateChoiceDigit"
	StateWaitingRival = "StateWaitingRival"
)

var GameSessionCounter atomic.Int64

func InitGameSessionCounter() {
	GameSessionCounter.Store(1)
}

type GameSession struct {
	player1 player
	player2 player
	ID      int64
}

type MyGameSessions struct {
	Map map[int64]GameSession
	mu  sync.Mutex
}

var GameSessionMap *MyGameSessions

func InitGameSessionMap() {
	GameSessionMap = &MyGameSessions{
		Map: make(map[int64]GameSession),
	}
}

type player struct {
	ID         int64
	points     int8
	lastChoice int8
}

func (m *MyGameSessions) CreateGameSession(idOfGameSession, idOfplayer1, idOfplayer2 int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Map[idOfGameSession]; exists {
		log.Printf("Игровая сессия %d не добавлена", idOfGameSession)
		return fmt.Errorf("Игровая сессия %d уже существует", idOfGameSession)
	}

	m.Map[idOfGameSession] = GameSession{
		player1: player{
			ID:     idOfplayer1,
			points: 0,
		},
		player2: player{
			ID:     idOfplayer2,
			points: 0,
		},
		ID: idOfGameSession,
	}

	log.Printf("Игровая сессия %d добавлена", idOfGameSession)
	return nil
}

func (m *MyGameSessions) DeleteGameSession(id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Map[id]; !exists {
		log.Printf("Игровая сессия %d не существует", id)
		return fmt.Errorf("Игровая сессия %d не существует", id)
	}

	delete(m.Map, id)

	log.Printf("Игровая сессия %d удалена", id)
	return nil
}

func (m *MyGameSessions) IsWaitingSmb(sessionId int64) (bool, int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id1 := m.Map[sessionId].player1.ID
	state1, _ := UserMap.GetUserState(id1)
	if state1 == StateWaitingRival {
		return true, id1
	}

	id2 := m.Map[sessionId].player2.ID
	state2, _ := UserMap.GetUserState(id2)
	if state2 == StateWaitingRival {
		return true, id2
	}

	return false, 0
}

func (m *MyGameSessions) UdatePlayerLastChoice(sessionId, playerId int64, newChoice int8) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.Map[sessionId]
	if !exists {
		return fmt.Errorf("сессия %d не найдена", sessionId)
	}

	if playerId == session.player1.ID {

		session.player1.lastChoice = newChoice
		m.Map[sessionId] = session

	} else if playerId == session.player2.ID {

		session.player2.lastChoice = newChoice
		m.Map[sessionId] = session

	}
	return nil
}

func (m *MyGameSessions) GetPlayerLastChoice(sessionId, playerId int64) (int8, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.Map[sessionId]
	if !exists {
		return 0, fmt.Errorf("сессия %d не найдена", sessionId)
	}

	if m.isFirstPlayer(sessionId, playerId) {
		return session.player1.lastChoice, nil
	} else {
		return session.player2.lastChoice, nil
	}

}

func (m *MyGameSessions) GetPlayerPoints(sessionId, playerId int64) (int8, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.Map[sessionId]
	if !exists {
		return 0, fmt.Errorf("сессия %d не найдена", sessionId)
	}

	if m.isFirstPlayer(sessionId, playerId) {
		return session.player1.points, nil
	} else {
		return session.player2.points, nil
	}
}

func (m *MyGameSessions) UdatePlayerPoints(sessionId, playerId int64, newPoints int8) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.Map[sessionId]
	if !exists {
		return fmt.Errorf("сессия %d не найдена", sessionId)
	}

	if m.isFirstPlayer(sessionId, playerId) {

		session.player1.points = newPoints
		m.Map[sessionId] = session
	} else {
		session.player2.points = newPoints
		m.Map[sessionId] = session
	}
	return nil
}

func (m *MyGameSessions) isFirstPlayer(sessionId, playerId int64) bool {

	if m.Map[sessionId].player1.ID == playerId {
		return true
	} else {
		return false
	}

}
func (m *MyGameSessions) IsFirstPlayer(sessionId, playerId int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Map[sessionId].player1.ID == playerId {
		return true
	} else {
		return false
	}

}
