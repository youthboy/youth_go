package model

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config struct {
}

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	currentPath, _ := os.Getwd()
	viper.AddConfigPath(".")

	fmt.Println(currentPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Config Error : %s", err.Error()))
	}
}
