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

func TestElasticSearch(c echo.Context) error {

	results := services.SearchElastic()

	return c.JSON(http.StatusOK, results)

}
