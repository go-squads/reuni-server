package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-squads/reuni-server/response"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userdata userv
	err := json.NewDecoder(r.Body).Decode(&userdata)
	defer r.Body.Close()

	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, response.ContentJson, "CreateUserHandler: error parsing body"+err.Error())
		return
	}

	err = createUserProcessor(userdata)
	if err != nil {
		response.RespondWithError(w, http.StatusNotFound, response.ContentJson, "CreateUserHandler: error writing to database"+err.Error())
		return
	}

	response.ResponseHelper(w, http.StatusCreated, response.ContentJson, "201 Created")
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var logindata userv
	err := json.NewDecoder(r.Body).Decode(&logindata)
	defer r.Body.Close()

	logindata.Password = createUserEncryptPassword(logindata.Username, logindata.Password)
	fmt.Println(logindata.Username, logindata.Password)
	if err != nil {
		response.RespondWithError(w, http.StatusNotFound, response.ContentJson, "LoginUserHandler: error parsing body")
		return
	}

	userData, err := loginUser(logindata)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			response.RespondWithError(w, http.StatusNotFound, response.ContentJson, "Username/Password wrong")
		default:
			response.RespondWithError(w, http.StatusInternalServerError, response.ContentJson, string(err.Error()))
		}
		return
	}

	response.ResponseHelper(w, http.StatusOK, response.ContentJson, string(userData))
}
