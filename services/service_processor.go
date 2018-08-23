package services

import (
	"github.com/go-squads/reuni-server/helper"
)

type serviceProcessorInterface interface {
	createServiceProcessor(servicedata servicev, organizationId int) error
	deleteServiceProcessor(organizationId int, servicedata servicev) error
	ValidateTokenProcessor(organizationId int, serviceName string, inputToken string) (bool, error)
	FindOneServiceByName(organizationId int, name string) (*service, error)
	TranslateNameToIdProcessor(name string) (int, error)
	getAllServicesBasedOnOrganizationProcessor(organizationId int) ([]service, error)
}

type serviceProcessor struct {
	repo serviceRepositoryInterface
}

func (p *serviceProcessor) createServiceProcessor(servicedata servicev, organizationId int) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	serviceStore.OrganizationId = organizationId
	serviceStore.CreatedBy = servicedata.CreatedBy
	if serviceStore.Name == "" {
		return helper.NewHttpError(400, "Service name not defined")
	}
	_, err := p.repo.findOneServiceByName(organizationId, serviceStore.Name)
	if err == nil {
		return helper.NewHttpError(409, "Service already exist")
	}
	serviceStore.AuthorizationToken = p.repo.generateToken()
	return p.repo.createService(serviceStore)
}

func (p *serviceProcessor) deleteServiceProcessor(organizationId int, servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	serviceStore.OrganizationId = organizationId
	if serviceStore.Name == "" {
		return helper.NewHttpError(400, "Service name not defined")
	}
	return p.repo.deleteService(serviceStore)
}

func (p *serviceProcessor) ValidateTokenProcessor(organizationId int, serviceName string, inputToken string) (bool, error) {
	token, err := p.repo.getServiceToken(organizationId, serviceName)
	if err != nil {
		return false, err
	}
	if token.Token == inputToken {
		return true, nil
	} else {
		return false, nil
	}
}

func ValidateTokenProcessor(q helper.QueryExecuter, organizationId int, serviceName string, inputToken string) (bool, error) {
	token, err := initRepository(q).getServiceToken(organizationId, serviceName)
	if err != nil {
		return false, err
	}
	if token.Token == inputToken {
		return true, nil
	} else {
		return false, nil
	}
}

func (p *serviceProcessor) FindOneServiceByName(organizationId int, name string) (*service, error) {
	return p.repo.findOneServiceByName(organizationId, name)
}

func (p *serviceProcessor) TranslateNameToIdProcessor(name string) (int, error) {
	return p.repo.translateNameToIdRepository(name)
}

func FindOneServiceByName(q helper.QueryExecuter, organizationId int, name string) (*service, error) {
	return initRepository(q).findOneServiceByName(organizationId, name)
}

func TranslateNameToIdProcessor(q helper.QueryExecuter, name string) (int, error) {
	return initRepository(q).translateNameToIdRepository(name)
}

func (p *serviceProcessor) getAllServicesBasedOnOrganizationProcessor(organizationId int) ([]service, error) {
	return p.repo.getAll(organizationId)
}
