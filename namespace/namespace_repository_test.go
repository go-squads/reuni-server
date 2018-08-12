package namespace

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/appcontext"
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
	rep := initRepository(makeMockRow(map[string]interface{}{"count": int64(50)}, nil))
	res, err := rep.isNamespaceExist(1, "namespace")
	assert.True(t, res)
	assert.NoError(t, err)
}

func TestIsNamespaceExistReturnFalse(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"count": int64(0)}, nil))
	res, err := rep.isNamespaceExist(1, "namespace")
	assert.False(t, res)
	assert.NoError(t, err)
}

func TestIsNamespaceExistReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("testError")))
	res, err := rep.isNamespaceExist(1, "namespace")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestIsNamespaceExistReturnErrorWhenCoundNotInt(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"count": "x"}, nil))
	res, err := rep.isNamespaceExist(1, "namespace")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestCreateConfigurationShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	err := rep.createConfiguration(1, "name", map[string]interface{}{"DB_HOST": "123"})
	assert.NoError(t, err)
}

func TestCreateConfigurationShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("Test Error")))
	err := rep.createConfiguration(1, "name", map[string]interface{}{"DB_HOST": "123"})
	assert.Error(t, err)
}

func TestCreateConfigurationShouldReturnErrorWhenConfigNotMarshalable(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("Test Error")))
	err := rep.createConfiguration(1, "name", map[string]interface{}{"DB_HOST": make(chan int)})
	assert.Error(t, err)
}

func TestCreateNamespaceShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	err := rep.createNewNamespace(newNamespaceStruct(0, 1, "service", 1))
	assert.NoError(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("Test Error")))
	err := rep.createNewNamespace(newNamespaceStruct(0, 1, "service", 1))
	assert.Error(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenDataIsEmptyFirst(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	err := rep.createNewNamespace(newNamespaceStruct(0, 1, "", 1))
	assert.Error(t, err)
}

func TestCreateNamespaceShouldReturnErrorWhenDataIsEmptySecond(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	err := rep.createNewNamespace(newNamespaceStruct(0, 0, "service", 1))
	assert.Error(t, err)
}

func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnNoData(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))
	res, err := rep.retrieveAllNamespace(1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
	assert.Empty(t, res)
}

func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnOneData(t *testing.T) {
	rep := initRepository(makeMockRows(makeRows(map[string]interface{}{"namespace": "default", "active_version": 1}), nil))
	res, err := rep.retrieveAllNamespace(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "default", res[0].Namespace)
}
func TestRetrieveAllNamespacesShouldNotReturnErrorWhenQueryReturnMoreThanOneData(t *testing.T) {
	rep := initRepository(makeMockRows(makeRows(map[string]interface{}{"namespace": "default", "active_version": 1}, map[string]interface{}{"namespace": "production", "active_version": 5}), nil))
	res, err := rep.retrieveAllNamespace(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "default", res[0].Namespace)
	assert.Equal(t, "production", res[1].Namespace)

}

func TestRetrieveAllNamespacesShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("TestError")))
	res, err := rep.retrieveAllNamespace(1)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestRetrieveAllNamespaceShouldReturnErrorWhenQueryReturnResultThatCannotBeParsed(t *testing.T) {
	rep := initRepository(makeMockRows(makeRows(map[string]interface{}{"namespace": make(chan string), "active_version": 1}, map[string]interface{}{"namespace": "production", "active_version": 5}), nil))
	res, err := rep.retrieveAllNamespace(1)
	assert.Error(t, err)
	assert.Nil(t, res)

}

func TestGetServiceIdShouldReturnId(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": 2}, nil))
	appcontext.InitMockContext(
		rep.execer,
	)
	res, err := rep.getServiceId("test-service")
	assert.NoError(t, err)
	assert.Equal(t, 2, res)

}

func TestGetServiceIdShouldReturnErrorWhenNoData(t *testing.T) {
	rep := initRepository(makeMockRow(nil, nil))
	appcontext.InitMockContext(
		rep.execer,
	)
	res, err := rep.getServiceId("test-service2")
	assert.Error(t, err)
	assert.Empty(t, res)

}

func TestGetServiceIdShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	appcontext.InitMockContext(
		rep.execer,
	)
	res, err := rep.getServiceId("test-service2")
	assert.Error(t, err)
	assert.Empty(t, res)

}
