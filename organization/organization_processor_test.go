package organization

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrganizationProcessorShouldReturnErrorWhenCannotCreateOrgnazation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(0), errors.New("Test Error"))
	err := createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldReturnErrorWhenAddUserReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(errors.New("Test Error"))
	err := createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(nil)
	err := createNewOrganizationProcessor("test", int64(1))
	assert.NoError(t, err)
}

func TestGetRepositoryWhenActiveRepoNil(t *testing.T) {
	appcontext.InitMockContext(&helper.QueryMockHelper{Data: nil, Err: nil})
	activeRepository = nil
	repo := getRepository()
	assert.NotNil(t, repo)

}
