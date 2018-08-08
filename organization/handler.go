package organization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/response"
)

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
	uid, err := strconv.ParseInt(getFromContext(r, "userId"), 10, 64)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	err = createNewOrganizationProcessor(data.Name, uid)
	if err != nil {
		response.ResponseError("CreateOrganization", getFromContext(r, "username"), w, err)
		return
	}
	defer r.Body.Close()

}
