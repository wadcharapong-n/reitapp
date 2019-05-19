package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/wadcharapong/reitapp/models"
	"github.com/wadcharapong/reitapp/services"
)

type AdminController interface {
	HandleAdminLogin(c echo.Context) error
	HandleGetUserAll(c echo.Context) error
	HandleDeleteUser(c echo.Context) error
	HandleGetPlaceAll(c echo.Context) error
	HandleaAddPlace(c echo.Context) error
	HandleDeletePlace(c echo.Context) error
	HandleGetReitAll(c echo.Context) error
	HandleGetPlaceID(c echo.Context) error
}

type Admin_Handler struct {
	c            echo.Context
	userProfiles []models.UserProfile
	userProfile  models.UserProfile
	Place        models.Place
	Places       []models.Place
	ReitItems    []models.ReitItem
	reitServicer services.Reit_Service
	err          error
}

func (self Admin_Handler) HandleAdminLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "password" {
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusBadRequest, "fail")
}

func (self Admin_Handler) HandleGetUserAll(c echo.Context) error {
	self.userProfiles, self.err = self.reitServicer.GetUserAll()

	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.userProfiles)
}

func (self Admin_Handler) HandleDeleteUser(c echo.Context) error {
	userId := c.FormValue("userId")
	self.err = self.reitServicer.DeleteUser(userId)

	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Admin_Handler) HandleGetPlaceAll(c echo.Context) error {
	self.Places, self.err = self.reitServicer.GetPlaceAll()

	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.Places)
}

func (self Admin_Handler) HandleaAddPlace(c echo.Context) error {

	placeId := c.FormValue("placeId")
	name := c.FormValue("name")
	address := c.FormValue("address")
	symbol := c.FormValue("symbol")
	lat, _ := strconv.ParseFloat(c.FormValue("lat"), 64)
	long, _ := strconv.ParseFloat(c.FormValue("long"), 64)

	if len(strings.TrimSpace(placeId)) == 0 ||
		len(strings.TrimSpace(name)) == 0 ||
		len(strings.TrimSpace(address)) == 0 ||
		len(strings.TrimSpace(symbol)) == 0 ||
		len(strings.TrimSpace(c.FormValue("lat"))) == 0 ||
		len(strings.TrimSpace(c.FormValue("long"))) == 0 {

		return c.String(http.StatusBadRequest, "fail")
	}

	self.err = self.reitServicer.AddPlace(placeId, name, address, symbol, lat, long)

	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")

}

func (self Admin_Handler) HandleDeletePlace(c echo.Context) error {
	id := c.FormValue("id")
	self.err = self.reitServicer.DeletePlace(id)

	if self.err != nil {
		return c.String(http.StatusBadRequest, "fail")
	}
	return c.String(http.StatusOK, "success")
}

func (self Admin_Handler) HandleGetReitAll(c echo.Context) error {
	self.ReitItems, self.err = self.reitServicer.GetReitAll()

	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.ReitItems)
}

func (self Admin_Handler) HandleGetPlaceID(c echo.Context) error {
	id := c.Param("id")
	self.Place, self.err = self.reitServicer.GetPlaceID(id)

	if self.err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}
	return c.JSON(http.StatusOK, self.Place)
}
