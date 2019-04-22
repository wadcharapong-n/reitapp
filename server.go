package main

import (
	"github.com/wadcharapong/REIT_APP_API/route"
)

func main() {
	//Route
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
