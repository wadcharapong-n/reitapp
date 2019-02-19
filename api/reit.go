package api

import (
	"../models"
	"../services"
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
