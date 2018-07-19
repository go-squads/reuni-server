package services

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllServices(w http.ResponseWriter, r *http.Request) {

	servicesjson, err := json.Marshal(GetAll())
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/xml")
	w.Write(servicesjson)
}
