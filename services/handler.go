package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/helper"

	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/response"
	"github.com/gorilla/mux"
)

func getUsername(r *http.Request) string {
	usr := r.Context().Value("username")
	if usr != nil {
		return usr.(string)
	}
	return ""

}

func toString(obj interface{}) string {
	js, _ := json.Marshal(obj)
	return string(js)
}

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := translateNameToIdRepository(appcontext.GetHelper(), organizationName)
	if err != nil {
		response.ResponseError("GetAllService", getUsername(r),w,err)
		return
	}
	services, err := getAll(context.GetHelper(), organizationId)
	if err != nil {
		response.ResponseError("GetAllService", getUsername(r), w, err)
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
		response.ResponseError("CreateService", getUsername(r), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	defer r.Body.Close()
	organizationName := mux.Vars(r)["organization_name"]
	organizationId, err := translateNameToIdRepository(appcontext.GetHelper(), organizationName)
	if err != nil {
		response.ResponseError("CreateService", getUsername(r), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = createServiceProcessor(servicedata, organizationId)
	if err != nil {
		response.ResponseError("CreateService", getUsername(r), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	var servicedata servicev
	err := json.NewDecoder(r.Body).Decode(&servicedata)
	defer r.Body.Close()

	if err != nil {
		response.ResponseError("DeleteServiceHandler", getUsername(r), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	err = deleteServiceProcessor(servicedata)
	if err != nil {
		response.ResponseError("DeleteServiceHandler", getUsername(r), w, err)
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service_name"]
	token := r.Header.Get("Authorization")
	result, err := ValidateTokenProcessor(serviceName, token)
	if err != nil {
		response.ResponseError("ValidateToken", getUsername(r), w, err)
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
	token, err := getServiceToken(context.GetHelper(), serviceName)
	if err != nil {
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
	}
	tokenJSON, _ := json.Marshal(token)
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(tokenJSON))
}
