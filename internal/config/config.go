package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func LoadConfig() (*Config, error) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_HOST is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PORT is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_NAME is not set")
	}

	return &Config{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBName:     dbName,
		DBPort:     dbPort,
	}, nil
}
