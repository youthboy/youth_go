package main

import (
	"fmt"
	"github.com/spf13/viper"
	"youthmod/config"
)

func main() {
	config.InitConfig()
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Print("这是我第一个GO程序")

}
