package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	ContentJson = "application/json"
	ContentText = "text/plain"
)

func ResponseHelper(w http.ResponseWriter, http_status int, content_type string, body string) {
	w.Header().Set("Content-Type", content_type)
	w.WriteHeader(http_status)
	w.Write([]byte(body))
}

func RespondWithError(w http.ResponseWriter, code int, content_type string, message string) {
	errorMessage, _ := json.Marshal(map[string]string{"error": message})
	ResponseHelper(w, code, content_type, string(errorMessage))
}

func ResponseError(caller, user string, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	log.Error(err.Error())
	if httpErr, ok := err.(*helper.HttpError); ok && httpErr.Status != http.StatusInternalServerError {
		msg, _ := json.Marshal(httpErr)
		http.Error(w, string(msg), httpErr.Status)
	} else {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code)
			log.Println(pqErr.Message)
		}
		http.Error(w, `{"status": 500, "message": "Internal Server Error"}`, http.StatusInternalServerError)
	}

}
