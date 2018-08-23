package namespace

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/helper"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNamespaceProcessorShouldReturnErrorWhenOrganizationIsNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(0, errors.New("Service Not Exist"))
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenNamespaceIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenNamespaceIsExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "test-service", "default").Return(true, nil)
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenIsNamespaceExistReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "test-service", "default").Return(false, errors.New("Internal Error"))
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenCreateNamespaceReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "test-service", "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(1, "test-service", "default", 1)).Return(errors.New("Internal Error"))
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}

func TestCreateNamespaceProcessorShouldReturnErrorWhenCreateConfigurationReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "test-service", "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(1, "test-service", "default", 1)).Return(nil)
	mock.EXPECT().createConfiguration(1, "test-service", "default", map[string]interface{}{"DB_HOST": "127.0.0.1"}).Return(errors.New("Internal Error"))
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.Error(t, err)
}
func TestCreateNamespaceProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().isNamespaceExist(1, "test-service", "default").Return(false, nil)
	mock.EXPECT().createNewNamespace(newNamespaceStruct(1, "test-service", "default", 1)).Return(nil)
	mock.EXPECT().createConfiguration(1, "test-service", "default", map[string]interface{}{"DB_HOST": "127.0.0.1"}).Return(nil)
	err := procs.createNewNamespaceProcessor("test-organization", "test-service", newNamespaceViewStruct("default", map[string]interface{}{"DB_HOST": "127.0.0.1"}))
	assert.NoError(t, err)
}

func TestRetrieveAllNamespaceProcessorShouldReturnErrorWhenOrganizationNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(0, helper.NewHttpError(404, "Not Found"))
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().retrieveAllNamespace(1, "test-service").Return([]namespaceStore{}, nil)
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.NoError(t, err)
	assert.Equal(t, []byte("[]"), res)
}
func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultOneEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, "test-service", "default", 1)}

	mock.EXPECT().retrieveAllNamespace(1, "test-service").Return(data, nil)
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllNamespaceProcessorShouldNotReturnErrorWhenResultMoreThanOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, "test-service", "default", 1), *newNamespaceStruct(2, "test-service", "prod", 2)}

	mock.EXPECT().retrieveAllNamespace(1, "test-service").Return(data, nil)
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllNamespaceProcessorShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	data := []namespaceStore{*newNamespaceStruct(1, "test-service", "default", 1), *newNamespaceStruct(2, "test-service", "prod", 2)}
	mock.EXPECT().retrieveAllNamespace(1, "test-service").Return(data, errors.New("Error"))
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAllNamespaceProcessorShouldReturnEmptyArrayWhenRetrieveReturnNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMocknamespaceRepositoryInterface(ctrl)
	procs := mainProcessor{repo: mock}
	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().retrieveAllNamespace(1, "test-service").Return(nil, nil)
	res, err := procs.retrieveAllNamespaceProcessor("test-organization", "test-service")
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
