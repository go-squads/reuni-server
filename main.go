package main

import (
	"net/http"
	"log"
	"github.com/go-squads/reuni-server/server"
	"github.com/go-squads/reuni-server/appcontext"
)

func main() {
	appcontext.InitContext()
	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}