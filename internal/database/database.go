package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() (string, error) {
	var err error
	DB, err = sql.Open("sqlite3", "file:bookworm.db?_foreign_keys=on")
	if err != nil {
		return "", err
	}

	if err = DB.Ping(); err != nil {
		return "", err
	}

	return "Database initialized successfully", nil
}
