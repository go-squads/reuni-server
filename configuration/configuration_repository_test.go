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
func makeMockSlice(data []interface{}, err error) *helper.QueryMockHelper {
	return &helper.QueryMockHelper{
		Slice: data,
		Err:   err,
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

func TestGetLatestVersionForNamespaceShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	res, err := rep.getLatestVersionForNamespace(1, "test-services")
	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestGetLatestVersionForNamespaceShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"latest": int64(1), "configs": []byte(`{"test":"123"}`)}, nil))
	res, err := rep.getLatestVersionForNamespace(1, "test-services")
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCreateNewVersionShouldNotReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	err := rep.createNewVersion(1, "test-services", configView{}, 1)
	assert.Error(t, err)
}

func TestCreateNewVersionShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"latest": int64(1), "configs": []byte(`{"test":"123"}`)}, nil))
	err := rep.createNewVersion(1, "test-services", configView{}, 1)
	assert.NoError(t, err)
}

func TestUpdateNamespaceActiveVersionShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	err := rep.updateNamespaceActiveVersion(1, "test-services", 1)
	assert.Error(t, err)
}

func TestUpdateNamespaceActiveVersionShoulNotdReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"latest": int64(1), "configs": []byte(`{"test":"123"}`)}, nil))
	err := rep.updateNamespaceActiveVersion(1, "test-services", 1)
	assert.NoError(t, err)
}

func TestGetVersionsShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockSlice(nil, errors.New("error")))
	version, err := rep.getVersions(1, "test-services")
	assert.Nil(t, version)
	assert.Error(t, err)
}

func TestGetVersionsShouldReturnErrorWhenDataCantBeParsed(t *testing.T) {
	var data []interface{}
	data = append(data, "error")
	rep := initRepository(makeMockSlice(data, nil))
	version, err := rep.getVersions(1, "test-services")
	assert.Nil(t, version)
	assert.Error(t, err)
}

func TestGetVersionsShouldNotReturnError(t *testing.T) {
	var data []interface{}
	data = append(data, 1)
	rep := initRepository(makeMockSlice(data, nil))
	version, err := rep.getVersions(1, "test-services")
	assert.NotNil(t, version)
	assert.NoError(t, err)
}

func TestGetServiceIdShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	serviceId, err := rep.getServiceId("test-services")
	assert.Empty(t, serviceId)
	assert.Error(t, err)
}

func TestGetServiceIdShouldReturnErrorWhenIdNotFound(t *testing.T) {
	data := make(map[string]interface{})
	rep := initRepository(makeMockRow(data, nil))
	serviceId, err := rep.getServiceId("test-services")
	assert.Empty(t, serviceId)
	assert.Error(t, err)
}

func TestGetServiceIdShouldNotReturnError(t *testing.T) {
	data := make(map[string]interface{})
	data["id"] = 1
	rep := initRepository(makeMockRow(data, nil))
	serviceId, err := rep.getServiceId("test-services")
	assert.NotNil(t, serviceId)
	assert.NoError(t, err)
}
