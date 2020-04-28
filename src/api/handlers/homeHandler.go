package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Echo!")
}
