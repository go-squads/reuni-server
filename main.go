package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/cmd"
	"github.com/go-squads/reuni-server/server"
)

func main() {
	if len(os.Args) < 2 {
		appcontext.InitContext()
		router := server.CreateRouter()
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		switch os.Args[1] {
		case "keygen":
			{
				cmd.GenerateRSAKey()
			}
		}
	}
}
