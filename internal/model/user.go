package cah

import (
	"time"
)

type UserStore interface {
	Create(username string, password []byte) (User, error)
	ByName(name string) (User, error)
	ByID(id int) (User, error)
}

type UserUsecases interface {
	Register(username, password string) (User, error)
	Login(name, pass string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type User struct {
	ID        int       `json:"id"       db:"user_account"`
	Username  string    `json:"username" db:"username"`
	Password  []byte    `json:"-"        db:"password"`
	CreatedAt time.Time `json:"-"        db:"created_at"`
}
