package cah

import (
	"gorm.io/gorm"
)

type UserStore interface {
	Create(username string, password []byte) (User, error)
	ByName(name string) (User, error)
	ByID(id uint) (User, error)
}

type UserUsecases interface {
	Register(username, password string) (User, error)
	Login(name, pass string) (u User, ok bool)
	ByID(id uint) (u User, ok bool)
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password []byte `json:"-" gorm:"default:null;not null"`
}
