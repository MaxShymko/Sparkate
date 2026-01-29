package matchmaking

import (
	bg "github.com/ilya-shymko/Sparkate/internal/games/blindgame"
	tele "gopkg.in/telebot.v4"
)

var (
	Queue            *BGQueue
	MapOfMMBGSession *MMBGSessionMap
	MapOfUser        *UsersMap
)

type BGQueue struct {
	Players []tele.Context
}
type MMBGSessionMap struct {
	Map map[string]*MMBGSession
}

type UsersMap struct {
	Map map[int64]*User
}
type User struct {
	BGsessionId string
}

type MMBGSession struct {
	Player1 tele.Context
	Player2 tele.Context
	Session *bg.BGSession
}

func NewMMBGSession(c1, c2 tele.Context, session *bg.BGSession) *MMBGSession {
	return &MMBGSession{
		Player1: c1,
		Player2: c2,
		Session: session,
	}
}

func InitMM() {
	InitBGQueue()
	InitMMBGSessionMap()
	InitUsersMap()
}

func InitBGQueue() {
	Queue = &BGQueue{
		Players: make([]tele.Context, 0, 100),
	}
}

func (q *BGQueue) AddUserInQueue(ctx tele.Context) {
	q.Players = append(q.Players, ctx)
}

func (q *BGQueue) IsReady() bool {
	if len(q.Players) >= 2 {
		return true
	} else {
		return false
	}
}

func InitMMBGSessionMap() {
	MapOfMMBGSession = &MMBGSessionMap{
		Map: make(map[string]*MMBGSession),
	}
}

func InitUsersMap() {
	MapOfUser = &UsersMap{
		Map: make(map[int64]*User),
	}
}
