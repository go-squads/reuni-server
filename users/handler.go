package users

import (
	"encoding/json"
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
