package services

import (
	"crypto/rand"
	"encoding/base64"

	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
)

func createServiceProcessor(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	if serviceStore.Name == "" {
		return helper.NewHttpError(400, "Service name not defined")
	}
	_, err := findOneServiceByName(context.GetHelper(), serviceStore.Name)
	if err == nil {
		return helper.NewHttpError(409, "Service already exist")
	}
	serviceStore.AuthorizationToken = generateTokenProcessor()
	return createService(context.GetHelper(), serviceStore)
}

func deleteServiceProcessor(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	if serviceStore.Name == "" {
		return helper.NewHttpError(400, "Service name not defined")
	}
	return deleteService(context.GetHelper(), serviceStore)
}

func generateTokenProcessor() string {
	randomBytes := make([]byte, 64)
	rand.Read(randomBytes)
	return base64.StdEncoding.EncodeToString(randomBytes)[:64]
}

func ValidateTokenProcessor(serviceName string, inputToken string) (bool, error) {
	token, err := getServiceToken(context.GetHelper(), serviceName)
	if err != nil {
		return false, err
	}
	if token.Token == inputToken {
		return true, nil
	} else {
		return false, nil
	}
}

func FindOneServiceByName(name string) (*service, error) {
	return findOneServiceByName(context.GetHelper(), name)
}
