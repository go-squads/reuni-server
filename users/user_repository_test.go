package users

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

func TestCreateUserShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": int64(1)}, nil))
	user := user{
		Name:     "user test",
		Username: "usertest",
		Password: "password",
		Email:    "test@gmail.com",
	}
	err := rep.createUser(user)
	assert.NoError(t, err)
}

func TestCreateUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("error")))
	user := user{
		Name:     "user test",
		Username: "usertest",
		Password: "password",
		Email:    "test@gmail.com",
	}
	err := rep.createUser(user)
	assert.Error(t, err)
}

func TestLoginUserShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": int64(1), "name": "test", "username": "test", "email": "test"}, nil))
	user := userv{
		Username: "usertest",
		Password: "password",
	}
	data, dataRefreshToken, err := rep.loginUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.NotNil(t, dataRefreshToken)
}

func TestLoginUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": 1, "name": "test", "username": "test", "email": "test"}, errors.New("error")))
	user := userv{
		Username: "usertest",
		Password: "password",
	}
	data, dataRefreshToken, err := rep.loginUser(user)
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, dataRefreshToken)
}

func TestLoginUserShouldReturnErrorUnauthorizedWhenNotValid(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": int64(0), "name": "test", "username": "test", "email": "test"}, nil))
	user := userv{
		Username: "usertest",
		Password: "password",
	}
	data, dataRefreshToken, err := rep.loginUser(user)
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, dataRefreshToken)
}

func TestGetAllUserRepositoryShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("error query")))
	data, err := rep.getAllUser()
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestGetAllUserRepositoryShouldReturnErrorWhenFailedToParseData(t *testing.T) {
	datum := map[string]interface{}{"id": errors.New("error")}
	data := []map[string]interface{}{datum}
	rep := initRepository(makeMockRows(data, nil))
	res, err := rep.getAllUser()
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetAllUserRepositoryShouldNotReturnError(t *testing.T) {
	datum := map[string]interface{}{"id": 1}
	data := []map[string]interface{}{datum}
	rep := initRepository(makeMockRows(data, nil))
	res, err := rep.getAllUser()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetUserDataRepositoryShouldReturnErrorWhenQueryError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("error")))
	res, err := rep.getUserData("test")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetUserDataRepositoryShouldNotReturnError(t *testing.T) {
	datum := map[string]interface{}{"id": int64(1), "name": "tester", "username": "test", "email": "test@gmail.com"}
	rep := initRepository(makeMockRow(datum, nil))
	res, err := rep.getUserData("test")
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
