package organization

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
func makeRows(args ...map[string]interface{}) []map[string]interface{} {
	return args
}

func TestCreateOrganizationShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": int64(1)}, nil))
	id, err := rep.createNewOrganization("test-org")
	assert.Equal(t, int64(1), id)
	assert.NoError(t, err)
}
func TestCreateOrganizationShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("Test Error")))
	id, err := rep.createNewOrganization("test-org")
	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestCreateOrganizationShouldReturnErrorWhenIdCannotBeParsed(t *testing.T) {
	rep := initRepository(makeMockRow(map[string]interface{}{"id": "random"}, nil))
	id, err := rep.createNewOrganization("test-org")
	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestAddUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Test Error")))
	err := rep.addUser(int64(1), int64(1), "adsd")
	assert.Error(t, err)
}

func TestAddUserShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))
	err := rep.addUser(int64(1), int64(1), "Admin")
	assert.NoError(t, err)
}

func TestDeleteUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Test Error")))
	err := rep.deleteUser(int64(1), int64(1))
	assert.Error(t, err)
}

func TestDeleteUserShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))
	err := rep.deleteUser(int64(1), int64(1))
	assert.NoError(t, err)
}

func TestUpdateRoleOfUserShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Test Error")))
	err := rep.updateRoleOfUser("dmn", int64(1), int64(1))
	assert.Error(t, err)
}

func TestUpdateRoleOfUserShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))
	err := rep.updateRoleOfUser("Admin", int64(1), int64(1))
	assert.NoError(t, err)
}

func TestGetAllMemberOfOrganizationShouldNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows([]map[string]interface{}{}, nil))
	data, err := rep.getAllMemberOfOrganization(int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestGetAllMemberOfOrganizationShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Internal error")))
	data, err := rep.getAllMemberOfOrganization(int64(1))
	assert.Error(t, err)
	assert.Nil(t, data)
}
