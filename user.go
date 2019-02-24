package cah

import (
	"time"
)

type UserStore interface {
	Create(username, password string) (User, error)
	ByName(name string) (User, error)
	ByID(id int) (User, error)
}

type UserUsecases interface {
	Register(username, password string) (User, error)
	Login(name, pass string) (u User, ok bool)
	ByID(id int) (u User, ok bool)
}

type User struct {
	ID        int       `json:"id" db:"user"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	// Games played, games won, cards played...
}
