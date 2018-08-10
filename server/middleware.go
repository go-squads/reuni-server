package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-squads/reuni-server/authorizator"
	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/response"
	"github.com/go-squads/reuni-server/services"
	"github.com/gorilla/mux"

	"github.com/go-squads/reuni-server/authenticator"
)

func withAuthenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimLeft(token, "Bearer")
		token = strings.TrimLeft(token, " ")
		obj, err := authenticator.VerifyUserJWToken(token, appcontext.GetKeys().PublicKey)
		if obj != nil {
			log.Println("User", obj, "access from", r.RemoteAddr)
			ctx := context.WithValue(r.Context(), "userId", obj["id"])
			ctx = context.WithValue(ctx, "username", obj["username"])

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Println("Access from", r.RemoteAddr, "not authorized")
			log.Println(err.Error())
			response.ResponseHelper(w, http.StatusUnauthorized, response.ContentText, "")
			return
		}

	}
}

func withAuthorizator(next http.HandlerFunc, permission rune) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := authorizator.New(appcontext.GetHelper())
		serviceName := mux.Vars(r)["service_name"]
		service, err := services.FindOneServiceByName(serviceName)
		if err != nil {
			response.ResponseError("AuthorizatorMiddleware", "", w, helper.NewHttpError(http.StatusNotFound, "Not Found"))
			return
		}
		uid, err := strconv.ParseInt(fmt.Sprintf("%v", r.Context().Value("userId")), 10, 64)
		if err != nil {
			response.ResponseError("AuthorizationMiddleware", "", w, helper.NewHttpError(http.StatusInternalServerError, "Cannot parse userId"))
		}
		res := auth.Authorize(int(uid), service.Id, permission)
		if res {
			next.ServeHTTP(w, r)
		} else {
			response.ResponseError("AuthorizationMiddleware", "", w, helper.NewHttpError(http.StatusForbidden, "Forbidden"))
		}
	}
}

func validateAgentTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		serviceName := mux.Vars(r)["service_name"]
		res, err := services.ValidateTokenProcessor(serviceName, token)
		if err != nil {
			response.ResponseHelper(w, http.StatusInternalServerError, response.ContentText, "")
			return
		}
		if res {
			next.ServeHTTP(w, r)
		} else {
			response.ResponseHelper(w, http.StatusForbidden, response.ContentText, "")
		}
	}
}
