package cmd

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr string

	PostgresConn string
	RedisConn    string
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	config := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	viper.SetEnvPrefix("BOILERPLATE")
	viper.KeyDelimiter("__")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
