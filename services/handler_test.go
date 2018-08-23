package services

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func ServeRequest(rr *httptest.ResponseRecorder, req *http.Request, handler http.HandlerFunc) {
	handler.ServeHTTP(rr, req)
}
func TestToString(t *testing.T) {
	data := make(map[string]interface{})
	data["id"] = 1
	res := toString(data)
	expected, _ := json.Marshal(data)
	assert.Equal(t, string(expected), res)
}

func TestGetAllHandlerShouldReturnOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	services := []service{}
	services = append(services, service{Name: "test"})

	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().getAllServicesBasedOnOrganizationProcessor(1).Return(services, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/services", strings.NewReader(""))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", GetAllServicesHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestGetAllHandlerShouldNotPanic(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().getAllServicesBasedOnOrganizationProcessor(1).Return([]service{}, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/services", strings.NewReader(""))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", GetAllServicesHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllHandlerShouldReturnErrorWhenBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().getAllServicesBasedOnOrganizationProcessor(1).Return([]service{}, helper.NewHttpError(400, "Bad Request"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/services", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", GetAllServicesHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllHandlerShouldReturnErrorWhenFailtoTranslateId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, helper.NewHttpError(http.StatusBadRequest, "Bad request"))
	mock.EXPECT().getAllServicesBasedOnOrganizationProcessor(1).Return([]service{}, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/services", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", GetAllServicesHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateServiceHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().createServiceProcessor(servicev{Name: "test"}, 1).Return(nil)

	payload := `
		{
			"name": "test"
		}
	`

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", CreateServiceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateServiceHandlerShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().createServiceProcessor(servicev{Name: "test"}, 1).Return(helper.NewHttpError(500, "Test"))

	payload := `
		{
			"name": "test"
		}
	`

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", CreateServiceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateServiceHandlerShouldReturnErrorWhenPayloadIsempty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, helper.NewHttpError(http.StatusBadRequest, "bad request"))
	mock.EXPECT().createServiceProcessor(servicev{Name: "test"}, 1).Return(nil)

	payload := `
		{
			"name": "test"
		}
	`

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", CreateServiceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateServiceHandlerShouldReturnErrorWhenPayloadMalformed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().createServiceProcessor(servicev{Name: "test"}, 1).Return(helper.NewHttpError(500, "Test"))

	payload := `
		{
			name: ""
			"asdsd
		}
	`
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", CreateServiceHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteServiceHandlerShouldNotReturnError(t *testing.T) {
	payload := `
		{
			"name": "test"
		}
	`

	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().deleteServiceProcessor(1, servicev{Name: "test"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", DeleteServiceHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteServiceHandlerShouldReturnErrorWhenTranslateNameReturnError(t *testing.T) {
	payload := `
		{
			"name": "test"
		}
	`

	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(0, helper.NewHttpError(http.StatusBadRequest, "bad request"))
	mock.EXPECT().deleteServiceProcessor(1, servicev{Name: "test"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", DeleteServiceHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteServiceHandlerShouldNotReturnErrorWhenServiceNotExist(t *testing.T) {
	payload := `
		{
			"name": "test"
		}
	`

	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().deleteServiceProcessor(1, servicev{Name: "test"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", DeleteServiceHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteServiceHandlerShouldReturnErrorWhenPayloadNotUnmarshalable(t *testing.T) {
	payload := `
		{
			Test:
			"name": "test-service",
		}
	`
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().deleteServiceProcessor(1, servicev{Name: "test"}).Return(nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", DeleteServiceHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteServiceHandlerShouldReturnErrorWhenProcessorError(t *testing.T) {
	payload := `
		{
			"name": "test"
		}
	`
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock
	service_expected := servicev{
		Name: "test",
	}
	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().deleteServiceProcessor(1, service_expected).Return(helper.NewHttpError(http.StatusInternalServerError, "Test Error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test/services", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/services", DeleteServiceHandler).Methods("DELETE")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestValidateTokenShouldReturnValidTrue(t *testing.T) {
	payload := `
		{
			"name": "test"
		}
	`
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("test").Return(1, nil)
	mock.EXPECT().ValidateTokenProcessor(1, "test", "HelloWorld!").Return(true, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/test/validatetoken", strings.NewReader(payload))
	req.Header.Set("Authorization", "HelloWorld!")

	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/validatetoken", ValidateToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]bool
	json.NewDecoder(rr.Body).Decode(&data)
	assert.True(t, data["valid"])
}

func TestValidateTokenShouldReturnValidFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(1, nil)
	mock.EXPECT().ValidateTokenProcessor(1, "test", "HelloWorld!").Return(false, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/validatetoken", strings.NewReader(""))
	req.Header.Set("Authorization", "HelloWorld!")

	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/validatetoken", ValidateToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]bool
	json.NewDecoder(rr.Body).Decode(&data)
	assert.False(t, data["valid"])
}

func TestValidateTokenShouldReturnErrorWhenTranslateNameReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(0, helper.NewHttpError(http.StatusBadRequest, "bad request"))
	mock.EXPECT().ValidateTokenProcessor(1, "test", "").Return(false, helper.NewHttpError(500, "TestError"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/validatetoken", strings.NewReader(""))

	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/validatetoken", ValidateToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestValidateTokenShouldReturnErrorWhenProcessorError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(1, nil)
	mock.EXPECT().ValidateTokenProcessor(1, "test", "").Return(false, helper.NewHttpError(500, "TestError"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/validatetoken", strings.NewReader(""))

	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/validatetoken", ValidateToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetTokenHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(1, nil)

	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "testToken"}},
		Err:  nil,
	}
	appcontext.InitMockContext(q)

	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test-services/token", nil)
	// ServeRequest(rr, req, GetToken)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/token", GetToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]string
	json.NewDecoder(rr.Body).Decode(&data)
	assert.Equal(t, "testToken", data["authorization_token"])
}

func TestGetTokenHandlerShouldReturnErrorWhenGetTokenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(1, nil)

	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "testToken"}},
		Err:  helper.NewHttpError(500, "Bad Error"),
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test-services/token", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/token", GetToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var data map[string]string
	json.NewDecoder(rr.Body).Decode(&data)
	assert.Equal(t, "", data["authorization_token"])
}

func TestGetTokenHandlerShouldReturnErrorWhenTranslateNameReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceProcessorInterface(ctrl)
	proc = mock

	mock.EXPECT().TranslateNameToIdProcessor("org").Return(0, helper.NewHttpError(http.StatusBadRequest, "bad request"))

	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "testToken"}},
		Err:  helper.NewHttpError(500, "Bad Error"),
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test-services/token", nil)
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/token", GetToken).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var data map[string]string
	json.NewDecoder(rr.Body).Decode(&data)
	assert.Equal(t, "", data["authorization_token"])
}

func TestGetFromContextShouldReturnTheRightValue(t *testing.T) {
	r, _ := http.NewRequest("GET", "/org/test-services/token", nil)
	ctx := context.WithValue(r.Context(), "username", "go-squads")
	r = r.WithContext(ctx)
	assert.Equal(t, "go-squads", getFromContext(r, "username"))
}

func TestGetUsernameShouldReturnEmptyString(t *testing.T) {
	r, _ := http.NewRequest("GET", "/org/test-services/token", nil)
	assert.Empty(t, getFromContext(r, "username"))
}
