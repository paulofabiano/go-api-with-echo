package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeAdminHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Admin!")
}
