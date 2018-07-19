package main

import (
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/server"
)

func main() {
	appcontext.InitContext()
	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
