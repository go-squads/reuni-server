package organization

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
