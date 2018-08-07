package namespace

import (
	"encoding/json"
	"errors"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/services"
)

var activeRepo namespaceRepositoryInterface

func getActiveRepo() namespaceRepositoryInterface {
	if activeRepo == nil {
		initRepository(appcontext.GetHelper())
	}
	return activeRepo
}

func createNewNamespaceProcessor(serviceName string, namespacev namespaceView) error {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return err
	}
	var namespaceStore = namespaceStore{}
	namespaceStore.ServiceId = service.Id
	namespaceStore.Namespace = namespacev.Namespace
	namespaceStore.ActiveVersion = 1
	configurations := namespacev.Configuration
	isNamespaceExist, err := getActiveRepo().isNamespaceExist(namespaceStore.ServiceId, namespaceStore.Namespace)
	if err != nil {
		return err
	}

	if isNamespaceExist {
		return errors.New("Namespace already exist for the service")
	}

	err = getActiveRepo().createNewNamespace(&namespaceStore)
	if err != nil {
		return err
	}
	err = getActiveRepo().createConfiguration(service.Id, namespacev.Namespace, configurations)

	return err
}

func retrieveAllNamespaceProcess(serviceName string) ([]byte, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return nil, err
	}

	namespaces, err := getActiveRepo().retrieveAllNamespace(service.Id)
	if err != nil {
		return nil, err
	}
	namespaceJSON, err := json.Marshal(namespaces)
	if err != nil {
		return nil, err
	}
	return namespaceJSON, nil
}
