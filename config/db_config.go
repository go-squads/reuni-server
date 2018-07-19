package config

import (
	"fmt"
	"os"
)

type databaseConfig struct {
	host     string
	port     string
	username string
	password string
	name     string
}

const (
	SSL_MODE = "disable"
)

func getDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		name:     os.Getenv("DB_NAME"),
	}
}

func GetConnectionString() string {
	dbConfig := getDatabaseConfig()
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbConfig.username, dbConfig.password, dbConfig.name, dbConfig.host, dbConfig.port, SSL_MODE)
}
