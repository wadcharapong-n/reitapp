package main

import (
	"github.com/wadcharapong/reitapp/route"
)

func main() {
	//Route
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
