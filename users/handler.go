package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/authenticator"
	"github.com/go-squads/reuni-server/response"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userdata userv
	err := json.NewDecoder(r.Body).Decode(&userdata)
	defer r.Body.Close()

	if err != nil {
		log.Println("CreateUserHandler: " + err.Error())
		response.RespondWithError(w, http.StatusBadRequest, response.ContentJson, "error parsing body")
		return
	}

	err = createUserProcessor(userdata)
	if err != nil {
		log.Println("CreateUserHandler: " + err.Error())
		response.RespondWithError(w, http.StatusConflict, response.ContentJson, "error writing to database")
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var logindata userv
	err := json.NewDecoder(r.Body).Decode(&logindata)
	defer r.Body.Close()

	logindata.Password = createUserEncryptPassword(logindata.Username, logindata.Password)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, response.ContentJson, "error parsing body")
		return
	}

	userData, err := loginUser(logindata)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			response.RespondWithError(w, http.StatusUnauthorized, response.ContentJson, "wrong username or password")
		default:
			response.RespondWithError(w, http.StatusInternalServerError, response.ContentText, "")
		}
		log.Println("LoginUserHandler: ", err.Error())
		return
	}
	log.Println("LoginUserHandler: ", string(userData), "succesfully login")
	token, err := authenticator.CreateUserJWToken(userData, appcontext.GetKeys().PrivateKey)
	if err != nil {
		log.Println(err.Error())
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, fmt.Sprintf("{\"token\": \"%v\"}", token))
}
