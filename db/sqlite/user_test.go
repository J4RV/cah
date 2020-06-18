package sqlite

import (
	"testing"

	"github.com/j4rv/cah"
)

func userTestSetup(t *testing.T) (*userStore, func()) {
	InitDB(":memory:")
	CreateTables()
	return NewUserStore(cah.DataStore{}), func() {
		db.Close()
	}
}

func TestUserCreate(t *testing.T) {
	us, teardown := userTestSetup(t)
	defer teardown()
	cases := []struct {
		name        string
		username    string
		password    []byte
		errExpected bool
	}{
		{"valid", "Rojo", []byte("rojopass"), false},
		{"empty username", "", []byte("pass"), true},
		{"empty password", "Admin", []byte(""), true},
		{"Name too long", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", []byte("pass"), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := us.Create(tc.username, tc.password)
			if !tc.errExpected && err != nil {
				t.Fatal(err.Error())
			}
			if tc.errExpected && err == nil {
				t.Fatal("Expected error but found nil")
			}
			if err == nil && (u.Username != tc.username /* FIXME || u.Password != tc.password*/) {
				t.Fatalf("The user was created with wrong fields, case: %+v, got %+v", tc, u)
			}
		})
	}
}

func TestUserGetByID(t *testing.T) {
	us, teardown := userTestSetup(t)
	defer teardown()
	db.MustExec(`INSERT INTO user (username, password) VALUES
		("first", "first"), ("second", "second"), ("third", "third")`)
	cases := []struct {
		name        string
		id          int
		errExpected bool
	}{
		{"first", 1, false},
		{"last", 3, false},
		{"zero", 0, true},
		{"minus one", -1, true},
		{"large number", 99999, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := us.ByID(tc.id)
			if !tc.errExpected && err != nil {
				t.Fatal(err.Error())
			}
			/*if !tc.errExpected && (u == cah.User{}) {
				t.Fatal("Found unexpected empty user")
			}*/
			if tc.errExpected && err == nil {
				t.Fatal("Expected error but found nil")
			}
		})
	}
}

func TestUserGetByName(t *testing.T) {
	us, teardown := userTestSetup(t)
	defer teardown()
	db.MustExec(`INSERT INTO user (username, password) VALUES
		("first", "first"), ("second", "second"), ("third", "third")`)
	cases := []struct {
		name        string
		namesearch  string
		errExpected bool
	}{
		{"first", "first", false},
		{"last", "third", false},
		{"empty", "", true},
		{"non existant", "anon", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := us.ByName(tc.namesearch)
			if !tc.errExpected && err != nil {
				t.Fatal(err.Error())
			}
			/*if !tc.errExpected && (u == cah.User{}) {
				t.Fatal("Found unexpected empty user")
			}*/
			if tc.errExpected && err == nil {
				t.Fatal("Expected error but found nil")
			}
		})
	}
}
