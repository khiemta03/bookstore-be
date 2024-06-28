package config

import "github.com/spf13/viper"

// Config stores all configurations of the service
type Config struct {
	ServerHost                   string `mapstructure:"SERVER_HOST"`
	ServerPort                   string `mapstructure:"SERVER_PORT"`
	UserServiceAddress           string `mapstructure:"USER_SERVICE_ADDRESS"`
	AuthenticationServiceAddress string `mapstructure:"AUTHENTICATION_SERVICE_ADDRESS"`
	BookServiceAddress           string `mapstructure:"BOOK_SERVICE_ADDRESS"`
	OrderServiceAddress          string `mapstructure:"ORDER_SERVICE_ADDRESS"`
}

// LoadConfig reads all configurations from .env file
func LoadConfig(path string, name string, ext string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(ext)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
