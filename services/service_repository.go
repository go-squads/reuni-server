package services

import (
	"log"

	context "github.com/go-squads/reuni-server/appcontext"
)

const (
	getAllServicesQuery       = "SELECT id,name,created_at FROM services"
	createServiceQuery        = "INSERT INTO services(name,authorization_token) VALUES ($1,$2)"
	deleteServiceQuery        = "DELETE FROM services WHERE name = $1"
	findOneServiceByNameQuery = "SELECT id, name, created_at FROM services WHERE name = $1"
	getServiceTokenQuery      = "SELECT authorization_token FROM services WHERE name = $1"
)

func getAll() ([]service, error) {
	var services []service

	db := context.GetDB()
	rows, err := db.Query(getAllServicesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var service service
		err = rows.Scan(&service.Id, &service.Name, &service.CreatedAt)

		if err != nil {
			log.Fatal(err)
		}
		services = append(services, service)
	}
	log.Printf("%v", services)
	return services, nil
}

func createService(servicestore service) error {
	db := context.GetDB()
	_, err := db.Query(createServiceQuery, servicestore.Name, servicestore.AuthorizationToken)
	return err
}

func deleteService(servicestore service) error {
	db := context.GetDB()
	_, err := db.Query(deleteServiceQuery, servicestore.Name)
	return err
}

func FindOneServiceByName(name string) (service, error) {
	service := service{}
	db := context.GetDB()
	row := db.QueryRow(findOneServiceByNameQuery, name)
	err := row.Scan(&service.Id, &service.Name, &service.CreatedAt)
	return service, err
}

func getServiceToken(name string) (string, error) {
	var token string
	row := context.GetDB().QueryRow(getServiceTokenQuery, name)
	err := row.Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}
