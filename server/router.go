package server

import (
	"github.com/go-squads/reuni-server/configuration"
	"github.com/go-squads/reuni-server/namespace"
	"github.com/go-squads/reuni-server/services"
	"github.com/go-squads/reuni-server/users"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/services", withAuthenticator(services.GetAllServicesHandler)).Methods("GET")
	router.HandleFunc("/services", services.CreateServiceHandler).Methods("POST")
	router.HandleFunc("/services", services.DeleteServiceHandler).Methods("DELETE")
	router.HandleFunc("/services/{service_name}/namespaces", withAuthenticator(namespace.RetrieveAllNamespaceHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/namespaces", namespace.CreateNamespace).Methods("POST")
	router.HandleFunc("/services/{service_name}/validatetoken", services.ValidateToken).Methods("GET")
	router.HandleFunc("/services/{service_name}/token", services.GetToken).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/latest", configuration.GetLatestVersionHandler).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/agent", validateAgentTokenMiddleware(configuration.GetLatestVersionHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/{version}/agent", validateAgentTokenMiddleware(configuration.GetConfigurationHandler))
	router.HandleFunc("/services/{service_name}/{namespace}/{version}", configuration.GetConfigurationHandler).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}", configuration.CreateNewVersionHandler).Methods("POST")
	router.HandleFunc("/signup", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", users.LoginUserHandler).Methods("POST")
	return router
}
