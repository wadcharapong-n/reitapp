package config

var RedirectURL_Google = "http://localhost:1323/Auth/GoogleCallback"
var ClientID_Google = "727520883795-jgv79pgu957lu3oojf8cv30akgbeep7j.apps.googleusercontent.com"
var ClientSecret_Google = "C5Y7I3MYKuVoaNBczMn8qoub"
var Scopes_Google = []string{"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email"}
var URL_access_token_Google = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var RedirectURL_Facebook = "http://localhost:1323/Auth/FacebookCallback"
var ClientID_Facebook = "385086385375815"
var ClientSecret_Facebook = "baa3dce5c60b202469c4387f0b588766"
var Scopes_Facebook = []string{"public_profile", "email"}
var URL_access_token_Facebook = "https://graph.facebook.com/me?access_token="
