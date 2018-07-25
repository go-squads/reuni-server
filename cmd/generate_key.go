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

	fmt.Println(privText)
	fmt.Println()
	fmt.Println(pubText)
	os.Exit(0)

}
