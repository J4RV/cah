package sqlite

import (
	"github.com/jmoiron/sqlx"

	// driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB(dbFileName string) {
	db = sqlx.MustOpen("sqlite3", dbFileName)
	if db.Ping() != nil {
		panic("DB did not answer ping")
	}
	CreateTables()
}

func CreateTables() {
	createTableUser()
}
