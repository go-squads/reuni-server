package authenticator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"hash"
	"log"
	"strings"

	"github.com/go-squads/reuni-server/helper"
)

func createJWTHeader() []byte {
	var header map[string]string
	header = make(map[string]string)
	header["alg"] = "RS256"
	header["typ"] = "JWT"
	headerJSON, _ := json.Marshal(header)
	return headerJSON
}

func hashJWT(token string) hash.Hash {
	hasher := sha256.New()
	hasher.Write([]byte(token))
	return hasher
}

func CreateUserJWToken(payload []byte, key *rsa.PrivateKey) (string, error) {
	log.Println("Creating JWT with ", payload)
	header := createJWTHeader()
	tokenBase := helper.EncodeSegment(header) + "." + helper.EncodeSegment(payload)
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashJWT(tokenBase).Sum(nil))
	if err != nil {
		return "", err
	}
	return tokenBase + "." + helper.EncodeSegment(signature), nil
}

func parseToken(token string) ([]string, error) {
	segments := strings.Split(token, ".")
	if len(segments) < 3 {
		return nil, errors.New("Failed to parse token")
	}
	return segments, nil
}

func VerifyUserJWToken(token string, key *rsa.PublicKey) (map[string]interface{}, error) {
	segments, err := parseToken(token)
	if err != nil {
		return nil, err
	}
	tokenBase := segments[0] + "." + segments[1]
	signature, err := helper.DecodeSegment(segments[2])
	if err != nil {
		return nil, err
	}
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashJWT(tokenBase).Sum(nil), signature)
	if err != nil {
		return nil, err
	}
	payload, err := helper.DecodeSegment(segments[1])
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return payloadMap, nil
}
