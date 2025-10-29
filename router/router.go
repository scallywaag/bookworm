package router

import "github.com/labstack/echo/v4"

func mockHandler(c echo.Context) error {
	return nil
}

func setupRoutes(e *echo.Echo) {
	e.GET("/", mockHandler)
	e.POST("/login", mockHandler)
	e.POST("/register", mockHandler)

	e.GET("/health", mockHandler)

	e.POST("/users", mockHandler)
	e.GET("/users/:id", mockHandler)
	e.PUT("/users/:id", mockHandler)

	admin := e.Group("/admin")
	admin.GET("/stats", mockHandler)
	admin.GET("/users", mockHandler)
	admin.DELETE("/users/:id", mockHandler)
}
