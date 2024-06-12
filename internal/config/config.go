package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	LogLevel   string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		LogLevel:   viper.GetString("LOG_LEVEL"),
	}

	return config, nil
}
