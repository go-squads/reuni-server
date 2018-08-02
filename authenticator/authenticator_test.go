package authenticator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJWTHeader(t *testing.T) {
	json := "{\"alg\":\"RS256\",\"typ\":\"JWT\"}"
	fetchedJSON := createJWTHeader()
	assert.Equal(t, json, string(fetchedJSON))
}
