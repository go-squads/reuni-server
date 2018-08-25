package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/response"

	"github.com/gorilla/mux"
)

type Configuration interface {
	GetConfigurationHandler(w http.ResponseWriter, r *http.Request)
	GetLatestVersionHandler(w http.ResponseWriter, r *http.Request)
	CreateNewVersionHandler(w http.ResponseWriter, r *http.Request)
	GetConfigurationVersionsHandler(w http.ResponseWriter, r *http.Request)
}

type mainConfiguration struct {
	processor Processor
}

func New(init interface{}) Configuration {
	switch v := init.(type) {
	case helper.QueryExecuter:
		return &mainConfiguration{processor: &mainProcessor{repo: &mainRepository{v}}}
	case Repository:
		return &mainConfiguration{processor: &mainProcessor{repo: v}}
	case Processor:
		return &mainConfiguration{processor: v}
	case Configuration:
		return v
	default:
		return nil
	}

}

func getFromContext(r *http.Request, key string) string {
	data := r.Context().Value(key)
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}

func (s *mainConfiguration) GetConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	organizationName := routerVar["organization_name"]
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	version, err := strconv.Atoi(routerVar["version"])
	if err != nil {
		log.Println("GetConfig: Cannot parse version to Integer.")
		response.RespondWithError(w, http.StatusBadRequest, response.ContentJson, "Cannot Parse Version")
		return
	}
	config, err := s.processor.getConfigurationProcess(organizationName, serviceName, namespace, version)
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

func (s *mainConfiguration) GetLatestVersionHandler(w http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	organizationName := routerVar["organization_name"]
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	version, err := s.processor.getLatestVersionProcess(organizationName, serviceName, namespace)
	versionv := versionView{Version: version}
	if err != nil {
		log.Println("GetConfig: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	versionJSON, err := json.Marshal(versionv)
	if err != nil {
		log.Println("GetConfig: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(versionJSON))
}

func (s *mainConfiguration) CreateNewVersionHandler(w http.ResponseWriter, r *http.Request) {
	var config configView
	routerVar := mux.Vars(r)
	organizationName := routerVar["organization_name"]
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Println("CreateNewConfigVersion: ", err.Error())
		response.ResponseHelper(w, http.StatusBadRequest, response.ContentText, "")
		return
	}
	if len(config.Configuration) == 0 {
		log.Println("CreateNewConfigVersion: ", "CONFIG EMPTY")
		response.ResponseHelper(w, http.StatusBadRequest, response.ContentText, "Configuration cant be empty")
		return
	}

	var newVersion int
	config.Created_by = getFromContext(r, "username")
	newVersion, err = s.processor.createNewVersionProcess(organizationName, serviceName, namespace, config)
	if err != nil {
		log.Println("CreateNewConfigVersion: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	err = appcontext.GetRedisHelper().Publish(organizationName+"_"+serviceName+"_"+namespace, newVersion)
	if err != nil {
		log.Println("CreateNewConfigVersion: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "")
}

func (s *mainConfiguration) GetConfigurationVersionsHandler(w http.ResponseWriter, r *http.Request) {
	routerVar := mux.Vars(r)
	organizationName := routerVar["organization_name"]
	serviceName := routerVar["service_name"]
	namespace := routerVar["namespace"]
	resp, err := s.processor.getConfigurationVersionsProcess(organizationName, serviceName, namespace)
	if err != nil {
		log.Println("GetConfigurationHandler:", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, resp)
}
