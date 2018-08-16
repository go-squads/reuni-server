package configuration

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigurationProcessShouldReturnErrorWhenServiceNameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(0, errors.New("service name doesnt exist"))
	mock.EXPECT().getConfiguration(1, "test", 1)

	config, err := proc.getConfigurationProcess("test", "test", 1)
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetConfigurationProcessShouldReturnErrorWhenConfigDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getConfiguration(1, "test", 1).Return(nil, errors.New("config not found"))

	config, err := proc.getConfigurationProcess("test", "test", 1)
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetConfigurationProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getConfiguration(1, "test", 1).Return(&configView{}, nil)

	config, err := proc.getConfigurationProcess("test", "test", 1)
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestGetLatestVersionProcessShouldReturnErrorWhenServiceNameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(0, errors.New("service name doesnt exist"))
	mock.EXPECT().getLatestVersionForNamespace(1, "test")

	config, err := proc.getLatestVersionProcess("test", "test")
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetLatestVersionProcessShouldReturnErrorWhenConfigDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(0, errors.New("config not found"))

	config, err := proc.getLatestVersionProcess("test", "test")
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetLatestVersionProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(1, nil)

	config, err := proc.getLatestVersionProcess("test", "test")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenServicenameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(0, errors.New("error namespace doesnt exist"))
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test", configView{}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test", "test", configView{})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenNamespaceDoesntHaveLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(0, errors.New("error latest version not found"))
	mock.EXPECT().createNewVersion(1, "test", configView{}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test", "test", configView{})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test", configView{}, 2).Return(errors.New("create query error"))
	mock.EXPECT().updateNamespaceActiveVersion(1, "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test", "test", configView{})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenUpdateActiveVersionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test", configView{}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test", 2).Return(errors.New("update version error"))

	version, err := proc.createNewVersionProcess("test", "test", configView{})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test", configView{}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test", "test", configView{})
	assert.NoError(t, err)
	assert.NotEmpty(t, version)
}

func TestGetConfigurationVersionsProcessShouldReturnErrorWhenServiceIdDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, errors.New("serviceid cant be found"))
	mock.EXPECT().getVersions(1, "test").Return([]int{1}, nil)

	version, err := proc.getConfigurationVersionsProcess("test", "test")
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestGetConfigurationVersionsProcessShouldReturnErrorWhenGetVersionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getVersions(1, "test").Return(nil, errors.New("versions cant be found"))

	version, err := proc.getConfigurationVersionsProcess("test", "test")
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestGetConfigurationVersionsProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getServiceId("test").Return(1, nil)
	mock.EXPECT().getVersions(1, "test").Return([]int{1}, nil)

	version, err := proc.getConfigurationVersionsProcess("test", "test")
	assert.NoError(t, err)
	assert.NotNil(t, version)
}
