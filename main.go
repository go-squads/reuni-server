package main

import (
	"net/http"
	"log"
	"github.com/go-squads/reuni-server/server"
)


func multiply(x,y int) int {
	return x*y
}


func main() {
	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}