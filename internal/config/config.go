package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "myuser")
	viper.SetDefault("DB_PASSWORD", "mypassword")
	viper.SetDefault("DB_NAME", "marketplace")
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("JWT_SECRET", "pelindo888")

	_ = viper.ReadInConfig()

	fmt.Println("DB_USER dari viper:", viper.GetString("DB_USER"))

	cfg := &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		ServerPort: viper.GetString("SERVER_PORT"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	}
	return cfg, nil
}
