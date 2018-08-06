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
