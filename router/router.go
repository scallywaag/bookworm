package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func mockHandler(c echo.Context) error {
	return c.String(http.StatusOK, "This is a test response")
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}

func SetupRoutes(e *echo.Echo) {
	e.GET("/", mockHandler)
	e.POST("/login", mockHandler)
	e.POST("/register", mockHandler)

	e.GET("/ping", pong)

	e.POST("/users", mockHandler)
	e.GET("/users/:id", mockHandler)
	e.PUT("/users/:id", mockHandler)

	admin := e.Group("/admin")
	admin.GET("/stats", mockHandler)
	admin.GET("/users", mockHandler)
	admin.DELETE("/users/:id", mockHandler)
}
