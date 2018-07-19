package server

import (
	"github.com/go-squads/reuni-server/configurations"
	"github.com/go-squads/reuni-server/services"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/services", services.GetAllServicesHandler).Methods("GET")
	router.HandleFunc("/services", services.CreateServiceHandler).Methods("POST")
	router.HandleFunc("/services", services.DeleteServiceHandler).Methods("DELETE")
	router.HandleFunc("/services/{service_name}/configurations", configurations.CreateNamespace).Methods("POST")
	return router
}
