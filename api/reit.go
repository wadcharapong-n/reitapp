package api

import (
	"../models"
	"../services"
	"../util"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type ReitController interface {
	GetReitAll(c echo.Context) error
	GetReitBySymbol(c echo.Context) error
	GetFavoriteReitAll(c echo.Context) error
	SaveFavoriteReit(c echo.Context) error
	DeleteFavoriteReit(c echo.Context) error
}

type Reit struct {
	services.Reit
}


func DeleteFavoriteReitProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.DeleteFavoriteReit(c)
}

func GetReitAllProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetReitAll(c)
}

func GetReitBySymbolProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetReitBySymbol(c)
}

func GetFavoriteReitAllProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.GetFavoriteReitAll(c)
}

func SaveFavoriteReitProcess(c echo.Context) error {
	var reitController ReitController
	reitController = Reit {}
	return reitController.SaveFavoriteReit(c)
}

// Handler
func (self Reit) GetReitAll(c echo.Context) error {
	var reitService  services.ReitServicer
	reitService = self.Reit
	results ,err := services.GetReitAllProcess(reitService)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, results)
}

func (self Reit) GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	var reitService services.ReitServicer
	reitService = self.Reit
	result, err := services.GetReitBySymbolProcess(reitService,symbol)
	if result == (models.ReitItem{}) {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (self Reit) GetFavoriteReitAll(c echo.Context) error {
	fmt.Println("start : GetFavoriteReitAll")
	userID := c.Param("id")
	var reitService services.ReitServicer
	reitService = services.Reit{}
	results := services.GetReitFavoriteByUserIDJoinProcess(reitService,userID)
	return c.JSON(http.StatusOK, results)
}

func (self Reit) SaveFavoriteReit(c echo.Context) error {
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

func (self Reit) DeleteFavoriteReit(c echo.Context) error {
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
