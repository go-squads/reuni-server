package services

import (
	"errors"
	"testing"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service_expected := service{
		Name:               "test",
		OrganizationId:     1,
		AuthorizationToken: "123",
	}
	mock.EXPECT().findOneServiceByName(1, "test").Return(nil, errors.New("error"))
	mock.EXPECT().generateToken().Return("123")
	mock.EXPECT().createService(service_expected).Return(nil)

	serv := servicev{Name: "test"}
	err := proc.createServiceProcessor(serv, 1)
	assert.NoError(t, err)
}

func TestCreateServiceProcessorShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service_expected := service{
		Name:               "test",
		OrganizationId:     1,
		AuthorizationToken: "123",
	}
	mock.EXPECT().findOneServiceByName(1, "test").Return(nil, errors.New("error"))
	mock.EXPECT().generateToken().Return("123")
	mock.EXPECT().createService(service_expected).Return(errors.New("Duplicate key"))

	serv := servicev{Name: "test"}
	err := proc.createServiceProcessor(serv, 1)
	assert.Error(t, err)
}

func TestCreateServiceProcessorShouldReturnErrorWhenServiceNameIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service_expected := service{
		Name:               "test",
		OrganizationId:     1,
		AuthorizationToken: "123",
	}
	mock.EXPECT().findOneServiceByName(1, "test").Return(nil, errors.New("error"))
	mock.EXPECT().generateToken().Return("123")
	mock.EXPECT().createService(service_expected).Return(nil)

	serv := servicev{}
	err := proc.createServiceProcessor(serv, 1)
	assert.Error(t, err)
	httpErr, ok := err.(*helper.HttpError)
	assert.True(t, ok)
	assert.Equal(t, 400, httpErr.Status)
}

func TestCreateServiceProcessorShouldReturnErrorWhenServiceNameIsDuplicate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service_expected := service{
		Name:               "test",
		OrganizationId:     1,
		AuthorizationToken: "123",
	}
	mock.EXPECT().findOneServiceByName(1, "test").Return(&service_expected, nil)
	mock.EXPECT().generateToken().Return("123")
	mock.EXPECT().createService(service_expected).Return(nil)

	serv := servicev{Name: "test"}
	err := proc.createServiceProcessor(serv, 1)
	assert.Error(t, err)
	httpErr, ok := err.(*helper.HttpError)
	assert.True(t, ok)
	assert.Equal(t, 409, httpErr.Status)
}

func TestDeleteServiceProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service := service{
		OrganizationId: 1,
		Name:           "test",
	}
	mock.EXPECT().deleteService(service).Return(nil)
	serv := servicev{Name: "test"}
	err := proc.deleteServiceProcessor(1, serv)
	assert.NoError(t, err)
}

func TestDeleteServiceProcessorShouldReturnErrorWhenServiceNameIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	service := service{}
	mock.EXPECT().deleteService(service).Return(nil)

	serv := servicev{}
	err := proc.deleteServiceProcessor(1, serv)
	assert.Error(t, err)
}

func TestValidateTokenProcessorShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	serv := &servicev{
		Name: "test",
	}
	serviceToken_expected := &serviceToken{
		Token: "123",
	}
	mock.EXPECT().getServiceToken(1, "test").Return(serviceToken_expected, nil)
	res, err := proc.ValidateTokenProcessor(1, serv.Name, "123")
	assert.True(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorShouldNotReturnErrorWhenTokenIsNotValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	serv := &servicev{
		Name: "test",
	}
	serviceToken_expected := &serviceToken{
		Token: "123",
	}
	mock.EXPECT().getServiceToken(1, "test").Return(serviceToken_expected, nil)

	res, err := proc.ValidateTokenProcessor(1, serv.Name, "155")
	assert.False(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorShouldReturnErrorWhenTokenIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}
	serv := &servicev{
		Name: "test",
	}
	serviceToken_expected := &serviceToken{
		Token: "",
	}
	mock.EXPECT().getServiceToken(1, "test").Return(serviceToken_expected, errors.New("token is nil"))
	res, err := proc.ValidateTokenProcessor(1, serv.Name, "")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestValidateTokenProcessorWithoutReceiverShouldNotReturnError(t *testing.T) {
	data := make(map[string]interface{})
	data["authorization_token"] = "123"
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{data},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	serv := &servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(q, 1, serv.Name, "123")
	assert.True(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorWithoutReceiverShouldNotReturnErrorWhenTokenIsNotValid(t *testing.T) {
	data := make(map[string]interface{})
	data["authorization_token"] = "123"
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{data},
		Err:  nil,
	}
	appcontext.InitMockContext(q)
	serv := &servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(q, 1, serv.Name, "155")
	assert.False(t, res)
	assert.NoError(t, err)
}

func TestValidateTokenProcessorWithoutProcessorShouldReturnErrorWhenTokenIsNil(t *testing.T) {
	data := make(map[string]interface{})
	data["authorization_token"] = nil
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{data},
		Err:  errors.New("token is nil"),
	}
	appcontext.InitMockContext(q)
	serv := &servicev{
		Name: "test",
	}
	res, err := ValidateTokenProcessor(q, 1, serv.Name, "")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestFindOneServiceByNameWithContextShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().findOneServiceByName(1, "test").Return(nil, errors.New("error data"))

	service, err := proc.FindOneServiceByName(1, "test")
	assert.Nil(t, service)
	assert.Error(t, err)
}

func TestFindOneServiceByNameWithContextShouldNotReturnError(t *testing.T) {

	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().findOneServiceByName(1, "test").Return(&service{Name: "test"}, nil)

	service, err := proc.FindOneServiceByName(1, "test")
	assert.Equal(t, service.Name, "test")
	assert.NoError(t, err)
}

func TestFindOneServiceByNameWithoutReceiverWithContextShouldReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("Test Error"),
	}
	appcontext.InitMockContext(q)

	service, err := FindOneServiceByName(q, 1, "test")
	assert.Nil(t, service)
	assert.Error(t, err)
}

func TestFindOneServiceByNameWithoutReceiverWithContextShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: []map[string]interface{}{MockServiceMap(1, "test")},
		Err:  nil,
	}
	appcontext.InitMockContext(q)

	service, err := FindOneServiceByName(q, 1, "test")
	assert.Equal(t, service.Name, "test")
	assert.NoError(t, err)
}

func TestTranslateNameToIdWithContextShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().translateNameToIdRepository("test").Return(0, errors.New("error data"))

	id, err := proc.TranslateNameToIdProcessor("test")
	assert.Equal(t, 0, id)
	assert.Error(t, err)
}

func TestTranslateNameToIdWithContextShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().translateNameToIdRepository("test").Return(1, nil)

	id, err := proc.TranslateNameToIdProcessor("test")
	assert.Equal(t, id, 1)
	assert.NoError(t, err)
}

func TestTranslateNameToIdWithoutReceiverWithContextShouldReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Data: nil,
		Err:  errors.New("Test Error"),
	}
	appcontext.InitMockContext(q)

	id, err := TranslateNameToIdProcessor(q, "test")
	assert.Equal(t, 0, id)
	assert.Error(t, err)
}

func TestTranslateNameToIdWithoutReceiverWithContextShouldNotReturnError(t *testing.T) {
	q := &helper.QueryMockHelper{
		Row: map[string]interface{}{
			"id": int64(1),
		},
		Err: nil,
	}
	appcontext.InitMockContext(q)

	id, err := TranslateNameToIdProcessor(q, "test")
	assert.Equal(t, id, 1)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldNotReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().getAll(1).Return([]service{}, nil)

	services, err := proc.getAllServicesBasedOnOrganizationProcessor(1)
	assert.NotNil(t, services)
	assert.NoError(t, err)
}

func TestGetAllServiceShouldReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockserviceRepositoryInterface(ctrl)
	proc := serviceProcessor{repo: mock}

	mock.EXPECT().getAll(1).Return(nil, errors.New("error"))

	services, err := proc.getAllServicesBasedOnOrganizationProcessor(1)
	assert.Nil(t, services)
	assert.Error(t, err)
}
