package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configurations of the service
type Config struct {
	DBDriver                 string        `mapstructure:"DB_DRIVER"`
	DBSource                 string        `mapstructure:"DB_SOURCE"`
	ServerHost               string        `mapstructure:"SERVER_HOST"`
	ServerPort               string        `mapstructure:"SERVER_PORT"`
	TokenSecretKey           string        `mapstructure:"TOKEN_SECRET_KEY"`
	AccessTokenDuration      time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration     time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	UserServiceServerAddress string        `mapstructure:"USER_SERVICE_SERVER_ADDRESS"`
}

// LoadConfig reads all configurations from .env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
