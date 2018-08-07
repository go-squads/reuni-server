package namespace

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
func makeRows(args ...map[string]interface{}) []map[string]interface{} {
	return args
}

func TestIsNamespaceExistReturnTrue(t *testing.T) {
	q := makeMockRow(map[string]interface{}{"count": 50}, nil)
	res, err := isNamespaceExist(q, 1, "namespace")
	assert.True(t, res)
	assert.NoError(t, err)
}

func TestIsNamespaceExistReturnFalse(t *testing.T) {
	q := makeMockRow(map[string]interface{}{"count": 0}, nil)
	res, err := isNamespaceExist(q, 1, "namespace")
	assert.False(t, res)
	assert.NoError(t, err)
}

func TestIsNamespaceExistReturnErrorWhenQueryError(t *testing.T) {
	q := makeMockRow(nil, errors.New("testError"))
	res, err := isNamespaceExist(q, 1, "namespace")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestIsNamespaceExistReturnErrorWhenCoundNotInt(t *testing.T) {
	q := makeMockRow(map[string]interface{}{"count": "x"}, nil)
	res, err := isNamespaceExist(q, 1, "namespace")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestCreateConfigurationShouldNotReturnError(t *testing.T) {
	q := makeMockRow(nil, nil)
	err := createConfiguration(q, 1, "name", map[string]interface{}{"DB_HOST": "123"})
	assert.NoError(t, err)
}

func TestCreateConfigurationShouldReturnErrorWhenQueryError(t *testing.T) {
	q := makeMockRow(nil, errors.New("Test Error"))
	err := createConfiguration(q, 1, "name", map[string]interface{}{"DB_HOST": "123"})
	assert.Error(t, err)
}

func TestCreateConfigurationShouldReturnErrorWhenConfigNotMarshalable(t *testing.T) {
	q := makeMockRow(nil, errors.New("Test Error"))
	err := createConfiguration(q, 1, "name", map[string]interface{}{"DB_HOST": make(chan int)})
	assert.Error(t, err)
}

func TestCreateNamespaceShouldNotReturnError(t *testing.T) {
	q := makeMockRow(nil, nil)
	err := createNewNamespace(q, newNamespaceStruct(0, 1, "service", 1))
	assert.NoError(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	q := makeMockRow(nil, errors.New("Test Error"))
	err := createNewNamespace(q, newNamespaceStruct(0, 1, "service", 1))
	assert.Error(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenDataIsEmptyFirst(t *testing.T) {
	q := makeMockRow(nil, nil)
	err := createNewNamespace(q, newNamespaceStruct(0, 1, "", 1))
	assert.Error(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenDataIsEmptySecond(t *testing.T) {
	q := makeMockRow(nil, nil)
	err := createNewNamespace(q, newNamespaceStruct(0, 0, "service", 1))
	assert.Error(t, err)
}

func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnNoData(t *testing.T) {
	q := makeMockRows(nil, nil)
	res, err := retrieveAllNamespace(q, 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
	assert.Empty(t, res)
}

func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnOneData(t *testing.T) {
	q := makeMockRows(makeRows(map[string]interface{}{"namespace": "default", "active_version": 1}), nil)
	res, err := retrieveAllNamespace(q, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "default", res[0].Namespace)
}
func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnMoreThanOneData(t *testing.T) {
	q := makeMockRows(makeRows(map[string]interface{}{"namespace": "default", "active_version": 1}, map[string]interface{}{"namespace": "production", "active_version": 5}), nil)
	res, err := retrieveAllNamespace(q, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "default", res[0].Namespace)
	assert.Equal(t, "production", res[1].Namespace)

}

func TestRetrieveAllNamespacesShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	q := makeMockRows(nil, errors.New("TestError"))
	res, err := retrieveAllNamespace(q, 1)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestRetrieveAllNamespaceShouldReturnErrorWhenQueryReturnResultThatCannotBeParsed(t *testing.T) {
	q := makeMockRows(makeRows(map[string]interface{}{"namespace": make(chan string), "active_version": 1}, map[string]interface{}{"namespace": "production", "active_version": 5}), nil)
	res, err := retrieveAllNamespace(q, 1)
	assert.Error(t, err)
	assert.Nil(t, res)

}
