package sqlite

import (
	"testing"

	cah "github.com/j4rv/cah/internal/model"
)

func userTestSetup(t *testing.T) (*userStore, func()) {
	InitDB(":memory:")
	CreateTables()
	return NewUserStore(cah.DataStore{}), func() {}
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
		{"repeated username", "Rojo", []byte("rojopass"), true},
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
				t.Fatal("Expected error but found nil", u.Username)
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
	db.Create(&cah.User{Username: "first", Password: []byte("first")})
	db.Create(&cah.User{Username: "second", Password: []byte("second")})
	db.Create(&cah.User{Username: "third", Password: []byte("third")})
	cases := []struct {
		name        string
		id          uint
		errExpected bool
	}{
		{"first", 1, false},
		{"last", 3, false},
		{"zero", 0, true},
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
	db.Create(&cah.User{Username: "first", Password: []byte("first")})
	db.Create(&cah.User{Username: "second", Password: []byte("second")})
	db.Create(&cah.User{Username: "third", Password: []byte("third")})
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
