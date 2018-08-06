package services

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/go-squads/reuni-server/helper"

	context "github.com/go-squads/reuni-server/appcontext"
)

func getAllProcess() ([]service, error) {
	data, err := getAll(context.GetHelper())
	if err != nil {
		return nil, err
	}
	var services []service
	err = helper.ParseMaps(data, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func createServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	serviceStore.AuthorizationToken = generateToken()
	return createService(context.GetHelper(), serviceStore)
}

func deleteServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return deleteService(context.GetHelper(), serviceStore)
}

func generateToken() string {
	randomBytes := make([]byte, 64)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(randomBytes)[:64]
}

func ValidateTokenProcess(serviceName string, inputToken string) (bool, error) {
	token, err := getServiceToken(context.GetHelper(), serviceName)
	if err != nil {
		return false, err
	}
	if token == inputToken {
		return true, nil
	} else {
		return false, nil
	}
}
