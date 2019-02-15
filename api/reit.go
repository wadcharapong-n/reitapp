package api

import (
	"../services"
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func GetReitAll(c echo.Context) error {
	results := services.GetReitAll()
	if results == nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, results)
}
