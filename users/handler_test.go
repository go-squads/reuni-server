package users

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-squads/reuni-server/helper"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandlerShouldReturnErrorWhenBodyCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username":,
		}
	`
	mock.EXPECT().createUserProcessor(userv{Username: "test", Password: "test"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/signup", CreateUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateUserHandlerShouldReturnErrorWhenPasswordEncryptFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username":"test",
			"password":"test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("")
	mock.EXPECT().createUserProcessor(userv{Username: "test", Password: "test"}).Return(helper.NewHttpError(http.StatusConflict, "error writing to database"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/signup", CreateUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateUserHandlerShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username":"test",
			"password":"test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("testtest")
	mock.EXPECT().createUserProcessor(userv{Username: "test", Password: "testtest"}).Return(helper.NewHttpError(http.StatusConflict, "error writing to database"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/signup", CreateUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestCreateUserHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username":"test",
			"password":"test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("testtest")
	mock.EXPECT().createUserProcessor(userv{Username: "test", Password: "testtest"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/signup", CreateUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestLoginUserHandlerShouldReturnErrorWhenDataCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username"
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestLoginUserHandlerShouldReturnErrorWhenPasswordCantBeEncrypted(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username": "test",
			"password": "test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("")
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestLoginUserHandlerShouldReturnErrorWhenQueryLoginErrorSqlNoRows(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username": "test",
			"password": "test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("testtest")
	mock.EXPECT().loginUserProcessor(userv{Username: "test", Password: "testtest"}).Return(nil, sql.ErrNoRows)

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestLoginUserHandlerShouldReturnErrorWhenQueryLoginErrorDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserProcessorInterface(ctrl)
	proc = mock
	payload := `
		{
			"username": "test",
			"password": "test"
		}
	`
	mock.EXPECT().createUserEncryptPassword("test", "test").Return("testtest")
	mock.EXPECT().loginUserProcessor(userv{Username: "test", Password: "testtest"}).Return(nil, errors.New("internal error"))

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
