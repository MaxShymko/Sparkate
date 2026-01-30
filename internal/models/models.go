package models

type User struct {
	Status string
}

type UserMap struct {
	Map map[int64]User
}

var userMap *UserMap
