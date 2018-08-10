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
	"github.com/stretchr/testify/assert"
)

func ServeRequest(rr *httptest.ResponseRecorder, req *http.Request, handler http.HandlerFunc) {
	handler.ServeHTTP(rr, req)
}
func TestGetAllHandlerShouldNotPanic(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test-service")},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	req, err := http.NewRequest("GET", "/services", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	ServeRequest(rr, req, GetAllServicesHandler)
	assert.Equal(t, rr.Code, http.StatusOK)
	exp, _ := json.Marshal([]service{MockServiceStruct(1, "test-service")})
	assert.Equal(t, string(exp), rr.Body.String())
}

func TestGetAllHandlerShouldReturnErrorWhenBadRequest(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  helper.NewHttpError(400, "Bad Request"),
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services", nil)
	ServeRequest(rr, req, GetAllServicesHandler)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetAllHandlerShouldReturnError500WhenObjectNotMarshalable(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"test": make(chan int)}},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services", nil)
	ServeRequest(rr, req, GetAllServicesHandler)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// func TestCreateServiceHandlerShouldNotReturnError(t *testing.T) {
// 	q := &helper.QueryMockHelper{
// 		Data: []map[string]interface{}{map[string]interface{}{"test": make(chan int)}},
// 		Err:  nil,
// 	}
// 	payload := `
// 		{
// 			"name": "test-service"
// 		}
// 	`
// 	appcontext.InitMockContext(q)
// 	var rr = httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "1/services", strings.NewReader(payload))
// 	ServeRequest(rr, req, CreateServiceHandler)
// 	assert.Equal(t, http.StatusCreated, rr.Code)
// }

// func TestCreateServiceHandlerShouldReturnErrorWhenQueryReturnError(t *testing.T) {
// 	q := &helper.QueryMockHelper{
// 		Data: nil,
// 		Err:  helper.NewHttpError(500, "Test"),
// 	}
// 	payload := `
// 		{
// 			"name": "test-service"
// 		}
// 	`
// 	appcontext.InitMockContext(q)
// 	var rr = httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/services", strings.NewReader(payload))
// 	ServeRequest(rr, req, CreateServiceHandler)
// 	assert.Equal(t, http.StatusInternalServerError, rr.Code)
// }

// func TestCreateServiceHandlerShouldReturnErrorWhenPayloadIsempty(t *testing.T) {
// 	q := &helper.QueryMockHelper{
// 		Data: nil,
// 		Err:  nil,
// 	}
// 	payload := `
// 		{
// 			"name": ""
// 		}
// 	`
// 	appcontext.InitMockContext(q)
// 	var rr = httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/services", strings.NewReader(payload))
// 	ServeRequest(rr, req, CreateServiceHandler)
// 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// }

// func TestCreateServiceHandlerShouldReturnErrorWhenPayloadMalformed(t *testing.T) {
// 	q := &helper.QueryMockHelper{
// 		Data: nil,
// 		Err:  nil,
// 	}
// 	payload := `
// 		{
// 			name: ""
// 			"asdsd
// 		}
// 	`
// 	appcontext.InitMockContext(q)
// 	var rr = httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/services", strings.NewReader(payload))
// 	ServeRequest(rr, req, CreateServiceHandler)
// 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// }

// func TestCreateServiceHandlerShouldReturnErrorServiceExist(t *testing.T) {
// 	q := &helper.QueryMockHelper{
// 		Data: []map[string]interface{}{MockServiceMap(1, "test-service")},
// 		Err:  nil,
// 	}
// 	payload := `
// 		{
// 			"name": "test-service"
// 		}
// 	`
// 	appcontext.InitMockContext(q)
// 	var rr = httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/services", strings.NewReader(payload))
// 	ServeRequest(rr, req, CreateServiceHandler)
// 	assert.Equal(t, http.StatusConflict, rr.Code)
// }

func TestDeleteServiceHandlerShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test-service")},
		Err:  nil,
	}
	payload := `
		{
			"name": "test-service"
		}
	`
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/services", strings.NewReader(payload))
	ServeRequest(rr, req, DeleteServiceHandler)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestDeleteServiceHandlerShouldNotReturnErrorWhenServiceNotExist(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	payload := `
		{
			"name": "test-service"
		}
	`
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/services", strings.NewReader(payload))
	ServeRequest(rr, req, DeleteServiceHandler)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteServiceHandlerShouldReturnErrorWhenPayloadNotUnmarshalable(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	payload := `
		{
			Test:
			"name": "test-service",
		}
	`
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/services", strings.NewReader(payload))
	ServeRequest(rr, req, DeleteServiceHandler)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteServiceHandlerShouldReturnErrorWhenProcessorError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  helper.NewHttpError(http.StatusInternalServerError, "Test Error"),
	}
	payload := `
		{
			"name": "test-service"
		}
	`
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/services", strings.NewReader(payload))
	ServeRequest(rr, req, DeleteServiceHandler)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestValidateTokenShouldReturnValidTrue(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "HelloWorld!"}},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/test-services/validatetoken", nil)
	req.Header.Set("Authorization", "HelloWorld!")
	ServeRequest(rr, req, ValidateToken)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]bool
	json.NewDecoder(rr.Body).Decode(&data)
	assert.True(t, data["valid"])
}

func TestValidateTokenShouldReturnValidFalse(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "HelloWorld!"}},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/test-services/validatetoken", nil)
	req.Header.Set("Authorization", "HelloWorld!!!!!!!")
	ServeRequest(rr, req, ValidateToken)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]bool
	json.NewDecoder(rr.Body).Decode(&data)
	assert.False(t, data["valid"])
}

func TestValidateTokenShouldReturnErrorWhenProcessorError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  helper.NewHttpError(500, "TestError"),
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/test-services/validatetoken", nil)
	ServeRequest(rr, req, ValidateToken)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetTokenHandlerShouldReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "testToken"}},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/test-services/token", nil)
	ServeRequest(rr, req, GetToken)
	assert.Equal(t, http.StatusOK, rr.Code)
	var data map[string]string
	json.NewDecoder(rr.Body).Decode(&data)
	assert.Equal(t, "testToken", data["authorization_token"])
}

func TestGetTokenHandlerShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"authorization_token": "testToken"}},
		Err:  helper.NewHttpError(500, "Bad Error"),
	}
	appcontext.InitMockContext(q)
	var rr = httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/test-services/token", nil)
	ServeRequest(rr, req, GetToken)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	var data map[string]string
	json.NewDecoder(rr.Body).Decode(&data)
	assert.Equal(t, "", data["authorization_token"])
}

func TestGetUsernameShouldReturnUsername(t *testing.T) {
	r, _ := http.NewRequest("GET", "/services/test-services/token", nil)
	ctx := context.WithValue(r.Context(), "username", "go-squads")
	r = r.WithContext(ctx)
	assert.Equal(t, "go-squads", getUsername(r))
}

func TestGetUsernameShouldReturnEmptyString(t *testing.T) {
	r, _ := http.NewRequest("GET", "/services/test-services/token", nil)
	assert.Empty(t, getUsername(r))
}
