package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wadcharapong/reitapp/api"
	"github.com/wadcharapong/reitapp/models"
)

func Init() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))


	reitController := api.Reit_Handler {}
	authController := api.Auth_Handler {}

	//Authenticate
	//e.GET("/", handleMain)
	a := e.Group("/Auth")
	a.GET("/GoogleLogin", authController.HandleGoogleLogin)
	a.GET("/GoogleCallback", authController.HandleGoogleCallback)

	a.GET("/FacebookLogin", authController.HandleFacebookLogin)
	a.GET("/FacebookCallback", authController.HandleFacebookCallback)

	a.GET("", authController.Authentication)

	// Unauthenticated route
	//a.GET("/", accessible)

	// API group
	r := e.Group("/api")
	//Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &models.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// Routes
	r.GET("/reit", reitController.GetReitAll)
	r.GET("/reitFavorite", reitController.GetFavoriteReitAll)
	r.POST("/reitFavorite", reitController.SaveFavoriteReit)
	r.DELETE("/reitFavorite", reitController.DeleteFavoriteReit)
	r.GET("/reit/:symbol", reitController.GetReitBySymbol)
	r.GET("/profile", reitController.GetUserProfile)
	r.GET("/refreshToken",authController.RefreshToken)

	r.GET("/search", reitController.Search)
	r.GET("/searchMap", reitController.SearchMap)
	r.GET("/syncElastic", reitController.SynData)
	r.POST("/addReit",reitController.AddReit)
	
	return e
}


