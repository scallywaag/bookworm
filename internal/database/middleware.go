package database

import "github.com/labstack/echo/v4"

func DBMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("db", DB)
		return next(c)
	}
}
