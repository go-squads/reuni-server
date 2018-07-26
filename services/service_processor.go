package services

import (
	"crypto/rand"
	"encoding/base64"
)

func createServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	serviceStore.AuthorizationToken = generateToken()
	return createService(serviceStore)
}

func deleteServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return deleteService(serviceStore)
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
	token, err := getServiceToken(serviceName)
	if err != nil {
		return false, err
	}
	if token == inputToken {
		return true, nil
	} else {
		return false, nil
	}
}
