package namespace

import (
	"encoding/json"
	"errors"

	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/appcontext"
)

var activeRepo namespaceRepositoryInterface

func getActiveRepo() namespaceRepositoryInterface {
	if activeRepo == nil {
		initRepository(appcontext.GetHelper())
	}
	return activeRepo
}

func parseData(serviceId int, view *namespaceView, data *namespaceStore) error {
	data.ServiceId = serviceId
	data.Namespace = view.Namespace
	data.ActiveVersion = 1
	if data.Namespace == "" {
		return helper.NewHttpError(400, "Namespace cannot be empty")
	}
	return nil
}

func createNewNamespaceProcessor(serviceName string, namespacev *namespaceView) error {
	serviceId, err := getActiveRepo().getServiceId(serviceName)
	if err != nil {
		return err
	}
	var namespaceStore namespaceStore
	err = parseData(serviceId, namespacev, &namespaceStore)
	if err != nil {
		return err
	}

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
	configurations := namespacev.Configuration
	err = getActiveRepo().createConfiguration(serviceId, namespacev.Namespace, configurations)

	return err
}

func retrieveAllNamespaceProcessor(serviceName string) ([]byte, error) {
	serviceId, err := getActiveRepo().getServiceId(serviceName)
	if err != nil {
		return nil, err
	}

	namespaces, err := getActiveRepo().retrieveAllNamespace(serviceId)
	if err != nil {
		return nil, err
	}
	namespaceJSON, err := json.Marshal(namespaces)
	return namespaceJSON, err
}
