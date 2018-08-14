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
	data, err := rep.loginUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestLoginUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": 1, "name": "test", "username": "test", "email": "test"}, errors.New("error")))
	user := userv{
		Username: "usertest",
		Password: "password",
	}
	data, err := rep.loginUser(user)
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestLoginUserShouldReturnErrorUnauthorizedWhenNotValid(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": int64(0), "name": "test", "username": "test", "email": "test"}, nil))
	user := userv{
		Username: "usertest",
		Password: "password",
	}
	data, err := rep.loginUser(user)
	assert.Error(t, err)
	assert.Nil(t, data)
}
