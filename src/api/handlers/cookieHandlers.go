package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CookieHandler(c echo.Context) error {
	return c.String(http.StatusOK, "The cookie group")
}
