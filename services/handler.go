package services

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {

	servicesjson, err := json.Marshal(getAll())
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("success"))
}

func createServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return createService(serviceStore)
}
