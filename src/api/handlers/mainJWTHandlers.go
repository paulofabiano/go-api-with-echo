package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func MainJwtHandler(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	if claim, ok := token.Claims.(jwt.MapClaims); !ok {
		log.Println("error typecasting")
	} else {
		log.Printf("user name: %v | user id: %v\n", claim["name"], claim["jti"])
	}

	return c.String(http.StatusOK, "The jwt group")
}
