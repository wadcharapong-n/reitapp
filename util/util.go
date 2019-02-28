package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"../models"
)

func GetUserFromToken(c echo.Context) (string,string)  {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTCustomClaims)
	userID := claims.ID
	site := claims.Site
	return userID,site
}