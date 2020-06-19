package mem

import (
	"fmt"

	cah "github.com/j4rv/cah/internal/model"
)

type stateMemStore struct {
	abstractMemStore
	gameStates map[int]*cah.GameState
}

var stateStore = &stateMemStore{
	gameStates: make(map[int]*cah.GameState),
}

// GetGameStateStore returns the global game state store
func GetGameStateStore(ds cah.DataStore) cah.GameStateStore {
	return stateStore
}

func (store *stateMemStore) Create(gs *cah.GameState) (*cah.GameState, error) {
	store.Lock()
	defer store.Unlock()
	gs.ID = store.nextID()
	store.gameStates[gs.ID] = gs
	return gs, nil
}

func (store *stateMemStore) ByID(id int) (*cah.GameState, error) {
	store.Lock()
	defer store.Unlock()
	return store.byID(id)
}

func (store *stateMemStore) byID(id int) (*cah.GameState, error) {
	gs, ok := store.gameStates[id]
	if !ok {
		return &cah.GameState{}, fmt.Errorf("No game found with ID %d", id)
	}
	return gs, nil
}

func (store *stateMemStore) Update(gs *cah.GameState) error {
	store.Lock()
	defer store.Unlock()
	_, err := store.byID(gs.ID)
	if err != nil {
		return err
	}
	store.gameStates[gs.ID] = gs
	return nil
}

func (store *stateMemStore) Delete(id int) error {
	store.Lock()
	defer store.Unlock()
	_, err := store.byID(id)
	if err != nil {
		return err
	}
	delete(store.gameStates, id)
	return nil
}
