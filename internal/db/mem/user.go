package mem

import (
	"errors"
	"time"

	cah "github.com/j4rv/cah/internal/model"
)

type userMemStore struct {
	abstractMemStore
	users map[uint]*cah.User
}

var userStore = &userMemStore{
	users: make(map[uint]*cah.User),
}

// GetUserStore returns the global user store
func GetUserStore(ds cah.DataStore) cah.UserStore {
	return userStore
}

func (store *userMemStore) Create(username string, password []byte) (cah.User, error) {
	store.Lock()
	defer store.Unlock()
	user := cah.User{}
	user.Username = username
	user.Password = password
	user.CreatedAt = time.Now()
	user.ID = uint(store.nextID())
	store.users[user.ID] = &user
	return user, nil
}

func (store *userMemStore) ByID(id uint) (cah.User, error) {
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
