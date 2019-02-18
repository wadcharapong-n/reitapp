package api

import (
	"../services"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func GetReitAll(c echo.Context) error {
	fmt.Println("start : GetReitAll")
	results := services.GetReitAll(c)
	return c.JSON(http.StatusOK, results)
}

func GetFavoriteReitAll(c echo.Context) error {
	fmt.Println("start : GetFavoriteReitAll")
	userID := c.Param("id")
	fmt.Println("userID :: " + userID)
	results := services.GetReitFavoriteByUserID(userID)
	return c.JSON(http.StatusOK, results)
}

func SaveFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : SaveFavoriteReit")
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	services.SaveReitFavorite(userID, ticker)
	return c.String(http.StatusOK, "userId:"+userID+", ticker:"+ticker)
}
