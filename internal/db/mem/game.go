package mem

import (
	"errors"
	"fmt"

	cah "github.com/j4rv/cah/internal/model"
)

type gameMemStore struct {
	abstractMemStore
	games map[int]cah.Game
}

var gameStore = &gameMemStore{
	games: map[int]cah.Game{},
}

// GetGameStore returns the global game store
func GetGameStore(ds cah.DataStore) cah.GameStore {
	gameStore.dataStore = ds
	return gameStore
}

func (store *gameMemStore) Create(g cah.Game) (cah.Game, error) {
	store.Lock()
	defer store.Unlock()
	if g.ID != 0 {
		return cah.Game{}, errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return g, nil
}

func (store *gameMemStore) ByID(id int) (cah.Game, error) {
	store.Lock()
	defer store.Unlock()
	g, ok := store.games[id]
	if !ok {
		return g, fmt.Errorf("No game found with id %d", id)
	}
	return g, nil
}

func (store *gameMemStore) ByGameStateID(id int) (cah.Game, error) {
	store.Lock()
	defer store.Unlock()
	for _, g := range store.games {
		if g.StateID == id {
			return g, nil
		}
	}
	return cah.Game{}, fmt.Errorf("No game found with game state id %d", id)
}

func (store *gameMemStore) ByPhase(started, finished bool) ([]cah.Game, error) {
	store.Lock()
	defer store.Unlock()
	ret := []cah.Game{}
	for _, g := range store.games {
		if g.Started == started && g.Finished == finished {
			ret = append(ret, g)
		}
	}
	return ret, nil
}

func (store *gameMemStore) Update(g cah.Game) error {
	store.Lock()
	defer store.Unlock()
	store.games[g.ID] = g
	return nil
}
