package authorizator_test

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/authorizator"
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
func TestAuthorizeShouldReturnTrueWhenAdminAccessReadPrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Admin"}, nil))
	res := auth.Authorize(1, 1, 'r')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnTrueWhenAdminAccessWritePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Admin"}, nil))
	res := auth.Authorize(1, 1, 'w')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnTrueWhenAdminAccessCreatePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Admin"}, nil))
	res := auth.Authorize(1, 1, 'c')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnTrueWhenAdminAccessDeleterivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Admin"}, nil))
	res := auth.Authorize(1, 1, 'd')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnFalseWhenDeveloperAccessWritePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Developer"}, nil))
	res := auth.Authorize(1, 1, 'c')
	assert.False(t, res)
}

func TestAuthorizeShouldReturnFalseWhenDeveloperAccessDeletePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Developer"}, nil))
	res := auth.Authorize(1, 1, 'd')
	assert.False(t, res)
}

func TestAuthorizeShouldReturnTrueWhenDeveloperAccessWritePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Developer"}, nil))
	res := auth.Authorize(1, 1, 'w')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnFalseWhenDeveloperAccessReadPrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Developer"}, nil))
	res := auth.Authorize(1, 1, 'r')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnTrueWhenAuditorAccessReadPrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Auditor"}, nil))
	res := auth.Authorize(1, 1, 'r')
	assert.True(t, res)
}

func TestAuthorizeShouldReturnFalseWhenAuditorAccessWritePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Auditor"}, nil))
	res := auth.Authorize(1, 1, 'w')
	assert.False(t, res)
}

func TestAuthorizeShouldReturnFalseWhenAuditorAccessDeletePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Auditor"}, nil))
	res := auth.Authorize(1, 1, 'd')
	assert.False(t, res)
}
func TestAuthorizeShouldReturnFalseWhenAuditorAccessCreatePrivilege(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Auditor"}, nil))
	res := auth.Authorize(1, 1, 'c')
	assert.False(t, res)
}

func TestAuthorizeShouldReturnFalseWheneverQueryReturnError(t *testing.T) {
	auth := authorizator.New(makeMockRow(map[string]interface{}{"role": "Auditor"}, errors.New("TestError")))
	res := auth.Authorize(1, 1, 'c')
	assert.False(t, res)
}
