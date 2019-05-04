package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wadcharapong/reitapp/route"
	"strings"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	//Route
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
