package mem

import "sync"

type abstractMemStore struct {
	lastID int
	sync.Mutex
}

func (s *abstractMemStore) nextID() int {
	s.lastID++
	return s.lastID
}
