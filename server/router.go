package server

import (
	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/configuration"
	"github.com/go-squads/reuni-server/namespace"
	"github.com/go-squads/reuni-server/organization"
	"github.com/go-squads/reuni-server/services"
	"github.com/go-squads/reuni-server/users"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	configuration := configuration.New(appcontext.GetHelper())
	router := mux.NewRouter()
	router.HandleFunc("/{organization_name}/services", withAuthenticator(services.GetAllServicesHandler)).Methods("GET")
	router.HandleFunc("/{organization_name}/services", withAuthenticator(services.CreateServiceHandler)).Methods("POST")
	router.HandleFunc("/{organization_name}/services", withAuthenticator(services.DeleteServiceHandler)).Methods("DELETE")
	router.HandleFunc("/services/{service_name}/namespaces", withAuthenticator(withAuthorizator(namespace.RetrieveAllNamespaceHandler, 'r'))).Methods("GET")
	router.HandleFunc("/services/{service_name}/namespaces", withAuthenticator(withAuthorizator(namespace.CreateNamespaceHandler, 'w'))).Methods("POST")
	router.HandleFunc("/services/{service_name}/validatetoken", services.ValidateToken).Methods("GET")
	router.HandleFunc("/services/{service_name}/token", withAuthenticator(withAuthorizator(services.GetToken, 'r'))).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/versions", withAuthenticator(withAuthorizator(configuration.GetConfigurationVersionsHandler, 'r'))).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/latest", withAuthenticator(withAuthorizator(configuration.GetLatestVersionHandler, 'r'))).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/agent", validateAgentTokenMiddleware(configuration.GetLatestVersionHandler)).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}/{version}/agent", validateAgentTokenMiddleware(configuration.GetConfigurationHandler))
	router.HandleFunc("/services/{service_name}/{namespace}/{version}", withAuthenticator(withAuthorizator(configuration.GetConfigurationHandler, 'r'))).Methods("GET")
	router.HandleFunc("/services/{service_name}/{namespace}", withAuthenticator(withAuthorizator(configuration.CreateNewVersionHandler, 'w'))).Methods("POST")
	router.HandleFunc("/organization", withAuthenticator(organization.GetAllHandler)).Methods("GET")
	router.HandleFunc("/organization", withAuthenticator(organization.CreateOrganizationHandler)).Methods("POST")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.AddUserHandler)).Methods("POST")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.DeleteUserFromGroupHandler)).Methods("DELETE")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.UpdateRoleOfUserHandler)).Methods("PATCH")
	router.HandleFunc("/organization/{org_id}/member", withAuthenticator(organization.GetAllMemberOfOrganizationHandler)).Methods("GET")
	router.HandleFunc("/signup", users.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", users.LoginUserHandler).Methods("POST")
	return router
}
