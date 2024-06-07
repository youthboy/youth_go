package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config struct {
}

func InitConfig() {
	viper.SetConfigFile("./setting/config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Config Error : %s", err.Error()))
	}
}
