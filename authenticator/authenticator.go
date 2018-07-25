package authenticator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"hash"
	"log"
	"strings"

	"github.com/go-squads/reuni-server/appcontext"
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

func CreateUserJWToken(payload []byte) string {
	log.Println("Creating JWT with ", payload)
	header := createJWTHeader()
	tokenBase := helper.EncodeSegment(header) + "." + helper.EncodeSegment(payload)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, appcontext.GetKeys().PrivateKey, crypto.SHA256, hashJWT(tokenBase).Sum(nil))
	return tokenBase + "." + helper.EncodeSegment(signature)
}

func VerifyUserJWToken(token string) (interface{}, bool) {
	segments := strings.Split(token, ".")
	tokenBase := segments[0] + "." + segments[1]
	signature, err := helper.DecodeSegment(segments[2])
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}
	err = rsa.VerifyPKCS1v15(appcontext.GetKeys().PublicKey, crypto.SHA256, hashJWT(tokenBase).Sum(nil), signature)
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}
	payload, _ := helper.DecodeSegment(segments[1])
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}
	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		log.Println(err.Error())
	}
	return payloadMap, true
}
