package namespace

import (
	"encoding/json"

	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/services"
)

func createNewNamespaceProcess(serviceName string, namespacev namespaceView) error {
	service, err := services.FindOneServiceByName(serviceName)
	var namespaceStore = namespaceStore{}
	namespaceStore.ServiceId = service.Id
	namespaceStore.Namespace = namespacev.Namespace
	namespaceStore.ActiveVersion = 1
	configurations := namespacev.Configuration

	err = createNewNamespace(context.GetHelper(), namespaceStore)
	if err != nil {
		return err
	}
	err = createConfiguration(context.GetHelper(), service.Id, namespacev.Namespace, configurations)

	return err
}

func retrieveAllNamespaceProcess(serviceName string) ([]byte, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return nil, err
	}
	namespaces, err := retrieveAllNamespace(context.GetHelper(), service.Id)
	if err != nil {
		return nil, err
	}
	namespaceJSON, err := json.Marshal(namespaces)
	if err != nil {
		return nil, err
	}
	return namespaceJSON, nil
}
