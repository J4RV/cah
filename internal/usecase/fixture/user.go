package fixture

import (
	cah "github.com/j4rv/cah/internal/model"
)

var users = []struct {
	name, pass string
}{
	{"Red", "Red"},
	{"Green", "Green"},
	{"Blue", "Blue"},
	{"Yellow", "Yellow"},
}

func PopulateUsers(s cah.UserUsecases) {
	for _, u := range users {
		s.Register(u.name, u.pass)
	}
}
