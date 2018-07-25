package config

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-squads/reuni-server/helper"
)

type Keys struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetKeys() (*Keys, error) {
	pubKeyByte, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))
	check(err)
	privKeyByte, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))
	check(err)
	log.Println(string(pubKeyByte))
	log.Println(string(privKeyByte))
	privKey, err := helper.ParseRsaPrivateKeyFromPemStr(string(privKeyByte))
	check(err)
	pubKey, err := helper.ParseRsaPublicKeyFromPemStr(string(pubKeyByte))
	return &Keys{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}, nil
}
