package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConfig(t *testing.T) {
	configVars := map[string]string{
		"DB_HOST": "127.0.0.1",
		"DB_PORT": "5432",
		"DB_USER": "reuni_test",
		"DB_PASS": "reuni_test123!",
		"DB_NAME": "reuni_test",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	dbConfig := getDatabaseConfig()
	assert.Equal(t, dbConfig.host, configVars["DB_HOST"])
	assert.Equal(t, dbConfig.password, configVars["DB_PASS"])
	assert.Equal(t, dbConfig.username, configVars["DB_USER"])
	assert.Equal(t, dbConfig.port, configVars["DB_PORT"])
	assert.Equal(t, dbConfig.name, configVars["DB_NAME"])
}

func TestDBConnectionString(t *testing.T) {
	configVars := map[string]string{
		"DB_HOST": "127.0.0.1",
		"DB_PORT": "5432",
		"DB_USER": "reuni_test",
		"DB_PASS": "reuni_test123!",
		"DB_NAME": "reuni_test",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}
	assert.Equal(t, GetConnectionString(), "user=reuni_test password=reuni_test123! dbname=reuni_test host=127.0.0.1 port=5432 sslmode=disable")
}
