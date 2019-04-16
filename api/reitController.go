package api

import (
	"../models"
	"../services"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type ReitController interface {
	GetReitAll(c echo.Context) error
	GetReitBySymbol(c echo.Context) error
	GetFavoriteReitAll(c echo.Context) error
	SaveFavoriteReit(c echo.Context) error
	DeleteFavoriteReit(c echo.Context) error
	GetUserProfile(c echo.Context) error
	GetUserFromToken(c echo.Context) (string,string)
	Search(c echo.Context) error
  	SynData(c echo.Context) error
}

type Reit_Handler struct {
	reitServicer services.Reit_Service
	reitItems []models.ReitItem
	reitItem models.ReitItem
	reitFavorite []*models.FavoriteInfo
	err error
	authHandler Auth_Handler
}

// Handler
func (self Reit_Handler) GetReitAll(c echo.Context) error {

	self.reitItems ,self.err = self.reitServicer.GetReitAll()
	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.reitItems)
}

func (self Reit_Handler) GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	self.reitItem, self.err = self.reitServicer.GetReitBySymbol(symbol)
	if self.reitItem.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	if self.err != nil {
		return echo.NewHTTPError(http.StatusOK, self.err)
	}
	return c.JSON(http.StatusOK, self.reitItem)
}

func (self Reit_Handler) GetFavoriteReitAll(c echo.Context) error {
	fmt.Println("start : GetFavoriteReitAll")
	//userID := c.Param("id")

	userID,_ := self.authHandler.GetUserFromToken(c);
	self.reitServicer = services.Reit_Service{}
	self.reitFavorite = self.reitServicer.GetReitFavoriteByUserIDJoin(userID)
	return c.JSON(http.StatusOK, self.reitFavorite)
}

func (self Reit_Handler) SaveFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : SaveFavoriteReit")
	//userID := c.FormValue("userId")
	userID,_ := self.authHandler.GetUserFromToken(c);
	ticker := c.FormValue("Ticker")
	self.err = self.reitServicer.SaveReitFavorite( userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Reit_Handler) DeleteFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : DeleteFavoriteReit")
	//userID := c.FormValue("userId")

	userID,_ := self.authHandler.GetUserFromToken(c);
	ticker := c.FormValue("Ticker")
	self.err = self.reitServicer.DeleteReitFavorite(userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Reit_Handler) GetUserProfile(c echo.Context) error {

	userID,site := self.authHandler.GetUserFromToken(c);
	profile := self.reitServicer.GetUserProfileByCriteria(userID, site)
	return c.JSON(http.StatusOK, profile)
}


func (self Reit_Handler) Search(c echo.Context) error {
	q := c.QueryParam("query")
	results := self.reitServicer.SearchElastic(q)
	return c.JSON(http.StatusOK, results)

}

func (self Reit_Handler) SearchMap(c echo.Context) error {
	latQ := c.QueryParam("lat")
	lonQ := c.QueryParam("lon")
	if len(latQ) > 0 && len(lonQ) > 0 {
		lat,_ := strconv.ParseFloat(latQ,64)
		lon,_ := strconv.ParseFloat(lonQ,64)
		results := self.reitServicer.SearchMap(lat,lon)
		//results := services.SearchMapV2(lat,lon)
		return c.JSON(http.StatusOK, results)
	}
	return c.JSON(http.StatusBadRequest,"" )

}
func (self Reit_Handler)  SynData(c echo.Context) error {
	reitItems ,err := self.reitServicer.GetReitAll()
	if err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	for _, reit := range reitItems {
		services.AddDataElastic(reit)
	}
	return c.String(http.StatusOK, "success")
}

