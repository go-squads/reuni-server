package configuration

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/helper"
	"github.com/stretchr/testify/assert"
)

func makeMockRow(data map[string]interface{}, err error) *helper.QueryMockHelper {
	return &helper.QueryMockHelper{
		Row: data,
		Err: err,
	}
}
func makeMockRows(data []map[string]interface{}, err error) *helper.QueryMockHelper {
	return &helper.QueryMockHelper{
		Data: data,
		Err:  err,
	}
}
func initRepository(q helper.QueryExecuter) Repository {
	return &mainRepository{
		execer: q,
	}
}
func TestGetConfigurationRepositoryShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("Test Error")))
	res, err := rep.getConfiguration(1, "test-sercvices", 1)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetConfigurationRepositoryShouldReturnErrorWhenReturnCannotBeParsed(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"test": make(chan bool)}, nil))
	res, err := rep.getConfiguration(1, "test-sercvices", 1)
	assert.Nil(t, res)
	assert.Error(t, err)
}
func TestGetConfigurationRepositoryShouldNotReturnErrorWhenQueryNil(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	res, err := rep.getConfiguration(1, "test-services", 1)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetConfigurationRepositoryShouldNotReturnErrorWhenQueryReturnEmptyRow(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{}, nil))
	res, err := rep.getConfiguration(1, "test-services", 1)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestGetConfigurationRepositoryShouldNotReturnErrorWhenQueryReturnData(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"version": 1, "configs": []byte(`{"test":"123"}`)}, nil))
	res, err := rep.getConfiguration(1, "test-services", 1)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}
