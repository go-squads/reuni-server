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
	proc := mainProcessor{repo: mock}
	mock.EXPECT().createNewOrganization("test").Return(int64(0), errors.New("Test Error"))
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldReturnErrorWhenAddUserReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	proc := mainProcessor{repo: mock}
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(errors.New("Test Error"))
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.Error(t, err)

}

func TestCreateOrganizationProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)
	proc := mainProcessor{repo: mock}
	mock.EXPECT().createNewOrganization("test").Return(int64(1), nil)
	mock.EXPECT().addUser(int64(1), int64(1), "Admin").Return(nil)
	err := proc.createNewOrganizationProcessor("test", int64(1))
	assert.NoError(t, err)
}
func TestAddUserProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
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
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}

	mock.EXPECT().deleteUser(int64(1), int64(1)).Return(nil)
	err := proc.deleteUserFromGroupProcessor(int64(1), int64(1))
	assert.NoError(t, err)
}

func TestDeleteUserProcessorShouldReturnError(t *testing.T) {
	member := &Member{
		OrgId:  int64(1),
		UserId: int64(1),
		Role:   "aosdkaos",
	}
	proc := mainProcessor{repo: nil}
	err := proc.addUserProcessor(member)
	assert.Error(t, err)
}

func TestUpdateRoleOfUserProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
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
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().getAllMemberOfOrganization(int64(1)).Return([]map[string]interface{}{}, nil)
	data, err := proc.getAllMemberOfOrganizationProcessor(int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetAllMemberOfOrganizationProcessorShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().getAllMemberOfOrganization(int64(1)).Return(nil, errors.New("Internal error"))
	data, err := proc.getAllMemberOfOrganizationProcessor(int64(1))
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestGetAllOrganizationProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().getAllOrganization(1).Return([]OrganizationMember{}, nil)
	data, err := proc.getAllOrganizationProcessor(1)
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetAllOrganizationProcessorShouldReturnErrorWhenQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().getAllOrganization(1).Return(nil, errors.New("internal error"))
	data, err := proc.getAllOrganizationProcessor(1)
	assert.Error(t, err)
	assert.Empty(t, data)
}
func TestGetAllOrganizationProcessorShouldNotReturnErrorWhenQueryReturnEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().getAllOrganization(1).Return(nil, nil)
	data, err := proc.getAllOrganizationProcessor(1)
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestTranslateNameToIdProcShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().translateNameToIdRepository("test").Return(1, nil)
	data, err := proc.translateNameToIdProcessor("test")
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestTranslateNameToIdProcShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockrepository(ctrl)

	proc := mainProcessor{repo: mock}
	mock.EXPECT().translateNameToIdRepository("test").Return(0, errors.New(""))
	data, err := proc.translateNameToIdProcessor("test")
	assert.Error(t, err)
	assert.Empty(t, data)
}

func TestTranslateNameToIdProcWithoutReceiverShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Row: map[string]interface{}{
			"id": int64(1),
		},
		Err: nil,
	}
	appcontext.InitMockContext(q)

	data, err := TranslateNameToIdProcessor(q, "test")
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestTranslateNameToIdProcWithoutReceiverShouldReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Row: map[string]interface{}{
			"id": int64(0),
		},
		Err: errors.New("error"),
	}
	appcontext.InitMockContext(q)
	data, err := TranslateNameToIdProcessor(q, "test")
	assert.Error(t, err)
	assert.Empty(t, data)
}
