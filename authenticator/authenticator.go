package authenticator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"hash"
	"log"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
)

func createJWTHeader() []byte {
	var header map[string]string
	header = make(map[string]string)
	header["alg"] = "RS512"
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
