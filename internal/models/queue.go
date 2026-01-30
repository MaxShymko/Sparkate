package models

import "sync"

type Queue struct {
	Users []User
	mu    sync.RWMutex
}

var QueueOfUsers *Queue

func InitQueue() {
	QueueOfUsers = &Queue{
		Users: make([]User, 0, 100),
	}
}

func (q *Queue) Push(user User) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Users = append(q.Users, user)
}

func (q *Queue) Pop() (*User, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.Users) == 0 {
		return nil, false
	}

	User := q.Users[0]
	q.Users = q.Users[1:]
	return &User, true
}
