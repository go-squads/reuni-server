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
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(0), errors.New("Test Error"))
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldReturnErrorWhenAddUserReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(errors.New("Test Error"))
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldNotReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	activeRepository = mock
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(nil)
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.NoError(t, err)
}

func TestGetRepositoryWhenActiveRepoNil(t *testing.T) {
	appcontext.InitMockContext(&helper.QueryMockHelper{Data: nil, Err: nil})
	activeRepository = nil
	repo := getRepository()
	assert.NotNil(t, repo)

}

func TestAddUserProcessorShouldNotReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	activeRepository = mock
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "Admin",
	}
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(nil)
	err := proc.addUserProcessor(member)
	assert.NoError(t, err)
}

func TestAddUserProcessorShouldReturnError(t *testing.T) {
	proc := mainProcessor{}
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "aosdkaos",
	}
	err := proc.addUserProcessor(member)
	assert.Error(t, err)
}

func TestDeleteUserProcessorShouldNotReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	activeRepository = mock

	mock.EXPECT().deleteUser(int64(1), int64(1)).Return(nil)
	err := proc.deleteUserFromGroupProcessor(int64(1), int64(1))
	assert.NoError(t, err)
}

func TestDeleteUserProcessorShouldReturnError(t *testing.T) {
	proc := mainProcessor{}
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "aosdkaos",
	}
	err := proc.addUserProcessor(member)
	assert.Error(t, err)
}

func TestUpdateRoleOfUserProcessorShouldNotReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	activeRepository = mock
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "Developer",
	}
	mock.EXPECT().updateRoleOfUser("Developer", int64(1), int64(1)).Return(nil)
	err := proc.updateRoleOfUserProcessor(member)
	assert.NoError(t, err)
}

func TestUpdateRoleOfUserProcessorShouldReturnError(t *testing.T) {
	proc := mainProcessor{}
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "aosdkaos",
	}
	err := proc.updateRoleOfUserProcessor(member)
	assert.Error(t, err)
}

func TestGetAllMemberOfOrganizationProcessorShouldNotReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	activeRepository = mock
	mock.EXPECT().getAllMemberOfOrganization(int64(1)).Return([]map[string]interface{}{}, nil)
	data, err := proc.getAllMemberOfOrganizationProcessor(int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetAllMemberOfOrganizationProcessorShouldReturnError(t *testing.T) {
	proc := mainProcessor{}
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	activeRepository = mock
	mock.EXPECT().getAllMemberOfOrganization(int64(1)).Return(nil, errors.New("Internal error"))
	data, err := proc.getAllMemberOfOrganizationProcessor(int64(1))
	assert.Error(t, err)
	assert.Nil(t, data)
}
