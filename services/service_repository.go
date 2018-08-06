package services

import (
	"github.com/go-squads/reuni-server/helper"
)

const (
	getAllServicesQuery       = "SELECT id,name,created_at FROM services"
	createServiceQuery        = "INSERT INTO services(name,authorization_token) VALUES ($1,$2)"
	deleteServiceQuery        = "DELETE FROM services WHERE name = $1"
	findOneServiceByNameQuery = "SELECT id, name, created_at FROM services WHERE name = $1"
	getServiceTokenQuery      = "SELECT authorization_token as token FROM services WHERE name = $1"
)

func getAll(q helper.QueryExecuter) ([]service, error) {
	data, err := q.DoQuery(getAllServicesQuery)
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

func createService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(createServiceQuery, servicestore.Name, servicestore.AuthorizationToken)
	return err
}

func deleteService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(deleteServiceQuery, servicestore.Name)
	return err
}

func findOneServiceByName(q helper.QueryExecuter, name string) (*service, error) {
	data, err := q.DoQuery(findOneServiceByNameQuery, name)
	if err != nil {
		return nil, err
	}
	var dest service
	err = helper.ParseMap(data[0], &dest)
	if err != nil {
		return nil, err
	}
	return &dest, err
}

func getServiceToken(q helper.QueryExecuter, name string) (*serviceToken, error) {
	var token serviceToken
	data, err := q.DoQuery(getServiceTokenQuery, name)
	if err != nil {
		return nil, err
	}
	err = helper.ParseMap(data[0], &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
