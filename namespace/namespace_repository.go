package namespace

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

const createNewNamespaceQuery = "INSERT INTO namespaces(organization_id, service_name, namespace,created_by) VALUES ($1,$2,$3,$4)"
const createNewConfigurationsQuery = "INSERT INTO configurations(organization_id, service_name, namespace, config_store, created_by) VALUES ($1,$2,$3,$4,$5)"
const retrieveAllNamespaceQuery = "SELECT organization_id,service_name,namespace,active_version as version,created_at,updated_at,created_by FROM namespaces WHERE organization_id =$1 AND service_name = $2"
const countNamespaceNameByService = "SELECT count(namespace) as count FROM namespaces WHERE organization_id =$1 AND service_name = $2 AND namespace=$3"
const translateNametoIdQuery = "SELECT id FROM organization WHERE name=$1"

type namespaceRepositoryInterface interface {
	isNamespaceExist(organizationId int, serviceName, namespace string) (bool, error)
	createConfiguration(organizationId int, serviceName, name string, configurations map[string]interface{}, createdBy string) error
	createNewNamespace(namespaceStore *namespaceStore) error
	retrieveAllNamespace(organizationId int, serviceName string) ([]namespaceStore, error)
	getOrganizationId(organizationName string) (int, error)
}

type namespaceRepository struct {
	execer helper.QueryExecuter
}

func initRepository(execer helper.QueryExecuter) *namespaceRepository {
	return &namespaceRepository{
		execer: execer,
	}
}

func (s *namespaceRepository) isNamespaceExist(organizationId int, serviceName, namespace string) (bool, error) {
	data, err := s.execer.DoQueryRow(countNamespaceNameByService, organizationId, serviceName, namespace)
	if err != nil {
		return false, err
	}

	count, ok := data["count"].(int64)
	if !ok {
		return false, errors.New("Query should result int")
	}

	return count > 0, nil
}
func (s *namespaceRepository) createConfiguration(organizationId int, serviceName, name string, configurations map[string]interface{}, createdBy string) error {
	configjson, err := json.Marshal(configurations)
	if err != nil {
		return err
	}
	_, err = s.execer.DoQuery(createNewConfigurationsQuery, organizationId, serviceName, name, configjson, createdBy)
	return err
}
func (s *namespaceRepository) createNewNamespace(namespaceStore *namespaceStore) error {
	if namespaceStore.OrganizationId == 0 || namespaceStore.ServiceName == "" || namespaceStore.Namespace == "" {
		return errors.New(fmt.Sprintf("Data not defined properly (organization_id, service_name, namespace): %v %v %v", namespaceStore.OrganizationId, namespaceStore.ServiceName, namespaceStore.Namespace))
	}
	_, err := s.execer.DoQuery(createNewNamespaceQuery, namespaceStore.OrganizationId, namespaceStore.ServiceName, namespaceStore.Namespace, namespaceStore.CreatedBy)
	return err
}

func (s *namespaceRepository) retrieveAllNamespace(organizationId int, serviceName string) ([]namespaceStore, error) {
	data, err := s.execer.DoQuery(retrieveAllNamespaceQuery, organizationId, serviceName)
	if err != nil {
		return nil, err
	}

	var namespaces []namespaceStore
	err = helper.ParseMap(data, &namespaces)
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}
func (s *namespaceRepository) getOrganizationId(organizationName string) (int, error) {
	ret, err := s.execer.DoQueryRow(translateNametoIdQuery, organizationName)
	if err != nil {
		return 0, err
	}
	res, ok := ret["id"].(int64)
	if !ok {
		return 0, helper.NewHttpError(http.StatusNotFound, "Not Found")
	}
	return int(res), nil
}
