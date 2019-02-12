package main

import (
	"./route"
)

func main() {
	//Route
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
