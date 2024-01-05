package config

import (
	"github.com/Fajar3108/go-boilerplate/internal/constants"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Debug       bool   `mapstructure:"DEBUG"`
}

var AppConfig Config

func IntializeAppConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return constants.ErrLoadConfig
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return constants.ErrParseConfig
	}

	if AppConfig.Port == 0 || AppConfig.Environment == "" {
		return constants.ErrEmptyVar
	}

	return nil
}
