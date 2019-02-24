package mem

import (
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

type stateMemStore struct {
	abstractMemStore
	games map[int]*cah.GameState
}

var stateStore = &stateMemStore{
	games: make(map[int]*cah.GameState),
}

func GetGameStateStore() *stateMemStore {
	return stateStore
}

func (store *stateMemStore) Create(g cah.GameState) (cah.GameState, error) {
	store.Lock()
	defer store.Unlock()
	g.ID = store.nextID()
	store.games[g.ID] = &g
	return g, nil
}

func (store *stateMemStore) ByID(id int) (cah.GameState, error) {
	store.Lock()
	defer store.Unlock()
	return store.byID(id)
}

func (store *stateMemStore) byID(id int) (cah.GameState, error) {
	g, ok := store.games[id]
	if !ok {
		return cah.GameState{}, fmt.Errorf("No game found with ID %d", id)
	}
	return *g, nil
}

func (store *stateMemStore) Update(g cah.GameState) error {
	store.Lock()
	defer store.Unlock()
	_, err := store.byID(g.ID)
	if err != nil {
		return err
	}
	store.games[g.ID] = &g
	return nil
}

func (store *stateMemStore) Delete(id int) error {
	store.Lock()
	defer store.Unlock()
	_, err := store.byID(id)
	if err != nil {
		return err
	}
	delete(store.games, id)
	return nil
}
