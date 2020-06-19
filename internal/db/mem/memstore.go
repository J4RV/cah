package mem

import (
	"sync"

	cah "github.com/j4rv/cah/internal/model"
)

type abstractMemStore struct {
	lastID    int
	dataStore cah.DataStore
	sync.Mutex
}

func (s *abstractMemStore) nextID() int {
	s.lastID++
	return s.lastID
}
