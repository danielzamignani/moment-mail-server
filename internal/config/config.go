package config

import "os"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func Load() *Config {
	databaseConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	serverConfig := ServerConfig{
		Port: os.Getenv("PORT"),
	}

	return &Config{
		Database: databaseConfig,
		Server:   serverConfig,
	}
}
