package configurations

import (
	"encoding/json"

	"github.com/go-squads/reuni-server/services"
)

func createNewNamespaceProcess(service_name string, configurationv configurationView) error {
	service, err := services.FindOneServiceByName(service_name)
	var configStore = configurationStore{}
	configStore.ServiceId = service.Id
	configStore.Namespace = configurationv.Namespace
	configStore.Version = 1
	configStore.Configurations = configurationv.Configuration

	err = createNewNamespace(configStore)
	return err
}

func retrieveAllNamespaceProcess(service_name string) ([]byte, error) {
	service, err := services.FindOneServiceByName(service_name)
	if err != nil {
		return nil, err
	}
	configurations, err := retrieveAllNamespace(service.Id)
	if err != nil {
		return nil, err
	}
	configjson, err := json.Marshal(configurations)
	if err != nil {
		return nil, err
	}
	return configjson, nil
}
