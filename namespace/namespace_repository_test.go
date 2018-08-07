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
