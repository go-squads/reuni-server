package authenticator

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-squads/reuni-server/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetJWTHeader(t *testing.T) {
	json := "{\"alg\":\"RS256\",\"typ\":\"JWT\"}"
	fetchedJSON := createJWTHeader()
	assert.Equal(t, json, string(fetchedJSON))
}

func TestGetHasher(t *testing.T) {
	json := string(createJWTHeader()) + "." + `{"username":"kenneth"}`
	hasher := hashJWT(json)
	assert.NotNil(t, hasher)
}

func TestCreateUserJWToken(t *testing.T) {
	json := `{"testing":"123"}`
	priv, _ := helper.GenerateRsaKeyPair()
	token, err := CreateUserJWToken([]byte(json), priv)
	assert.NotNil(t, token)
	assert.NoError(t, err)
}

func TestParseToken(t *testing.T) {
	segments, err := parseToken("aklsdfjlsaf.alskdfjklsjaf.fasdflkasfd")
	assert.NotNil(t, segments)
	assert.NoError(t, err)
}

func TestParseTokenShouldReturnError(t *testing.T) {
	segments, err := parseToken("")
	assert.Nil(t, segments)
	assert.EqualError(t, err, "Failed to parse token")
}

func TestVerifyUserJWTokenVerified(t *testing.T) {
	expires := fmt.Sprint(time.Now().Add(time.Minute * 100).Unix())
	json := `{
		"test":"test",
		"exp": ` + expires + `}`
	priv, pub := helper.GenerateRsaKeyPair()
	token, _ := CreateUserJWToken([]byte(json), priv)
	obj, err := VerifyUserJWToken(token, pub)
	assert.Equal(t, "test", obj["test"])
	assert.NoError(t, err)
}

func TestVerifyUserJWTokenShouldReturnErrorWhenExpiryTimeHasPassed(t *testing.T) {
	expires := fmt.Sprint(0)
	json := `{
		"test":"test",
		"exp": ` + expires + `}`
	priv, pub := helper.GenerateRsaKeyPair()
	token, _ := CreateUserJWToken([]byte(json), priv)
	obj, err := VerifyUserJWToken(token, pub)
	assert.Nil(t, obj)
	assert.Error(t, err)
}

func TestVerifyUserJWTokenShouldReturnErrorWhenTokenChangeByUser(t *testing.T) {
	json := `{"test":"test"}`
	priv, pub := helper.GenerateRsaKeyPair()
	token, _ := CreateUserJWToken([]byte(json), priv)
	obj, err := VerifyUserJWToken(token+"h", pub)
	assert.Nil(t, obj)
	assert.Error(t, err, "crypto/rsa: verification error")
}

func TestVerifyUserJWTokenShouldReturnErrorWhenTokenNotJWTformat(t *testing.T) {
	_, pub := helper.GenerateRsaKeyPair()
	token := "Hello World!"
	obj, err := VerifyUserJWToken(token, pub)
	assert.Nil(t, obj)
	assert.EqualError(t, err, "Failed to parse token")
}
