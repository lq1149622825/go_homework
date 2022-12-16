package utils

import (
	"github.com/spf13/viper"
	"os"
)

var MyProxy = GetConfigInfo()

func GetConfigInfo() *viper.Viper {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return config
}
