package system

import (
	"fmt"
	"log"
	"sync"

	tele "gopkg.in/telebot.v4"
)

const (
	StateMainMenu  = "MainMenu"
	StateRulesMenu = "RulesMenu"
	StateQueue     = "Queue"
)

//init...

var UserMap *MyUsers

func InitUserMap() {
	UserMap = &MyUsers{
		Map: make(map[int64]User),
	}
}

type User struct {
	State    string
	NickName string
	ID       int64
	GameID   int64
	LastCtx  tele.Context
}

type MyUsers struct {
	Map map[int64]User
	mu  sync.Mutex
}

func (m *MyUsers) AddUser(c tele.Context) error {
	id := c.Sender().ID
	nick := c.Sender().FirstName
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Map[id]; exists {
		return fmt.Errorf("пользователь %d уже существует", id)
	}

	m.Map[id] = User{
		State:    StateMainMenu,
		ID:       id,
		GameID:   0,
		NickName: nick,
	}

	log.Printf("Пользователь %d добавлен", id)
	return nil
}
func (m *MyUsers) GetUserState(id int64) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[id]
	if !exists {
		return "", fmt.Errorf("пользователь %d не найден", id)
	}

	return user.State, nil
}

func (m *MyUsers) UpdateUserState(id int64, newState string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. Получаем пользователя из map (это КОПИЯ)
	user, exists := m.Map[id]
	if !exists {
		return fmt.Errorf("пользователь %d не найден", id)
	}

	// 2. Меняем поле у копии
	user.State = newState

	// 3. Кладем ОБНОВЛЕННУЮ копию обратно в map
	m.Map[id] = user

	return nil
}

func (m *MyUsers) GetUserLastCtx(id int64) (tele.Context, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[id]
	if !exists {
		return user.LastCtx, fmt.Errorf("пользователь %d не найден", id)
	}

	return user.LastCtx, nil
}

func (m *MyUsers) UpdateUserLastCtx(id int64, c tele.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[id]
	if !exists {
		return fmt.Errorf("пользователь %d не найден", id)
	}

	user.LastCtx = c

	m.Map[id] = user

	return nil

}

func (m *MyUsers) GetUserNickName(id int64) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[id]
	if !exists {
		return "", fmt.Errorf("пользователь %d не найден", id)
	}

	return user.NickName, nil
}

func (m *MyUsers) UpdateUserGameID(userId, gameId int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[userId]
	if !exists {
		return fmt.Errorf("пользователь %d не найден", userId)
	}

	user.GameID = gameId

	m.Map[userId] = user

	return nil

}

func (m *MyUsers) GetUserGameID(userId int64) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.Map[userId]
	if !exists {
		return 0, fmt.Errorf("пользователь %d не найден", userId)
	}

	return user.GameID, nil
}
