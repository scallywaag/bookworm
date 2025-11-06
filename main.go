package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/scallywaag/bookworm/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "bookworm.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT);`

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialized")

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
