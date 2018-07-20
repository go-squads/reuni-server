package response

import "net/http"

const (
	ContentJson = "application/json"
	ContentText = "text/plain"
)

func ResponseHelper(w http.ResponseWriter, http_status int, content_type string, body string) {
	w.WriteHeader(http_status)
	w.Header().Set("Content-Type", content_type)
	w.Write([]byte(body))
}
