package route

import (
	"../api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/reit", api.GetReitAll)
	e.GET("/search", api.TestElasticSearch)
	return e
}
