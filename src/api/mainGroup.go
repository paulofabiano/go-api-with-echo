package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/static", handlers.HomeHandler)
	e.GET("/login", handlers.LoginHandler)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
