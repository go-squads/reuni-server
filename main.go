package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
)

func multiply(x,y int) int {
	return x*y
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(os.Getenv("POSTGRES_DB")))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}