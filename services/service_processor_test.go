package services

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/helper"
	"github.com/stretchr/testify/assert"

	context "github.com/go-squads/reuni-server/appcontext"
)

func TestTokenRandomizerDifferentAtLeastAHundredThousandTry(t *testing.T) {
	var data map[string]bool
	data = make(map[string]bool)
	for i := 0; i < 100000; i++ {
		token := generateTokenProcessor()
		if data[token] {
			t.Fail()
		} else {
			data[token] = true
		}
	}
}

func TestCreateServiceProcessorShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{Name: "Hello-World"}
	err := createServiceProcessor(serv)
	assert.NoError(t, err)
}

func TestCreateServiceProcessorShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("Duplicate key"),
	}
	context.InitMockContext(q)
	serv := servicev{Name: "Hello-World"}
	err := createServiceProcessor(serv)
	assert.Error(t, err)
}

func TestCreateServiceProcessorShouldReturnErrorWhenServiceNameIsEmpty(t *testing.T) {

	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{}
	err := createServiceProcessor(serv)
	assert.Error(t, err)
	httpErr, ok := err.(*helper.HttpError)
	assert.True(t, ok)
	assert.Equal(t, 400, httpErr.Status)
}

func TestCreateServiceProcessorShouldReturnErrorWhenServiceNameIsDuplicate(t *testing.T) {

	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test")},
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{Name: "Hello World"}
	err := createServiceProcessor(serv)
	assert.Error(t, err)
	httpErr, ok := err.(*helper.HttpError)
	assert.True(t, ok)
	assert.Equal(t, 409, httpErr.Status)
}

func TestDeleteServiceProcessorShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test")},
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{Name: "Hello World"}
	err := deleteServiceProcessor(serv)
	assert.NoError(t, err)
}
func TestDeleteServiceProcessorShouldReturnErrorWhenServiceNameIsEmpty(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test")},
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{}
	err := deleteServiceProcessor(serv)
	assert.Error(t, err)
}

func TestValidateTokenProcessorShouldNotReturnError(t *testing.T) {
	data := make(map[string]interface{})
	data["authorization_token"] = "123"
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{data},
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(serv.Name, "123")
	assert.True(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorShouldNotReturnErrorWhenTokenIsNotValid(t *testing.T) {
	data := make(map[string]interface{})
	data["authorization_token"] = "123"
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{data},
		Err:  nil,
	}
	context.InitMockContext(q)
	serv := servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(serv.Name, "155")
	assert.False(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorShouldReturnErrorWhenTokenIsNil(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{nil},
		Err:  errors.New("Test Error"),
	}
	context.InitMockContext(q)
	serv := servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(serv.Name, "155")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestFindOneServiceByNameWithContextShouldReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("Test Error"),
	}
	context.InitMockContext(q)
	service, err := FindOneServiceByName("test")
	assert.Nil(t, service)
	assert.Error(t, err)
}

func TestFindOneServiceByNameWithContextShouldNotReturnError(t *testing.T) {

	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test-service")},
		Err:  nil,
	}
	context.InitMockContext(q)
	service, err := FindOneServiceByName("test")
	assert.Equal(t, service.Name, "test-service")
	assert.Equal(t, service.Id, 1)
	assert.NoError(t, err)
}
