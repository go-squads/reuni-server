package services

import (
	"encoding/json"
	"log"
	"net/http"

	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/response"
	"github.com/gorilla/mux"
)

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, err := getAll(context.GetHelper())
	if err != nil {
		log.Println(err.Error())
		return
	}
	servicesjson, err := json.Marshal(services)
	if err != nil {
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(servicesjson))
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var servicedata servicev
	err := json.NewDecoder(r.Body).Decode(&servicedata)
	defer r.Body.Close()

	if err != nil {
		log.Println("CreateServiceHandler: error parsing body")
		return
	}

	err = createServiceProcessor(servicedata)
	if err != nil {
		log.Println("CreateServiceHandler: error writing to database", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	var servicedata servicev
	err := json.NewDecoder(r.Body).Decode(&servicedata)
	defer r.Body.Close()

	if err != nil {
		log.Println("DeleteerviceHandler: error parsing body")
		return
	}

	err = deleteServiceProcessor(servicedata)
	if err != nil {
		log.Println("DeleteServiceHandler: error writing to database", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	serviceName := mux.Vars(r)["service_name"]
	token := r.Header.Get("Authorization")
	result, err := ValidateTokenProcessor(serviceName, token)
	if err != nil {
		log.Println("ValidateToken: ", err.Error())
		response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
		return
	}
	if result {
		response.ResponseHelper(w, http.StatusOK, response.ContentText, "true")
		return
	} else {
		response.ResponseHelper(w, http.StatusOK, response.ContentText, "false")
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
