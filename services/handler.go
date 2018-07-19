package services

import (
	"encoding/json"
	"log"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(servicesjson)
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
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("201 Created"))
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("200 OK"))
}

func createServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return createService(serviceStore)
}

func deleteServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return deleteService(serviceStore)
}
