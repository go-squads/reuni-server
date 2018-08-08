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
	rep := initRepository(makeMockRows(nil, nil))
	err := rep.createNewOrganization("test-org")
	assert.NoError(t, err)
}
func TestCreateOrganizationShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Test Error")))
	err := rep.createNewOrganization("test-org")
	assert.Error(t, err)
}
