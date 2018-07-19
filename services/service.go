package services

import (
	"log"

	context "github.com/go-squads/reuni-server/appcontext"
)

type Service struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const getAllQuery = "SELECT id,name FROM services"

func GetAll() []Service {
	var services []Service

	db := context.GetDB()
	rows, err := db.Query(getAllQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var service Service
		err := rows.Scan(&service.Id, &service.Name)

		if err != nil {
			log.Fatal(err)
		}
		services = append(services, service)
	}
	log.Printf("%v", services)
	return services
}
