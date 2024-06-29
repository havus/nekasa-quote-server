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
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	viper.AutomaticEnv()

	config := &Config{
		Port:           viper.GetString("PORT"),
		TrustedProxies: viper.GetStringSlice("TRUSTED_PROXIES"),
		Version:        LoadVersion(),
	}
	return config
}
