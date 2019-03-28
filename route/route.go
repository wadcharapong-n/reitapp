package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../api"
	"../config"
	"../models"
	"../services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  config.RedirectURL_Google,
		ClientID:     config.ClientID_Google,
		ClientSecret: config.ClientSecret_Google,
		Scopes:       config.Scopes_Google,
		Endpoint:     google.Endpoint,
	}

	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  config.RedirectURL_Facebook,
		ClientID:     config.ClientID_Facebook,
		ClientSecret: config.ClientSecret_Facebook,
		Scopes:       config.Scopes_Facebook,
		Endpoint:     facebook.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "random"
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

	//Authenticate
	//e.GET("/", handleMain)
	a := e.Group("/Auth")
	a.GET("/GoogleLogin", handleGoogleLogin)
	a.GET("/GoogleCallback", handleGoogleCallback)

	a.GET("/FacebookLogin", handleFacebookLogin)
	a.GET("/FacebookCallback", handleFacebookCallback)

	a.GET("", authentication)

	// Unauthenticated route
	a.GET("/", accessible)

	// API group
	r := e.Group("/api")
	//Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &models.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// Routes
	r.GET("/reit", api.GetReitAllProcess)
	r.GET("/reitFavorite/", api.GetFavoriteReitAllProcess)
	r.POST("/reitFavorite", api.SaveFavoriteReitProcess)
	r.DELETE("/reitFavorite", api.DeleteFavoriteReitProcess)
	r.GET("/reit/:symbol", api.GetReitBySymbolProcess)
	r.GET("/profile", api.GetUserProfileProcess)
	r.GET("/refreshToken",refreshToken)

	r.GET("/search", api.TestElasticSearch)

	return e
}

func handleGoogleLogin(c echo.Context) error {
	fmt.Println("Start handleGoogleLogin")
	w := c.Response().Writer
	r := c.Request()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return c.String(http.StatusUnauthorized,"")
}

func handleGoogleCallback(c echo.Context) error {
	fmt.Println("Start handleGoogleCallback")
	w := c.Response().Writer
	r := c.Request()
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	response, err := http.Get(config.URL_access_token_Google + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err == nil {
		google := models.Google{}
		json.Unmarshal(contents, &google)
		var reitServicer services.ReitServicer
		reitServicer = services.Reit_Service{}
		services.CreateNewUserProfileProcess(reitServicer,models.Facebook{}, google)
		CreateTokenFromGoogle(c, google)
	}
	return c.String(http.StatusUnauthorized,"")
}

func handleFacebookLogin(c echo.Context) error {
	fmt.Println("Start handleFacebookLogin")
	w := c.Response().Writer
	r := c.Request()
	url := facebookOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return c.String(http.StatusUnauthorized,"")
}

func handleFacebookCallback(c echo.Context) error {
	fmt.Println("Start handleFacebookCallback")
	w := c.Response().Writer
	r := c.Request()
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	code := r.FormValue("code")
	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	response, err := http.Get(config.URL_access_token_Facebook + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err == nil {
		facebook := models.Facebook{}
		json.Unmarshal(contents, &facebook)
		var reitServicer services.ReitServicer
		reitServicer = services.Reit_Service{}
		services.CreateNewUserProfileProcess(reitServicer,facebook, models.Google{})
		CreateTokenFromFacebook(c, facebook)
	}
	return c.String(http.StatusUnauthorized,"")

}

func CreateTokenFromFacebook(c echo.Context, facebook models.Facebook) error {
	// Set custom claims
	claims := &models.JWTCustomClaims{
		facebook.ID,
		facebook.Name,
		"facebook",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	claimsRefresh := &models.JWTCustomClaims{
		facebook.ID,
		facebook.Name,
		"facebook",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	r, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
		"refreshToken" : r,
	})
}

func CreateTokenFromGoogle(c echo.Context, google models.Google) error {
	// Set custom claims
	claims := &models.JWTCustomClaims{
		google.ID,
		google.Name,
		"google",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	claimsRefresh := &models.JWTCustomClaims{
		google.ID,
		google.Name,
		"google",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	r, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
		"refreshToken" : r,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func authentication(c echo.Context) error {
	token := c.QueryParam("token")
	site := c.QueryParam("site")

	if site == "facebook" {
		getProfileFacebook(token,c)
	}else if site == "google" {
		getProfileGoogle(token,c)
	}

	return c.String(http.StatusUnauthorized, "")
}

func getProfileFacebook(token string,c echo.Context) error{
	response, err := http.Get(config.URL_access_token_Facebook + token)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err == nil {
		facebook := models.Facebook{}
		json.Unmarshal(contents, &facebook)
		var reitServicer services.ReitServicer
		reitServicer = services.Reit_Service{}
		services.CreateNewUserProfileProcess(reitServicer,facebook, models.Google{})
		CreateTokenFromFacebook(c, facebook)
	}
	return c.String(http.StatusUnauthorized, "")
}

func getProfileGoogle(token string,c echo.Context) error{
	response, err := http.Get(config.URL_access_token_Google + token)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err == nil {
		google := models.Google{}
		json.Unmarshal(contents, &google)
		var reitServicer services.ReitServicer
		reitServicer = services.Reit_Service{}
		services.CreateNewUserProfileProcess(reitServicer,models.Facebook{}, google)
		CreateTokenFromGoogle(c, google)
	}
	return c.String(http.StatusUnauthorized, "")

}

func refreshToken(c echo.Context)  error{
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTCustomClaims)
	userID := claims.ID
	site := claims.Site
	name := claims.Name

	// Set custom claims
	newClaims := &models.JWTCustomClaims{
		userID,
		name,
		site,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
