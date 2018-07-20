package configurations

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/response"
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
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func RetrieveAllNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	var service_name = mux.Vars(r)["service_name"]
	log.Printf("RetrieveAllNamespaces: Get Request to %v for retrieve all data", service_name)
	configsjson, err := retrieveAllNamespaceProcess(service_name)
	if err != nil {
		log.Println("RetrieveAllNamespaces: ", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(configsjson))

}
