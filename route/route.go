package route

import (
	"../api"
	"../config"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"time"
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

type Facebook struct {
	ID     string `bson:"id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
}

type Google struct {
	ID     string `bson:"id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	VerifiedEmail string `bson:"verified_email"`
	GivenName string `bson:"given_name"`
	FamilyName string `bson:"family_name"`
	Link string `bson:"link"`
	PictureURL string `bson:"picture"`
	Gender string `bson:"gender"`
	Locale string `bson:"locale"`
}

type jwtCustomClaims struct {
	ID     string `bson:"id"`
	Name string `bson:"name"`
	jwt.StandardClaims
}

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

	// Unauthenticated route
	a.GET("/", accessible)

	// API group
	r := e.Group("/api")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)
	// Routes

	r.GET("/reit", api.GetReitAll)
	r.GET("/reitFavorite/:id", api.GetFavoriteReitAll)
	r.POST("/reitFavorite", api.SaveFavoriteReit)
	r.GET("/reit/:symbol", api.GetReitBySymbol)

	e.GET("/search", api.TestElasticSearch)

	return e
}

func handleGoogleLogin(c echo.Context) error {
	fmt.Println("Start handleGoogleLogin")
	w := c.Response().Writer
	r := c.Request()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
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
	google := Google{}
	json.Unmarshal(contents,&google)
	CreateTokenFromGoogle(c,google)
	return nil
}

func handleFacebookLogin(c echo.Context) error {
	fmt.Println("Start handleFacebookLogin")
	w := c.Response().Writer
	r := c.Request()
	url := facebookOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
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
	facebook := Facebook{}
	json.Unmarshal(contents,&facebook)
	CreateTokenFromFacebook(c,facebook)
	return nil

}

func CreateTokenFromFacebook(c echo.Context, facebook Facebook) error {
	// Set custom claims
	claims := &jwtCustomClaims {
		facebook.ID,
		facebook.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func CreateTokenFromGoogle(c echo.Context, google Google) error {
	// Set custom claims
	claims := &jwtCustomClaims {
		google.ID,
		google.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
