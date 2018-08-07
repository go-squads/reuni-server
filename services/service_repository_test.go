package services

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/helper"
	"github.com/stretchr/testify/assert"
)

func MockServiceMap(id int, name string) map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = id
	m["name"] = name
	return m
}

func MockServiceStruct(id int, name string) service {
	return service{
		Id:   id,
		Name: name,
	}
}
func TestGetAllServiceShouldNotReturnErrorWhenQueryReturnOneData(t *testing.T) {
	var data []map[string]interface{}
	data = append(data, MockServiceMap(1, "go-pay-service"))

	mock := &helper.QueryMockHelper{
		Data: data,
		Err:  nil,
	}
	services, err := getAll(mock)
	var expected []service

	expected = append(expected, MockServiceStruct(1, "go-pay-service"))
	assert.Equal(t, services, expected)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenDataMoreThanOne(t *testing.T) {

	var data []map[string]interface{}
	data = append(data, MockServiceMap(1, "go-pay-service"))
	data = append(data, MockServiceMap(2, "go-ride-service"))

	mock := &helper.QueryMockHelper{
		Data: data,
		Err:  nil,
	}
	services, err := getAll(mock)
	var expected []service
	expected = append(expected, MockServiceStruct(1, "go-pay-service"))
	expected = append(expected, MockServiceStruct(2, "go-ride-service"))
	assert.Equal(t, services, expected)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenQueryDoesNotReturnData(t *testing.T) {
	var data []map[string]interface{}
	mock := &helper.QueryMockHelper{
		Data: data,
		Err:  nil,
	}
	services, err := getAll(mock)
	assert.Empty(t, services)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("Query Return Error"),
	}
	services, err := getAll(mock)
	assert.Nil(t, services)
	assert.Error(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenDataNotParseableToStruct(t *testing.T) {
	var data []map[string]interface{}
	row := make(map[string]interface{})
	row["test"] = 1
	data = append(data, row)
	mock := &helper.QueryMockHelper{
		Data: data,
		Err:  nil,
	}
	var expected []service
	expected = append(expected, MockServiceStruct(0, ""))
	services, err := getAll(mock)
	assert.Equal(t, expected, services)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldReturnErrorWhenDataNotMarshalable(t *testing.T) {
	var data []map[string]interface{}
	row := make(map[string]interface{})
	row["test"] = make(chan int)
	data = append(data, row)
	mock := &helper.QueryMockHelper{
		Data: data,
		Err:  nil,
	}
	services, err := getAll(mock)
	assert.Nil(t, services)
	assert.Error(t, err)
}

func TestCreateServiceShouldNotReturnErrorWhenQueryNotReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}

	err := createService(mock, service{Name: "Hello", AuthorizationToken: "World"})
	assert.NoError(t, err)
}

func TestCreateServiceShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("This is Error"),
	}

	err := createService(mock, service{Name: "Hello", AuthorizationToken: "World"})
	assert.Error(t, err)
}

func TestDeleteServiceShouldNotReturnErrorWhenQueryNotReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}

	err := deleteService(mock, service{Name: "Hello"})
	assert.NoError(t, err)
}

func TestFindOneServiceByNameShouldNotReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Err: nil,
	}
	var data = make(map[string]interface{})
	data["id"] = 7
	data["name"] = "New_Service"
	data["authorization_token"] = ""
	data["created_at"] = nil
	mock.Data = []map[string]interface{}{data}
	s, err := findOneServiceByName(mock, "Hello")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
func TestFindOneServiceByNameShouldReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Err: errors.New("Sample Error"),
	}
	s, err := findOneServiceByName(mock, "Hello")
	assert.Error(t, err)
	assert.Empty(t, s)
}

func TestFindOneServiceByNameShouldReturnErrorWhenDataIsEmpty(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	s, err := findOneServiceByName(mock, "Hello")
	assert.Error(t, err)
	assert.Empty(t, s)
}

func TestFindOneServiceByNameShouldReturnErrorIfDataNotMarshalable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["test"] = make(chan int)
	mock := &helper.QueryMockHelper{
		Data: []map[string]interface{}{datum},
		Err:  nil,
	}
	s, err := findOneServiceByName(mock, "Hello")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetTokenShouldNotReturnError(t *testing.T) {
	mock := &helper.QueryMockHelper{}
	var data = make(map[string]interface{})
	data["authorization_token"] = "asdfsdfa"
	mock.Data = []map[string]interface{}{data}
	token, err := getServiceToken(mock, "Hello")
	assert.NoError(t, err)
	assert.Equal(t, "asdfsdfa", token.Token)
}

func TestGetTokenShouldReturnErrorWhenThereIsNodata(t *testing.T) {
	mock := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	services, err := getServiceToken(mock, "hello")
	assert.Empty(t, services)
	assert.Error(t, err)
}

func TestGetTokenShouldReturnErrorWhenDataNotMarshalable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["test"] = make(chan int)
	mock := &helper.QueryMockHelper{
		Data: []map[string]interface{}{datum},
		Err:  nil,
	}
	services, err := getServiceToken(mock, "hello")
	assert.Empty(t, services)
	assert.Error(t, err)
}
