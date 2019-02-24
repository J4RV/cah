package mem

import (
	"errors"
	"fmt"

	"github.com/j4rv/golang-stuff/cah"
)

type gameMemStore struct {
	abstractMemStore
	games map[int]cah.Game
}

var gameStore = &gameMemStore{
	games: map[int]cah.Game{},
}

func GetGameStore() *gameMemStore {
	return gameStore
}

func (store *gameMemStore) Create(g cah.Game) error {
	store.Lock()
	defer store.Unlock()
	if g.ID != 0 {
		return errors.New("Tried to create a game but its ID was not zero")
	}
	g.ID = store.nextID()
	store.games[g.ID] = g
	return nil
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

func (store *gameMemStore) ByStatePhase(phases ...cah.Phase) []cah.Game {
	store.Lock()
	defer store.Unlock()
	ret := []cah.Game{}
	for _, g := range store.games {
		for _, p := range phases {
			if g.State.Phase == p {
				ret = append(ret, g)
				break
			}
		}
	}
	return ret
}

func (store *gameMemStore) Update(g cah.Game) error {
	store.Lock()
	defer store.Unlock()
	currG, ok := store.games[g.ID]
	if !ok {
		return fmt.Errorf("No game found with ID %d", g.ID)
	}
	store.games[g.ID] = g
	if !currG.State.Equal(g.State) {
		stateStore.Update(g.State)
	}
	return nil
}
