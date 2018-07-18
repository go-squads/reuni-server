package config
import (
	"os"
)
type databaseConfig struct {
	host        string
	port        string
	username    string
	password    string
	name        string
}

func getDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		name: os.Getenv("DB_NAME"),
	}	
}

