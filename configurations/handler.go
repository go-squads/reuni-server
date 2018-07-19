package configurations

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateNamespace(w http.ResponseWriter, r *http.Request) {
	var configdata configurationView
	var service_name = mux.Vars(r)["service_name"]
	log.Printf("CreateNamespace: Get Request to %v", service_name)
	err := json.NewDecoder(r.Body).Decode(&configdata)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateNamespace: error parsing body")
		return
	}
	log.Println(configdata)

	err = createNewNamespaceProcess(service_name, configdata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateNamespace: error writing to database", err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("201 Created"))
}
