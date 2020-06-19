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

// Passwords will be plaintext since Usecase is the one doing the hashing!
func PopulateUsers(s cah.UserStore) {
	for _, u := range users {
		s.Create(u.name, []byte(u.pass))
	}
}
