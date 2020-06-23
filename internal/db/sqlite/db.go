package sqlite

import (
	"errors"

	cah "github.com/j4rv/cah/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var errNotFound = errors.New("Not found")

func InitDB(dbFileName string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dbFileName), nil)
	if err != nil {
		panic(err)
	}
	CreateTables()
}

func CreateTables() {
	err := db.AutoMigrate(&cah.User{})
	if err != nil {
		panic(err)
	}
}
