package cmd

import (
	"fmt"
	"os"

	"github.com/go-squads/reuni-server/authenticator"
)

func GenerateRSAKey() {
	priv, pub := authenticator.GenerateRsaKeyPair()
	privText := authenticator.ExportRsaPrivateKeyAsPemStr(priv)
	pubText, err := authenticator.ExportRsaPublicKeyAsPemStr(pub)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Mkdir("./keys", os.ModePerm)
	privFile, fileErr := os.Create("./keys/private")
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	defer privFile.Close()

	pubFile, fileErr := os.Create("./key/public")
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	defer pubFile.Close()

	fmt.Fprintf(privFile, "%v\n", privText)
	fmt.Fprintf(pubFile, "%v\n", pubText)
	fmt.Println("Keys created at keys/")
}
