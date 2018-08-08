package namespace

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/appcontext"

	"github.com/go-squads/reuni-server/helper"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNamespaceProcessorShouldReturnErrorWhenServiceIsNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(0, errors.New("Service Not Exist"))
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenNamespaceIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenNamespaceIsExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "default").Return(true, nil)
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenIsNamespaceExistReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "default").Return(false, errors.New("Internal Error"))
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenCreateNamespaceReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(0, 1, "default", 1)).Return(errors.New("Internal Error"))
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenCreateConfigurationReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(0, 1, "default", 1)).Return(nil)
	mock.EXPECT().createConfiguration(1, "default", map[string]interface{}{"DB_HOST": "127.0.0.1"}).Return(errors.New("Internal Error"))
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}
func TestCreateNamespaceProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(0, 1, "default", 1)).Return(nil)
	mock.EXPECT().createConfiguration(1, "default", map[string]interface{}{"DB_HOST": "127.0.0.1"}).Return(nil)
	err := createNewNamespaceProcessor("test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.NoError(t, err)
}

func TestRetrieveAllNamespaceProcessorShouldReturnErrorWhenServiceNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(0, helper.NewHttpError(404, "Not Found"))
	res, err := retrieveAllNamespaceProcessor("test-service")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	mock.EXPECT().retrieveAllNamespace(1).Return([]namespaceStore{}, nil)
	res, err := retrieveAllNamespaceProcessor("test-service")
	assert.NoError(t, err)
	assert.Equal(t, []byte("[]"), res)
}
func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultOneEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, 1, "default", 1)}

	mock.EXPECT().retrieveAllNamespace(1).Return(data, nil)
	res, err := retrieveAllNamespaceProcessor("test-service")
	assert.NoError(t, err)
	assert.Equal(t, `[{"id":1,"service_id":1,"namespace":"default","version":1}]`, string(res))
}

func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultMoreThanOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, 1, "default", 1), *newNamespaceStruct(2, 1, "prod", 2)}

	mock.EXPECT().retrieveAllNamespace(1).Return(data, nil)
	res, err := retrieveAllNamespaceProcessor("test-service")
	assert.NoError(t, err)
	assert.Equal(t, `[{"id":1,"service_id":1,"namespace":"default","version":1},{"id":2,"service_id":1,"namespace":"prod","version":2}]`, string(res))
}

func TestGetAllNamespaceProcessorShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockNamespaceRepository(ctrl)
	activeRepo = mock
	mock.EXPECT().getServiceId("test-service").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, 1, "default", 1), *newNamespaceStruct(2, 1, "prod", 2)}

	mock.EXPECT().retrieveAllNamespace(1).Return(data, errors.New("Error"))
	res, err := retrieveAllNamespaceProcessor("test-service")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetActiveRepositoryWithMockContext(t *testing.T) {
	appcontext.InitMockContext(&helper.QueryMockHelper{
		Data: []map[string]interface{}{map[string]interface{}{"Test": 123}},
	})
	activeRepo = nil
	res := getActiveRepo()
	mock, ok := res.(*namespaceRepository)
	assert.True(t, ok)
	assert.NotNil(t, mock)
}
