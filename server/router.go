package server

import (
	"github.com/go-squads/reuni-server/services"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/services", services.GetAllServices).Methods("GET")
	return router
}
