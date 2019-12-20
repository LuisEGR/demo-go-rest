package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteUser Handler
func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
