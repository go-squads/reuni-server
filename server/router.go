package server

import (
	"github.com/go-squads/reuni-server/configuration"
	"github.com/go-squads/reuni-server/namespace"
	"github.com/go-squads/reuni-server/organization"
	"github.com/go-squads/reuni-server/services"
	"github.com/go-squads/reuni-server/users"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/services", withAuthenticator(services.GetAllServicesHandler)).Methods("GET")
	router.HandleFunc("/services", withAuthenticator(services.CreateServiceHandler)).Methods("POST")
	router.HandleFunc("/services", withAuthenticator(services.DeleteServiceHandler)).Methods("DELETE")
	router.HandleFunc("/services/{service_name}/namespaces", withAuthenticator(namespace.RetrieveAllNamespaceHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/namespaces", withAuthenticator(namespace.CreateNamespaceHandler)).Methods("POST")
	router.HandleFunc("/services/{service_name}/validatetoken", services.ValidateToken).Methods("GET")
	router.HandleFunc("/services/{service_name}/token", withAuthenticator(services.GetToken)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/versions", withAuthenticator(configuration.GetConfigurationVersionsHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/latest", withAuthenticator(configuration.GetLatestVersionHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/agent", validateAgentTokenMiddleware(configuration.GetLatestVersionHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/{version}/agent", validateAgentTokenMiddleware(configuration.GetConfigurationHandler))
	router.HandleFunc("/services/{service_name}/{namespace}/{version}", withAuthenticator(configuration.GetConfigurationHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}", withAuthenticator(configuration.CreateNewVersionHandler)).Methods("POST")
	router.HandleFunc("/organization", withAuthenticator(organization.CreateOrganizationHandler)).Methods("POST")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.AddUserHandler)).Methods("POST")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.DeleteUserFromGroupHandler)).Methods("DELETE")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.UpdateRoleOfUserHandler)).Methods("PATCH")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.GetAllMemberOfOrganizationHandler)).Methods("GET")
	router.HandleFunc("/signup", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", users.LoginUserHandler).Methods("POST")
	return router
}
