package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/scallywaag/bookworm/internal/database"
	"github.com/scallywaag/bookworm/internal/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	msg, err := database.Init()
	if err != nil {
		e.Logger.Fatal("Database init failed: ", err)
	}
	e.Logger.Info(msg)
	e.Use(database.DBMiddleware)

	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
