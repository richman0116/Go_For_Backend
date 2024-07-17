package handler

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal(err)
	}
}
