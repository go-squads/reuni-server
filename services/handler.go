package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/response"
)

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, err := getAll()
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

	err = createServiceProcess(servicedata)
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

	err = deleteServiceProcess(servicedata)
	if err != nil {
		log.Println("DeleteServiceHandler: error writing to database", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentText, "200 OK")
}
