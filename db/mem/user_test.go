package mem

import (
	"testing"

	"github.com/j4rv/golang-stuff/cah/db/mem/fixture"
)

func init() {
	fixture.PopulateUsers(userStore)
}

func TestGetUserByID(t *testing.T) {
	u, err := userStore.ByID(1)
	if err != nil {
		t.Error("Did not find user with id 1")
	} else {
		if u.Username != "Red" && u.ID == 1 {
			t.Fatal("GetUserByID is horribly broken")
		}
	}
	u, err = userStore.ByID(999)
	if err == nil {
		t.Error("Found an user with id 999 but expected none")
	}
}
