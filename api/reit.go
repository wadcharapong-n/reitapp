package api

import (
	"../models"
	"../services"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func GetReitAll(c echo.Context) error {
	results, err := services.GetReitAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, results)
}

func GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	result, err := services.GetReitBySymbol(symbol)
	if result == (models.ReitItem{}) {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, result)
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
