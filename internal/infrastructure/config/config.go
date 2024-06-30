package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	TrustedProxies []string
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
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
		DBUser:         viper.GetString("DB_USER"),
		DBPassword:     viper.GetString("DB_PASSWORD"),
		DBHost:         viper.GetString("DB_HOST"),
		DBPort:         viper.GetString("DB_PORT"),
		DBName:         viper.GetString("DB_NAME"),
		EnvMode:        configEnv,
		Port:           viper.GetString("PORT"),
		TrustedProxies: viper.GetStringSlice("TRUSTED_PROXIES"),
		Version:        LoadVersion(),
	}
	return config
}
