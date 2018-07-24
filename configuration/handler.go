package configuration

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-squads/reuni-server/response"

	"github.com/gorilla/mux"
)

func GetConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	version, err := strconv.Atoi(routerVar["version"])
	if err != nil {
		log.Println("GetConfig: Cannot parse version to Integer.")
		response.RespondWithError(w, http.StatusBadRequest, response.ContentJson, "Cannot Parse Version")
		return
	}
	config, err := getConfigurationProcess(serviceName, namespace, version)
	if err != nil {
		log.Println("GetConfig:", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	configJSON, err := json.Marshal(config)
	if err != nil {
		log.Println("GetConfig:", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(configJSON))
}

func GetLatestVersionHandler(w http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	version, err := getLatestVersionProcess(serviceName, namespace)
	if err != nil {
		log.Println("GetConfig: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	versionJSON, err := json.Marshal(version)
	if err != nil {
		log.Println("GetConfig: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(versionJSON))
}
