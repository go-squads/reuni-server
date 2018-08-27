package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/authenticator"
	"github.com/go-squads/reuni-server/helper"
	"github.com/go-squads/reuni-server/response"
)

var proc userProcessorInterface

func getProcessor() userProcessorInterface {
	if proc == nil {
		proc = &userProcessor{repo: initRepository(appcontext.GetHelper())}
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userdata userv
	err := json.NewDecoder(r.Body).Decode(&userdata)
	defer r.Body.Close()

	if err != nil {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	if !userdata.Valid() {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, "User data not valid"))
		return
	}
	userdata.Password = getProcessor().createUserEncryptPassword(userdata.Username, userdata.Password)
	if userdata.Password == "" {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, "password cant be encrypted"))
		return
	}

	err = getProcessor().createUserProcessor(userdata)
	if err != nil {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusConflict, err.Error()))
		return
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentText, "201 Created")
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var logindata userv
	err := json.NewDecoder(r.Body).Decode(&logindata)
	defer r.Body.Close()
	if err != nil {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	logindata.Password = getProcessor().createUserEncryptPassword(logindata.Username, logindata.Password)
	if logindata.Password == "" {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, helper.NewHttpError(http.StatusInternalServerError, "password cant be encrypted"))
		return
	}

	userData, userDataForRefreshToken, err := getProcessor().loginUserProcessor(logindata)
	if err != nil {
		response.ResponseError("CreateUser", getFromContext(r, "username"), w, err)
		return
	}
	log.Println("LoginUserHandler: ", string(userData), "succesfully login")
	token, err := authenticator.CreateUserJWToken(userData, appcontext.GetKeys().PrivateKey)
	if err != nil {
		log.Println(err.Error())
	}
	refreshToken, err := authenticator.CreateUserJWToken(userDataForRefreshToken, appcontext.GetKeys().PrivateKey)
	if err != nil {
		log.Println(err.Error())
	}
	response.ResponseHelper(w, http.StatusOK, response.ContentJson, fmt.Sprintf("{\"token\": \"%v\",\"refresh_token\": \"%v\"}", token, refreshToken))
}

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	usr, err := getProcessor().getAllUserProcessor()
	if err != nil {
		response.ResponseError("GetAllUserHandler", getFromContext(r, "username"), w, err)
		return
	}
	response.ResponseHelper(w, 200, response.ContentJson, usr)

}

func GetNewTokenHandler(w http.ResponseWriter, r *http.Request) {
	username := getFromContext(r, "username")
	userData, err := getProcessor().getUserDataProcessor(username)
	if err != nil {
		response.ResponseError("GetNewToken", getFromContext(r, "username"), w, err)
		return
	}
	newToken, err := authenticator.CreateUserJWToken(userData, appcontext.GetKeys().PrivateKey)
	if err != nil {
		log.Println(err.Error())
	}
	response.ResponseHelper(w, http.StatusCreated, response.ContentJson, fmt.Sprintf("{\"token\": \"%v\"}", newToken))
}
