package server

import (
	"github.com/gorilla/mux"
	"github.com/go-squads/reuni-server/handler"
)


func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.HomeHandler)
	return router
}