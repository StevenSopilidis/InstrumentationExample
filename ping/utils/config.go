package utils

import "github.com/spf13/viper"

type Config struct {
	PongServerAddress string `mapstructure:"PONG_SERVER_ADDRESS"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	// if file is just not found means we are running from docker
	viper.AutomaticEnv()
	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("PONG_SERVER_ADDRESS")

	err = viper.Unmarshal(&config)
	return config, err
}
