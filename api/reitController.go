package api

import (
	"../models"
	"../services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

type ReitController interface {
	GetReitAll(c echo.Context) error
	GetReitBySymbol(c echo.Context) error
	GetFavoriteReitAll(c echo.Context) error
	SaveFavoriteReit(c echo.Context) error
	DeleteFavoriteReit(c echo.Context) error
	GetUserProfile(c echo.Context) error
	GetUserFromToken(c echo.Context) (string,string)
}

type Reit struct {
	reitServicer services.Reit_Service
	reitItems []*models.ReitItem
	reitItem models.ReitItem
	reitFavorite []*models.FavoriteInfo
	err error
}

// Handler
func (self Reit) GetReitAll(c echo.Context) error {

	self.reitItems ,self.err = self.reitServicer.GetReitAll()
	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.reitItems)
}

func (self Reit) GetReitBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	self.reitItem, self.err = self.reitServicer.GetReitBySymbol(symbol)
	if self.reitItem == (models.ReitItem{}) {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	if self.err != nil {
		return echo.NewHTTPError(http.StatusOK, self.err)
	}
	return c.JSON(http.StatusOK, self.reitItem)
}

func (self Reit) GetFavoriteReitAll(c echo.Context) error {
	fmt.Println("start : GetFavoriteReitAll")
	//userID := c.Param("id")
	var reitController ReitController
	reitController = Reit{}
	userID,_ := reitController.GetUserFromToken(c);
	self.reitServicer = services.Reit_Service{}
	self.reitFavorite = self.reitServicer.GetReitFavoriteByUserIDJoin(userID)
	return c.JSON(http.StatusOK, self.reitFavorite)
}

func (self Reit) SaveFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : SaveFavoriteReit")
	//userID := c.FormValue("userId")
	var reitController ReitController
	reitController = Reit{}
	userID,_ := reitController.GetUserFromToken(c);
	ticker := c.FormValue("Ticker")
	self.err = self.reitServicer.SaveReitFavorite( userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Reit) DeleteFavoriteReit(c echo.Context) error {
	// Get name and email
	fmt.Println("start : DeleteFavoriteReit")
	//userID := c.FormValue("userId")
	var reitController ReitController
	reitController = Reit{}
	userID,_ := reitController.GetUserFromToken(c);
	ticker := c.FormValue("Ticker")
	self.err = self.reitServicer.DeleteReitFavorite(userID, ticker)
	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Reit) GetUserProfile(c echo.Context) error {
	var reitController ReitController
	reitController = Reit{}
	userID,site := reitController.GetUserFromToken(c);
	profile := self.reitServicer.GetUserProfileByCriteria(userID, site)
	return c.JSON(http.StatusOK, profile)
}

func (self Reit) GetUserFromToken(c echo.Context) (string,string)  {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTCustomClaims)
	userID := claims.ID
	site := claims.Site
	return userID,site
}

func TestElasticSearch(c echo.Context) error {
	q := c.QueryParam("query")
	results := services.SearchElastic(q)

	return c.JSON(http.StatusOK, results)

}

func  SynData(c echo.Context) error {
	reitServicer := services.Reit_Service{}
	reitItems ,err := reitServicer.GetReitAll()
	if err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	for _, reit := range reitItems {
		services.AddDataElastic(reit)
	}
	return c.String(http.StatusOK, "success")
}

