package sqlite

import (
	"fmt"
)

func createTableUser() {
	db.MustExec(`CREATE TABLE IF NOT EXISTS user (
			user INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password VARBINARY(1024) NOT NULL DEFAULT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CHECK(username <> '' AND LENGTH(password) <> 0 AND LENGTH(username) <= 36)
		);`)	
	createIndex("user", "username")
}

// methods for repetitive stuff

func createIndex(table, column string) {
	indexName := fmt.Sprintf("%s_%s", table, column)
	// Using Sprintf since this internal method does not use user inputs
	createIndexStatement := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s);", indexName, table, column)
	db.MustExec(createIndexStatement)
}
