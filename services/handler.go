package services

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := getProcessor().TranslateNameToIdProcessor(organizationName)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = getProcessor().createServiceProcessor(servicedata, organizationId)
	if err != nil {
		response.ResponseError("CreateService", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	var servicedata servicev
	err := json.NewDecoder(r.Body).Decode(&servicedata)
	defer r.Body.Close()

	if err != nil {
		response.ResponseError("DeleteServiceHandler", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	err = getProcessor().deleteServiceProcessor(servicedata)
	if err != nil {
		response.ResponseError("DeleteServiceHandler", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service_name"]
	token := r.Header.Get("Authorization")
	result, err := getProcessor().ValidateTokenProcessor(serviceName, token)
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
	serviceName := mux.Vars(r)["service_name"]
	token, err := initRepository(appcontext.GetHelper()).getServiceToken(serviceName)
	if err != nil {
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
	}
	tokenJSON, _ := json.Marshal(token)
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(tokenJSON))
}
