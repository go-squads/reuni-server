package configuration

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func ServeRequest(rr *httptest.ResponseRecorder, req *http.Request, handler http.HandlerFunc) {
	handler.ServeHTTP(rr, req)
}

func TestGetConfigurationHandlerShouldReturnErrorWhenVersionCantBeParsed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().getConfigurationProcess("org", "test", "test", 1).Return(nil, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/error", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/{version}", handler.GetConfigurationHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetConfigurationHandlerShouldReturnErrorWhenGetVersionReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().getConfigurationProcess("org", "test", "test", 1).Return(nil, errors.New("error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/1", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/{version}", handler.GetConfigurationHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetConfigurationHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().getConfigurationProcess("org", "test", "test", 1).Return(&configView{Version: 1}, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/1", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/{version}", handler.GetConfigurationHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetLatestVersionHandlerShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().getLatestVersionProcess("org", "test", "test").Return(1, errors.New("error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/latest", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/latest", handler.GetLatestVersionHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetLatestVersionHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id": 1
		}
	`
	mock.EXPECT().getLatestVersionProcess("org", "test", "test").Return(1, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/1", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/{version}", handler.GetLatestVersionHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateNewVersionHandlerShouldReturnErrorWhenUrlNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"makeiterror"
		}
	`
	mock.EXPECT().createNewVersionProcess("org", "test", "test", configView{Created_by: "tester"}).Return(1, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test/test", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}", handler.CreateNewVersionHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateNewVersionHandlerShouldReturnErrorWhenConfigEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id":1
		}
	`
	mock.EXPECT().createNewVersionProcess("org", "test", "test", configView{Created_by: "tester"}).Return(0, errors.New("error create query"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test/test", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}", handler.CreateNewVersionHandler).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateNewVersionHandlerShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id":1,
			"configuration": {
				"host": "localhost"
				}
		}
	`
	configuration_expected := make(map[string]string)
	configuration_expected["host"] = "localhost"
	mock.EXPECT().createNewVersionProcess("org", "test", "test", configView{Created_by: "tester", Configuration: configuration_expected}).Return(0, errors.New("error create query"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test/test", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}", handler.CreateNewVersionHandler).Methods("POST")
	r.ServeHTTP(rr, req.WithContext(context.WithValue(req.Context(), "username", "tester")))
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateNewVersionHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
	{
		"configuration": {
			"host": "localhost"
			}
	}
		`

	configuration_expected := make(map[string]string)
	configuration_expected["host"] = "localhost"
	mock.EXPECT().createNewVersionProcess("org", "test", "test", configView{Created_by: "tester", Configuration: configuration_expected}).Return(1, nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/org/test/test", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}", handler.CreateNewVersionHandler).Methods("POST")
	r.ServeHTTP(rr, req.WithContext(context.WithValue(req.Context(), "username", "tester")))
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetConfigurationVersionsHandlerShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id":1
		}
	`
	mock.EXPECT().getConfigurationVersionsProcess("org", "test", "test").Return("1", errors.New("error"))
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/versions", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/versions", handler.GetConfigurationVersionsHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetConfigurationVersionsHandlerShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockProcessor(ctrl)
	handler := mainConfiguration{processor: mock}
	payload := `
		{
			"user_id":1
		}
	`
	mock.EXPECT().getConfigurationVersionsProcess("org", "test", "test").Return("1", nil)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/org/test/test/versions", strings.NewReader(payload))
	r := mux.NewRouter()
	r.HandleFunc("/{organization_name}/{service_name}/{namespace}/versions", handler.GetConfigurationVersionsHandler).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
