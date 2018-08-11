package services

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

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

func MockServiceMap(id int, name string) map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = id
	m["name"] = name
	return m
}

func MockServiceStruct(id int, name string) service {
	return service{
		Id:   id,
		Name: name,
	}
}
func TestGetAllServiceShouldNotReturnErrorWhenQueryReturnOneData(t *testing.T) {
	var data []map[string]interface{}
	data = append(data, MockServiceMap(1, "go-pay-service"))

	rep := initRepository(makeMockRows(data, nil))
	services, err := rep.getAll(1)
	var expected []service

	expected = append(expected, MockServiceStruct(1, "go-pay-service"))
	assert.Equal(t, services, expected)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenDataMoreThanOne(t *testing.T) {

	var data []map[string]interface{}
	data = append(data, MockServiceMap(1, "go-pay-service"))
	data = append(data, MockServiceMap(2, "go-ride-service"))

	rep := initRepository(makeMockRows(data, nil))
	services, err := rep.getAll(1)
	var expected []service
	expected = append(expected, MockServiceStruct(1, "go-pay-service"))
	expected = append(expected, MockServiceStruct(2, "go-ride-service"))
	assert.Equal(t, services, expected)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenQueryDoesNotReturnData(t *testing.T) {
	var data []map[string]interface{}
	rep := initRepository(makeMockRows(data, nil))

	services, err := rep.getAll(1)
	assert.Empty(t, services)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldReturnErrorWhenQueryReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("Query Return Error")))

	services, err := rep.getAll(1)

	assert.Nil(t, services)
	assert.Error(t, err)
}

func TestGetAllServiceShouldNotReturnErrorWhenDataNotParseableToStruct(t *testing.T) {
	var data []map[string]interface{}
	row := make(map[string]interface{})
	row["test"] = 1
	data = append(data, row)
	rep := initRepository(makeMockRows(data, nil))

	var expected []service
	expected = append(expected, MockServiceStruct(0, ""))
	services, err := rep.getAll(1)
	assert.Equal(t, expected, services)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldReturnErrorWhenDataNotMarshalable(t *testing.T) {
	var data []map[string]interface{}
	row := make(map[string]interface{})
	row["test"] = make(chan int)
	data = append(data, row)
	rep := initRepository(makeMockRows(data, nil))

	services, err := rep.getAll(1)
	assert.Nil(t, services)
	assert.Error(t, err)
}

func TestCreateServiceShouldNotReturnErrorWhenQueryNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))

	err := rep.createService(service{Name: "Hello", AuthorizationToken: "World"})
	assert.NoError(t, err)
}

func TestCreateServiceShouldReturnErrorWhenQueryReturnError(t *testing.T) {

	rep := initRepository(makeMockRows(nil, errors.New("This is Error")))

	err := rep.createService(service{Name: "Hello", AuthorizationToken: "World"})
	assert.Error(t, err)
}

func TestDeleteServiceShouldNotReturnErrorWhenQueryNotReturnError(t *testing.T) {
	rep := initRepository(makeMockRows(nil, nil))

	err := rep.deleteService(service{Name: "Hello"})
	assert.NoError(t, err)
}

func TestFindOneServiceByNameShouldNotReturnError(t *testing.T) {

	var datum = make(map[string]interface{})
	datum["id"] = 7
	datum["name"] = "New_Service"
	datum["authorization_token"] = ""
	datum["created_at"] = nil
	data := []map[string]interface{}{datum}
	rep := initRepository(makeMockRows(data, nil))
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)

	mock.EXPECT().findOneServiceByName("Hello").Return(&service{Name: "Hello"}, nil)
	s, err := rep.findOneServiceByName("Hello")
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestFindOneServiceByNameShouldReturnError(t *testing.T) {
	rep := initRepository(makeMockRow(nil, errors.New("data error")))

	s, err := rep.findOneServiceByName("Hello")
	assert.Error(t, err)
	assert.Empty(t, s)
}

func TestFindOneServiceByNameShouldReturnErrorWhenDataIsEmpty(t *testing.T) {
	rep := initRepository(makeMockRows([]map[string]interface{}{}, nil))

	s, err := rep.findOneServiceByName("Hello")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestFindOneServiceByNameShouldReturnErrorIfDataNotParsable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["name"] = errors.New("error")
	data := []map[string]interface{}{datum}

	rep := initRepository(makeMockRows(data, nil))

	s, err := rep.findOneServiceByName("Hello")
	assert.Error(t, err)
	assert.Nil(t, s)
}
func TestFindOneServiceByNameShouldReturnErrorIfDataNotMarshalable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["test"] = make(chan int)
	data := []map[string]interface{}{datum}

	rep := initRepository(makeMockRows(data, errors.New("data error")))

	s, err := rep.findOneServiceByName("Hello")
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGetTokenShouldNotReturnError(t *testing.T) {
	var datum = make(map[string]interface{})
	datum["authorization_token"] = "asdfsdfa"
	var data []map[string]interface{}
	data = append(data, datum)
	rep := initRepository(makeMockRows(data, nil))

	token, err := rep.getServiceToken("Hello")
	assert.NoError(t, err)
	assert.Equal(t, "asdfsdfa", token.Token)
}

func TestGetTokenShouldReturnErrorWhenThereIsNodata(t *testing.T) {
	rep := initRepository(makeMockRows(nil, errors.New("error no data")))

	services, err := rep.getServiceToken("hello")
	assert.Empty(t, services)
	assert.Error(t, err)
}

func TestGetTokenShouldReturnErrorWhenDataNotParsable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["authorization_token"] = errors.New("err")
	data := []map[string]interface{}{datum}
	rep := initRepository(makeMockRows(data, nil))

	services, err := rep.getServiceToken("hello")
	assert.Empty(t, services)
	assert.Error(t, err)
}

func TestGetTokenShouldReturnErrorWhenDataNotMarshalable(t *testing.T) {
	datum := make(map[string]interface{})
	datum["test"] = make(chan int)
	var data []map[string]interface{}
	rep := initRepository(makeMockRows(data, nil))

	services, err := rep.getServiceToken("hello")
	assert.Empty(t, services)
	assert.Error(t, err)
}
func TestTranslateNameToIdRepositoryShouldReturnErrorOnQuery(t *testing.T) {
	mock := makeMockRow(nil, errors.New("error data"))
	rep := initRepository(mock)

	id, err := rep.translateNameToIdRepository("test")
	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestTranslateNameToIdRepositoryShouldNotReturnError(t *testing.T) {
	datum := make(map[string]interface{})
	datum["id"] = int64(1)
	mock := makeMockRow(datum, nil)
	rep := initRepository(mock)

	id, err := rep.translateNameToIdRepository("test")
	assert.NotEmpty(t, id)
	assert.NoError(t, err)
}

func TestTokenRandomizerDifferentAtLeastAHundredThousandTry(t *testing.T) {
	var data map[string]bool
	data = make(map[string]bool)
	mock := makeMockRow(nil, nil)
	rep := initRepository(mock)
	for i := 0; i < 100000; i++ {
		token := rep.generateToken()
		if data[token] {
			t.Fail()
		} else {
			data[token] = true
		}
	}
}
