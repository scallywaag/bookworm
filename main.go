package main

import (
	"github.com/labstack/echo/v4"
	"github.com/scallywaag/bookworm-echo/router"
)

func main() {
	e := echo.New()
	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
