package response

import (
	"encoding/json"
	"net/http"
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
