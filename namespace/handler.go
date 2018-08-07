package namespace

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/response"
	"github.com/gorilla/mux"
)

func CreateNamespace(w http.ResponseWriter, r *http.Request) {
	var namespaceData namespaceView
	var serviceName = mux.Vars(r)["service_name"]
	log.Printf("CreateNamespace: Get Request to %v", serviceName)
	err := json.NewDecoder(r.Body).Decode(&namespaceData)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateNamespace: error parsing body")
		return
	}
	log.Println(namespaceData)

	err = createNewNamespaceProcessor(serviceName, namespaceData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("CreateNamespace: error writing to database", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func RetrieveAllNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	var serviceName = mux.Vars(r)["service_name"]
	log.Printf("RetrieveAllNamespaces: Get Request to %v for retrieve all data", serviceName)
	configsjson, err := retrieveAllNamespaceProcess(serviceName)
	if err != nil {
		log.Println("RetrieveAllNamespaces: ", err.Error())
		return
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(configsjson))

}
