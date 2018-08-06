package services

import (
	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
)

const (
	getAllServicesQuery       = "SELECT id,name,created_at FROM services"
	createServiceQuery        = "INSERT INTO services(name,authorization_token) VALUES ($1,$2)"
	deleteServiceQuery        = "DELETE FROM services WHERE name = $1"
	findOneServiceByNameQuery = "SELECT id, name, created_at FROM services WHERE name = $1"
	getServiceTokenQuery      = "SELECT authorization_token FROM services WHERE name = $1"
)

func getAll(q helper.QueryExecuter) ([]map[string]interface{}, error) {
	data, err := q.DoQuery(getAllServicesQuery)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(createServiceQuery, servicestore.Name, servicestore.AuthorizationToken)
	return err
}

func deleteService(q helper.QueryExecuter, servicestore service) error {
	_, err := q.DoQuery(deleteServiceQuery, servicestore.Name)
	return err
}

func findOneServiceByName(q helper.QueryExecuter, name string) (service, error) {
	data, err := q.DoQuery(findOneServiceByNameQuery, name)
	if err != nil {
		return service{}, err
	}
	var dest service
	err = helper.ParseMaps(data[0], dest)
	return dest, err
}

func FindOneServiceByName(name string) (service, error) {
	return findOneServiceByName(context.GetHelper(), name)
}

func getServiceToken(q helper.QueryExecuter, name string) (string, error) {
	var token string
	data, err := q.DoQuery(getServiceTokenQuery, name)
	if err != nil {
		return "", err
	}
	helper.ParseMap(data[0]["token"], &token)
	return token, nil
}
