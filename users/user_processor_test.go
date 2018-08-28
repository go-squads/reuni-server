package users

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}

	assert.NotNil(t, proc.getRepository())
}

func TestCreateUserEncryptPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}

	assert.NotNil(t, proc.createUserEncryptPassword("username", "password"))
}

func TestCreateUserProcessorShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)
	proc := userProcessor{repo: mock}
	mock.EXPECT().createUser(user{Name: "test", Username: "test", Email: "test", Password: "123hash"}).Return(nil)
	err := proc.createUserProcessor(userv{Name: "test", Username: "test", Email: "test", Password: "123hash"})
	assert.NoError(t, err)
}

func TestCreateUserProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)
	proc := userProcessor{repo: mock}
	mock.EXPECT().createUser(user{Name: "test", Username: "test", Email: "test", Password: "123hash"}).Return(errors.New("internal error"))
	err := proc.createUserProcessor(userv{Name: "test", Username: "test", Email: "test", Password: "123hash"})
	assert.Error(t, err)
}

func TestLoginUserProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	mock.EXPECT().loginUser(userv{Username: "test", Password: "test"}).Return([]byte{}, []byte{}, nil)

	data, dataRefreshToken, err := proc.loginUserProcessor(userv{Username: "test", Password: "test"})
	assert.NotNil(t, data)
	assert.NoError(t, err)
	assert.NotNil(t, dataRefreshToken)
}

func TestLoginUserProcessorShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	mock.EXPECT().loginUser(userv{Username: "test", Password: "test"}).Return(nil, nil, errors.New("error login"))

	data, dataRefreshToken, err := proc.loginUserProcessor(userv{Username: "test", Password: "test"})
	assert.Nil(t, data)
	assert.Nil(t, dataRefreshToken)
	assert.Error(t, err)
}

func TestGetAllProcessorShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	mock.EXPECT().getAllUser().Return(nil, errors.New("error query"))

	data, err := proc.getAllUserProcessor()
	assert.Empty(t, data)
	assert.Error(t, err)
}

func TestGetAllProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	var users []user
	users = append(users, user{ID: 1})
	mock.EXPECT().getAllUser().Return(users, nil)

	data, err := proc.getAllUserProcessor()
	assert.NotEmpty(t, data)
	assert.NoError(t, err)
}

func TestGetUserDataProcessorShouldReturnErrorWhenRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	mock.EXPECT().getUserData("test").Return(nil, errors.New("error"))

	data, err := proc.getUserDataProcessor("test")
	assert.Nil(t, data)
	assert.Error(t, err)
}

func TestGetUserDataProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockuserRepositoryInterface(ctrl)

	proc := userProcessor{repo: mock}
	var users verifiedUser
	users.ID = 1
	usersJSON, _ := json.Marshal(users)
	mock.EXPECT().getUserData("test").Return(usersJSON, nil)

	data, err := proc.getUserDataProcessor("test")
	assert.NotNil(t, data)
	assert.NoError(t, err)
}
