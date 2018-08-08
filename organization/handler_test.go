package organization

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func ServeRequest(rr *httptest.ResponseRecorder, req *http.Request, handler http.HandlerFunc) {
	handler.ServeHTTP(rr, req)
}
func TestCreateOrganizationHandlerShouldReturn400WhenDataIsNotParsable(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			asdfasdfs
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req, CreateOrganizationHandler)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateOrganizationHandlerShouldReturn500WhenUserIdNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"name": "test"
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req, CreateOrganizationHandler)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateOrganizationHandlerShouldReturn400WhenOrganizationNameEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"name": ""
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req, CreateOrganizationHandler)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateOrganizationHandlerShouldReturn500WhenUserIdNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"name": "test"
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req.WithContext(context.WithValue(req.Context(), "userId", "something")), CreateOrganizationHandler)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateOrganizationHandlerShouldReturn500WhenProcessorError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"name": "test"
		}
	`
	mock.EXPECT().createNewOrganizationProcessor("test", int64(1)).Return(errors.New("Internal Error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req.WithContext(context.WithValue(req.Context(), "userId", 1)), CreateOrganizationHandler)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateOrganizationHandlerShouldReturn201(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"name": "test"
		}
	`
	mock.EXPECT().createNewOrganizationProcessor("test", int64(1)).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization", strings.NewReader(payload))
	ServeRequest(rr, req.WithContext(context.WithValue(req.Context(), "userId", 1)), CreateOrganizationHandler)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetProcessorShouldInstantiateMainRepository(t *testing.T) {
	proc = nil
	activeProc := getProcessor()
	_, ok := activeProc.(*mainProcessor)
	assert.True(t, ok)
}

func TestAddUserShouldReturn500WhenOrgIdCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"user_id": 1,
			"role": "Developer"
		}
	`
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "Developer",
	}
	mock.EXPECT().addUserProcessor(member).Return(errors.New("Internal error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization/error/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", AddUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestAddUserShouldReturn500WhenBodyCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			adasdas
		}
	`
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "Developer",
	}
	mock.EXPECT().addUserProcessor(member).Return(errors.New("Internal error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", AddUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestAddUserShouldReturnErrorWhenDataMemberIsNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"user_id": 1,
			"role": "ooa"
		}
	`
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "ooa",
	}
	mock.EXPECT().addUserProcessor(member).Return(errors.New("Internal error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", AddUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
func TestAddUserShouldReturn201WhenAddSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"user_id": 1,
			"role": "Developer"
		}
	`
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "Developer",
	}
	mock.EXPECT().addUserProcessor(member).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", AddUserHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteUserShouldReturn500WhenOrgIdCantBeParsed(t *testing.T) {
	payload := `
		{
			"user_id": 1
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organization/error/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", DeleteUserFromGroupHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteUserShouldReturn500WhenBodyCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			adasdas
		}
	`
	mock.EXPECT().deleteUserFromGroupProcessor(int64(1), int64(1)).Return(errors.New("Internal error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", DeleteUserFromGroupHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteUserShouldReturnErrorWhenDataIsNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().deleteUserFromGroupProcessor(int64(1), int64(1)).Return(errors.New("Internal error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", DeleteUserFromGroupHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
func TestDeleteUserShouldReturn200WhenDeleteSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockprocessor(ctrl)
	proc = mock
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().deleteUserFromGroupProcessor(int64(1), int64(1)).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organization/1/member", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/organization/{org_id}/member", DeleteUserFromGroupHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
