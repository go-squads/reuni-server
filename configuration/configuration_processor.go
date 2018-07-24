package configuration

import "github.com/go-squads/reuni-server/services"

func getConfigurationProcess(serviceName, namespace string, version int) (*configView, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return nil, err
	}
	config, err := getConfiguration(service.Id, namespace, version)
	if err != nil {
		return nil, err
	}
	return config, nil
}
