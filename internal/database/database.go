package database

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() (string, error) {
	var err error
	DB, err = sql.Open("sqlite3", "bookworm.db")
	if err != nil {
		return "", err
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT
    );`
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		return "", err
	}

	return "Database initialized successfully", nil
}

func GetDB(c echo.Context) *sql.DB {
	db, ok := c.Get("db").(*sql.DB)
	if !ok || db == nil {
		c.Logger().Error("Database not found in context")
		return nil
	}
	return db
}
