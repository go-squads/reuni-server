package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestDatabaseConfig(t *testing.T) {
	dbConfig := getDatabaseConfig()
	assert.Equal(t, dbConfig.host, os.Getenv("DB_HOST"))
	assert.Equal(t, dbConfig.password, os.Getenv("DB_PASS"))
	assert.Equal(t, dbConfig.username, os.Getenv("DB_USER"))
	assert.Equal(t, dbConfig.port, os.Getenv("DB_PORT"))
	assert.Equal(t, dbConfig.name, os.Getenv("DB_NAME"))

}
