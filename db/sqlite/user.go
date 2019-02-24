package sqlite

import (
	"github.com/j4rv/golang-stuff/cah"
)

type userStore struct{}

func NewUserStore() *userStore {
	return &userStore{}
}

func (store *userStore) Create(username, password string) (cah.User, error) {
	var user cah.User
	_, err := db.Exec(`INSERT INTO user (username, password) VALUES (?, ?)`,
		username, password)
	if err != nil {
		return user, err
	}
	err = db.Get(&user, `SELECT * FROM user WHERE user = last_insert_rowid()`)
	return user, err
}

func (store *userStore) ByID(id int) (cah.User, error) {
	res := cah.User{}
	if err := db.Get(&res, "SELECT * FROM user WHERE user = ?", id); err != nil {
		return res, err
	}
	return res, nil
}

func (store *userStore) ByName(name string) (cah.User, error) {
	res := cah.User{}
	if err := db.Get(&res, "SELECT * FROM user WHERE username = ?", name); err != nil {
		return res, err
	}
	return res, nil
}
