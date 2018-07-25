package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-squads/reuni-server/response"

	"github.com/go-squads/reuni-server/authenticator"
)

func withAuthenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimLeft(token, "Bearer")
		token = strings.TrimLeft(token, " ")
		obj, res := authenticator.VerifyUserJWToken(token)
		if res {
			log.Println("User", obj, "access from", r.RemoteAddr)
			next.ServeHTTP(w, r)
		} else {
			log.Println("Access from", r.RemoteAddr, "not authorized")
			response.ResponseHelper(w, http.StatusUnauthorized, response.ContentText, "")
		}

	}
}
