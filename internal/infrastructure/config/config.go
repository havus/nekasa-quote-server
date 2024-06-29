package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	TrustedProxies []string
	Version        string
	EnvMode        string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	viper.AutomaticEnv()

	configEnv := viper.GetString("CONFIG_ENV")

	config := &Config{
		Port:           viper.GetString("PORT"),
		TrustedProxies: viper.GetStringSlice("TRUSTED_PROXIES"),
		Version:        LoadVersion(),
		EnvMode:        configEnv,
	}
	return config
}
