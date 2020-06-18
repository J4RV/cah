package mem

import (
	"sync"

	"github.com/j4rv/cah"
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
