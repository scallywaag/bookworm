package router

import (
	"database/sql"
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
	admin.GET("/users", func(c echo.Context) error {
		db := c.Get("db").(*sql.DB)

		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer rows.Close()

		type User struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			users = append(users, u)
		}

		return c.JSON(http.StatusOK, users)
	})
	admin.DELETE("/users/:id", mockHandler)
}
