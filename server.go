package main

import (
	"github.com/wadcharapong/reitapp/route"
)

func main() {
	//viper.SetConfigName("config") // ชื่อ config file
	//	//viper.AddConfigPath(".") // ระบุ path ของ config file
	//	//viper.AutomaticEnv() // อ่าน value จาก ENV variable
	//	//// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	//	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//	//// อ่าน config
	//	//err := viper.ReadInConfig()
	//	//if err != nil {
	//	//	panic(fmt.Errorf("fatal error config file: %s \n", err))
	//	//}
	//Route
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
