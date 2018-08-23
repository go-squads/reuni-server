package configuration

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigurationProcessShouldReturnErrorWhenOrganizationNameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(0, errors.New("organization name doesnt exist"))
	mock.EXPECT().getConfiguration(1, "test-service", "test", 1)

	config, err := proc.getConfigurationProcess("test-organization", "test-service", "test", 1)
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetConfigurationProcessShouldReturnErrorWhenConfigDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getConfiguration(1, "test-service", "test", 1).Return(nil, errors.New("config not found"))

	config, err := proc.getConfigurationProcess("test-organization", "test-service", "test", 1)
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetConfigurationProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getConfiguration(1, "test-service", "test", 1).Return(&configView{}, nil)

	config, err := proc.getConfigurationProcess("test-organization", "test-service", "test", 1)
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestGetLatestVersionProcessShouldReturnErrorWhenOrganizationNameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(0, errors.New("organization name doesnt exist"))
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test")

	config, err := proc.getLatestVersionProcess("test-organization", "test-service", "test")
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetLatestVersionProcessShouldReturnErrorWhenConfigDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(0, errors.New("config not found"))

	config, err := proc.getLatestVersionProcess("test-organization", "test-service", "test")
	assert.Error(t, err)
	assert.Empty(t, config)
}

func TestGetLatestVersionProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(1, nil)

	config, err := proc.getLatestVersionProcess("test-organization", "test-service", "test")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenOrganizationNameDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(0, errors.New("error namespace doesnt exist"))
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test-service", "test", configView{Created_by: "tester"}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test-service", "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test-organization", "test", "test", configView{})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenNamespaceDoesntHaveLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(0, errors.New("error latest version not found"))
	mock.EXPECT().createNewVersion(1, "test-service", "test", configView{Created_by: "tester"}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test-service", "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test-organization", "test-service", "test", configView{Created_by: "tester"})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test-service", "test", configView{Created_by: "tester"}, 2).Return(errors.New("create query error"))
	mock.EXPECT().updateNamespaceActiveVersion(1, "test-service", "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test-organization", "test-service", "test", configView{Created_by: "tester"})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldReturnErrorWhenUpdateActiveVersionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test-service", "test", configView{Created_by: "tester"}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test-service", "test", 2).Return(errors.New("update version error"))

	version, err := proc.createNewVersionProcess("test-organization", "test-service", "test", configView{Created_by: "tester"})
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestCreateNewVersionProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getLatestVersionForNamespace(1, "test-service", "test").Return(1, nil)
	mock.EXPECT().createNewVersion(1, "test-service", "test", configView{Created_by: "tester"}, 2).Return(nil)
	mock.EXPECT().updateNamespaceActiveVersion(1, "test-service", "test", 2).Return(nil)

	version, err := proc.createNewVersionProcess("test-organization", "test-service", "test", configView{Created_by: "tester"})
	assert.NoError(t, err)
	assert.NotEmpty(t, version)
}

func TestGetConfigurationVersionsProcessShouldReturnErrorWhenServiceIdDoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, errors.New("organization cant be found"))
	mock.EXPECT().getVersions(1, "test-service", "test").Return([]int{1}, nil)

	version, err := proc.getConfigurationVersionsProcess("test-organization", "test-service", "test")
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestGetConfigurationVersionsProcessShouldReturnErrorWhenGetVersionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getVersions(1, "test-service", "test").Return(nil, errors.New("versions cant be found"))

	version, err := proc.getConfigurationVersionsProcess("test-organization", "test-service", "test")
	assert.Error(t, err)
	assert.Empty(t, version)
}

func TestGetConfigurationVersionsProcessShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockRepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().getOrganizationId("test-organization").Return(1, nil)
	mock.EXPECT().getVersions(1, "test-service", "test").Return([]int{1}, nil)

	version, err := proc.getConfigurationVersionsProcess("test-organization", "test-service", "test")
	assert.NoError(t, err)
	assert.NotNil(t, version)
}
