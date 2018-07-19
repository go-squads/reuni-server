package configurations

import (
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
