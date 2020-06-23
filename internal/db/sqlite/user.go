package sqlite

import (
	"errors"

	cah "github.com/j4rv/cah/internal/model"
	"gorm.io/gorm"
)

type userStore struct{}

func NewUserStore(ds cah.DataStore) *userStore {
	return &userStore{}
}

func (store *userStore) Create(username string, password []byte) (cah.User, error) {
	user := cah.User{Username: username, Password: password}
	if username == "" || len(password) == 0 {
		return user, errors.New("Username and password must not be empty")
	}
	if len(username) > 36 {
		return user, errors.New("Username too long, must not be greater than 36")
	}
	tx := db.Create(&user)
	return user, tx.Error
}

func (store *userStore) ByID(id uint) (cah.User, error) {
	var user cah.User
	if id == 0 {
		return user, errNotFound
	}
	tx := db.First(&user, cah.User{Model: gorm.Model{ID: id}})
	return user, tx.Error
}

func (store *userStore) ByName(name string) (cah.User, error) {
	var user cah.User
	if name == "" {
		return user, errNotFound
	}
	tx := db.First(&user, cah.User{Username: name})
	return user, tx.Error
}
