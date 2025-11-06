package database

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/scallywaag/bookworm/internal/persistence"
)

func DBMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("db", DB)
		c.Set("queries", persistence.New(DB))
		return next(c)
	}
}

func GetDB(c echo.Context) *sql.DB {
	db, ok := c.Get("db").(*sql.DB)
	if !ok || db == nil {
		c.Logger().Error("Database not found in context")
		return nil
	}
	return db
}

func GetQueries(c echo.Context) *persistence.Queries {
	q, ok := c.Get("queries").(*persistence.Queries)
	if !ok || q == nil {
		c.Logger().Error("Queries not found in context")
		return nil
	}
	return q
}
