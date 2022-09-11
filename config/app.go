package config

import (
	"github.com/spf13/viper"
)

var config *viper.Viper

func GetConfig() (*viper.Viper, error) {
	config = viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath("../../")

	config.AutomaticEnv()

	err := config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return config, nil
}
