package services

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/go-squads/reuni-server/helper"
)

const (
	getAllServicesQuery       = "SELECT id,name,created_at,created_by FROM services WHERE organization_id = $1"
	createServiceQuery        = "INSERT INTO services(name, organization_id,authorization_token, created_by) VALUES ($1,$2,$3,$4)"
	deleteServiceQuery        = "DELETE FROM services WHERE name = $1"
	findOneServiceByNameQuery = "SELECT id, name, created_by FROM services WHERE name = $1"
	getServiceTokenQuery      = "SELECT authorization_token FROM services WHERE name = $1"
	translateNameToIdQuery    = "SELECT id FROM organization WHERE name = $1"
)

type serviceRepositoryInterface interface {
	getAll(organizationId int) ([]service, error)
	createService(servicestore service) error
	deleteService(servicestore service) error
	getServiceToken(name string) (*serviceToken, error)
	findOneServiceByName(name string) (*service, error)
	translateNameToIdRepository(organizationName string) (int, error)
	generateToken() string
}

type serviceRepository struct {
	execer helper.QueryExecuter
}

func initRepository(execer helper.QueryExecuter) *serviceRepository {
	return &serviceRepository{
		execer: execer,
	}
}

func (s *serviceRepository) getAll(organizationId int) ([]service, error) {
	data, err := s.execer.DoQuery(getAllServicesQuery, organizationId)
	if err != nil {
		return nil, err
	}
	var services []service
	err = helper.ParseMap(data, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *serviceRepository) createService(servicestore service) error {
	_, err := s.execer.DoQuery(createServiceQuery, servicestore.Name, servicestore.OrganizationId, servicestore.AuthorizationToken, servicestore.CreatedBy)
	return err
}

func (s *serviceRepository) deleteService(servicestore service) error {
	_, err := s.execer.DoQuery(deleteServiceQuery, servicestore.Name)
	return err
}

func (s *serviceRepository) findOneServiceByName(name string) (*service, error) {
	data, err := s.execer.DoQuery(findOneServiceByNameQuery, name)
	if err != nil {
		return nil, err
	}
	var dest service
	if len(data) < 1 {
		return nil, helper.NewHttpError(404, "Not Found")
	}
	err = helper.ParseMap(data[0], &dest)
	if err != nil {
		return nil, err
	}
	return &dest, err
}

func (s *serviceRepository) getServiceToken(name string) (*serviceToken, error) {
	var token serviceToken
	data, err := s.execer.DoQuery(getServiceTokenQuery, name)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, helper.NewHttpError(404, "Not Found")
	}
	err = helper.ParseMap(data[0], &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *serviceRepository) translateNameToIdRepository(organizationName string) (int, error) {
	data, err := s.execer.DoQueryRow(translateNameToIdQuery, organizationName)
	if err != nil {
		return 0, err
	}
	id := int(data["id"].(int64))
	return id, nil
}

func (p *serviceRepository) generateToken() string {
	randomBytes := make([]byte, 64)
	rand.Read(randomBytes)
	return base64.StdEncoding.EncodeToString(randomBytes)[:64]
}
