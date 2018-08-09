package namespace

import (
	"encoding/json"
	"errors"

	"github.com/go-squads/reuni-server/helper"
)

type processor interface {
	parseData(serviceId int, view *namespaceView, data *namespaceStore) error
	createNewNamespaceProcessor(serviceName string, namespacev *namespaceView) error
	retrieveAllNamespaceProcessor(serviceName string) ([]byte, error)
}

type mainProcessor struct {
	repo namespaceRepositoryInterface
}

func (s *mainProcessor) parseData(serviceId int, view *namespaceView, data *namespaceStore) error {
	data.ServiceId = serviceId
	data.Namespace = view.Namespace
	data.ActiveVersion = 1
	if data.Namespace == "" {
		return helper.NewHttpError(400, "Namespace cannot be empty")
	}
	return nil
}

func (s *mainProcessor) createNewNamespaceProcessor(serviceName string, namespacev *namespaceView) error {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return err
	}
	var namespaceStore namespaceStore
	err = s.parseData(serviceId, namespacev, &namespaceStore)
	if err != nil {
		return err
	}
	isNamespaceExist, err := s.repo.isNamespaceExist(namespaceStore.ServiceId, namespaceStore.Namespace)
	if err != nil {
		return err
	}

	if isNamespaceExist {
		return errors.New("Namespace already exist for the service")
	}
	err = s.repo.createNewNamespace(&namespaceStore)
	if err != nil {
		return err
	}
	configurations := namespacev.Configuration
	err = s.repo.createConfiguration(serviceId, namespacev.Namespace, configurations)
	return err
}

func (s *mainProcessor) retrieveAllNamespaceProcessor(serviceName string) ([]byte, error) {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return nil, err
	}

	namespaces, err := s.repo.retrieveAllNamespace(serviceId)
	if err != nil {
		return nil, err
	}
	if namespaces == nil {
		return []byte("[]"), nil
	}
	namespaceJSON, err := json.Marshal(namespaces)
	return namespaceJSON, err
}
