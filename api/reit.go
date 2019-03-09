package api

import (
	"../models"
	"../services"
	"../util"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)



// Handler
func GetReitAll(c echo.Context) error {
	var reitService  services.ReitServicer
	reitService = services.Reit{}
	results ,err := services.GetReitAllProcess(reitService)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, results)
}

func GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	var reitService services.ReitServicer
	reitService = services.Reit{}
	result, err := services.GetReitBySymbolProcess(reitService,symbol)
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
	var reitService services.ReitServicer
	reitService = services.Reit{}
	results := services.GetReitFavoriteByUserIDJoinProcess(reitService,userID)
	return c.JSON(http.StatusOK, results)
}

func SaveFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : SaveFavoriteReit")
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	var reitService services.ReitServicer
	reitService = services.Reit{}
	err := services.SaveReitFavoriteProcess(reitService, userID, ticker)
	if err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func DeleteFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : DeleteFavoriteReit")
	userID := c.FormValue("userId")
	ticker := c.FormValue("Ticker")
	var reitService services.ReitServicer
	reitService = services.Reit{}
	err := services.DeleteReitFavoriteProcess(reitService, userID, ticker)
	if err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func GetUserProfile(c echo.Context) error {
	userID,site := util.GetUserFromToken(c);
	profile := services.GetUserProfileByCriteria(userID, site)
	return c.JSON(http.StatusOK, profile)
}
