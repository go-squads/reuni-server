package organization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/response"
)

var proc processor

func getProcessor() processor {
	if proc == nil {
		proc = &mainProcessor{}
	}
	return proc
}

func getFromContext(r *http.Request, key string) string {
	data := r.Context().Value(key)
	if data == nil {
		return ""
	}
	return fmt.Sprintf("%v", data)
}

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var data Organization
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	if data.Name == "" {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "Name cannot be empty"))
		return
	}
	uid, err := strconv.ParseInt(getFromContext(r, "userId"), 10, 64)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	err = getProcessor().createNewOrganizationProcessor(data.Name, uid)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, err)
		return
	}
	defer r.Body.Close()
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}
