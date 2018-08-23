package namespace

import (
	"encoding/json"

	"github.com/go-squads/reuni-server/helper"
)

type processor interface {
	parseData(organizationId int, serviceName string, view *namespaceView, data *namespaceStore) error
	createNewNamespaceProcessor(organizationName, serviceName string, namespacev *namespaceView) error
	retrieveAllNamespaceProcessor(organizationName, serviceName string) ([]byte, error)
}

type mainProcessor struct {
	repo namespaceRepositoryInterface
}

func (s *mainProcessor) parseData(organizationId int, serviceName string, view *namespaceView, data *namespaceStore) error {
	data.OrganizationId = organizationId
	data.ServiceName = serviceName
	data.Namespace = view.Namespace
	data.ActiveVersion = 1
	if data.Namespace == "" {
		return helper.NewHttpError(400, "Namespace cannot be empty")
	}
	return nil
}

func (s *mainProcessor) createNewNamespaceProcessor(organizationName, serviceName string, namespacev *namespaceView) error {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return err
	}
	var namespaceStore namespaceStore
	err = s.parseData(organizationId, serviceName, namespacev, &namespaceStore)
	namespaceStore.CreatedBy = namespacev.CreatedBy
	if err != nil {
		return err
	}
	isNamespaceExist, err := s.repo.isNamespaceExist(namespaceStore.OrganizationId, namespaceStore.ServiceName, namespaceStore.Namespace)
	if err != nil {
		return err
	}

	if isNamespaceExist {
		return helper.NewHttpError(409, "Namespace already exist for the service")
	}
	err = s.repo.createNewNamespace(&namespaceStore)
	if err != nil {
		return err
	}
	configurations := namespacev.Configuration
	err = s.repo.createConfiguration(organizationId, serviceName, namespacev.Namespace, configurations)
	return err
}

func (s *mainProcessor) retrieveAllNamespaceProcessor(organizationName, serviceName string) ([]byte, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return nil, err
	}

	namespaces, err := s.repo.retrieveAllNamespace(organizationId, serviceName)
	if err != nil {
		return nil, err
	}
	if namespaces == nil {
		return []byte("[]"), nil
	}
	namespaceJSON, err := json.Marshal(namespaces)
	return namespaceJSON, err
}
