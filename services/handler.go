package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
	"github.com/go-squads/reuni-server/response"
	"github.com/gorilla/mux"
)

var proc serviceProcessorInterface

func getProcessor() serviceProcessorInterface {
	if proc == nil {
		proc = &serviceProcessor{repo: initRepository(appcontext.GetHelper())}
	}
	return proc
}

func getFromContext(r *http.Request, key string) string {
	data := r.Context().Value(key)
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}
func toString(obj interface{}) string {
	js, _ := json.Marshal(obj)
	return string(js)
}

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	services, err := getProcessor().getAllServicesBasedOnOrganizationProcessor(organizationId)
	if err != nil {
		response.ResponseError("GetAllService", getFromContext(r, "username"), w, err)
		return
	}
	if len(services) == 0 {
		response.ResponseHelper(w, http.StatusOK, response.ContentJson, "[]")
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, toString(services))
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var servicedata servicev
	err := json.NewDecoder(r.Body).Decode(&servicedata)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	defer r.Body.Close()
	reg, _ := regexp.Compile(`^[^.|\s]+$`)
	if !reg.MatchString(servicedata.Name) {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Service name should not contain '.' or any whitespaces"))
		return
	}
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	servicedata.CreatedBy = getFromContext(r, "username")
	err = getProcessor().createServiceProcessor(servicedata, organizationId)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("DeleteServiceHandler", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	var servicedata servicev
	err = json.NewDecoder(r.Body).Decode(&servicedata)
	defer r.Body.Close()

	if err != nil {
		response.ResponseError("DeleteServiceHandler", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	err = getProcessor().deleteServiceProcessor(organizationId, servicedata)
	if err != nil {
		response.ResponseError("DeleteServiceHandler", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("ValidateToken", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	serviceName := mux.Vars(r)["service_name"]
	token := r.Header.Get("Authorization")
	result, err := getProcessor().ValidateTokenProcessor(organizationId, serviceName, token)
	if err != nil {
		response.ResponseError("ValidateToken", getFromContext(r, "username"), w, err)
		return
	}
	if result {
		response.ResponseHelper(w, http.StatusOK, response.ContentJson, `{"valid": true}`)
		return
	} else {
		response.ResponseHelper(w, http.StatusOK, response.ContentJson, `{"valid": false}`)
		return
	}
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("GetToken", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	serviceName := mux.Vars(r)["service_name"]
	log.Println("ORGID: " + fmt.Sprint(organizationId) + " SERVICENAME: " + serviceName)
	token, err := initRepository(appcontext.GetHelper()).getServiceToken(organizationId, serviceName)

	if err != nil {
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
	}
	log.Println(token)
	tokenJSON, _ := json.Marshal(token)
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(tokenJSON))
}
