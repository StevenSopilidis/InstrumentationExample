package utils

import "github.com/spf13/viper"

type Config struct {
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	TracingEndpoint string `mapstructure:"TRACING_ENDPOINT"`
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
	viper.BindEnv("SERVICE_NAME")
	viper.BindEnv("TRACING_ENDPOINT")

	err = viper.Unmarshal(&config)
	return config, err
}
