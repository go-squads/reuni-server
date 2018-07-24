package server

import (
	"github.com/go-squads/reuni-server/configurations"
	"github.com/go-squads/reuni-server/services"
	"github.com/go-squads/reuni-server/users"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/services", services.GetAllServicesHandler).Methods("GET")
	router.HandleFunc("/services", services.CreateServiceHandler).Methods("POST")
	router.HandleFunc("/services", services.DeleteServiceHandler).Methods("DELETE")
	router.HandleFunc("/services/{service_name}/namespaces", configurations.RetrieveAllNamespaceHandler).Methods("GET")
	router.HandleFunc("/services/{service_name}/namespaces", configurations.CreateNamespace).Methods("POST")
	router.HandleFunc("/users", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/loginuser", users.LoginUserHandler).Methods("POST")

	return router
}
