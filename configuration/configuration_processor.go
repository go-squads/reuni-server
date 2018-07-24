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

func getLatestVersionProcess(serviceName, namespace string) (int, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return 0, err
	}
	version, err := getLatestVersionForNamespace(service.Id, namespace)
	if err != nil {
		return 0, err
	}
	return version, nil
}

func createNewVersionProcess(serviceName, namespace string, config configView) error {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return err
	}
	createNewVersion(service.Id, namespace, config)
	return nil
}
