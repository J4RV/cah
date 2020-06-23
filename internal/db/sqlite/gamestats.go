package sqlite

import cah "github.com/j4rv/cah/internal/model"

type gameStatsStore struct{}

func NewGameStatsStore(ds cah.DataStore) *gameStatsStore {
	return &gameStatsStore{}
}
