package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PostUser Handler
func PostUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
