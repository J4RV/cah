package mem

import (
	"errors"
	"time"

	"github.com/j4rv/golang-stuff/cah"
)

type userMemStore struct {
	abstractMemStore
	users map[int]*cah.User
}

var userStore = &userMemStore{
	users: make(map[int]*cah.User),
}

func GetUserStore() *userMemStore {
	return userStore
}

func (store *userMemStore) Create(username, password string) (cah.User, error) {
	store.Lock()
	defer store.Unlock()
	user := cah.User{}
	user.Username = username
	user.Password = password
	user.CreatedAt = time.Now()
	user.ID = store.nextID()
	store.users[user.ID] = &user
	return user, nil
}

func (store *userMemStore) ByID(id int) (cah.User, error) {
	store.Lock()
	defer store.Unlock()
	u, ok := store.users[id]
	if !ok {
		return cah.User{}, errors.New("User not found")
	}
	return *u, nil
}

func (store *userMemStore) ByName(name string) (cah.User, error) {
	store.Lock()
	defer store.Unlock()
	for _, u := range store.users {
		if u.Username == name {
			return *u, nil
		}
	}
	return cah.User{}, errors.New("User not found")
}
