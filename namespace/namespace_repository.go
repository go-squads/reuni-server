package namespace

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

const createNewNamespaceQuery = "INSERT INTO namespaces(service_id, namespace) VALUES ($1,$2)"
const createNewConfigurationsQuery = "INSERT INTO configurations(service_id, namespace, config_store) VALUES ($1,$2,$3)"
const retrieveAllNamespaceQuery = "SELECT id,namespace,active_version as version FROM namespaces WHERE service_id = $1"
const countNamespaceNameByService = "SELECT count(namespace) as count FROM namespaces WHERE service_id=$1 AND namespace=$2"
const findServiceIdFromName = "SELECT id FROM services WHERE name=$1"

type namespaceRepositoryInterface interface {
	isNamespaceExist(service_id int, namespace string) (bool, error)
	createConfiguration(serviceId int, name string, configurations map[string]interface{}) error
	createNewNamespace(namespaceStore *namespaceStore) error
	retrieveAllNamespace(service_id int) ([]namespaceStore, error)
	getServiceId(serviceName string) (int, error)
}

type namespaceRepository struct {
	execer helper.QueryExecuter
}

func initRepository(execer helper.QueryExecuter) *namespaceRepository {
	return &namespaceRepository{
		execer: execer,
	}
}

func (s *namespaceRepository) isNamespaceExist(service_id int, namespace string) (bool, error) {
	data, err := s.execer.DoQueryRow(countNamespaceNameByService, service_id, namespace)
	if err != nil {
		return false, err
	}

	count, ok := data["count"].(int64)
	if !ok {
		return false, errors.New("Query should result int")
	}

	return count > 0, nil
}
func (s *namespaceRepository) createConfiguration(serviceId int, name string, configurations map[string]interface{}) error {
	configjson, err := json.Marshal(configurations)
	if err != nil {
		return err
	}
	_, err = s.execer.DoQuery(createNewConfigurationsQuery, serviceId, name, configjson)
	return err
}
func (s *namespaceRepository) createNewNamespace(namespaceStore *namespaceStore) error {
	if namespaceStore.ServiceId == 0 || namespaceStore.Namespace == "" {
		return errors.New(fmt.Sprintf("Data not defined properly (id, namespace): %v %v", namespaceStore.ServiceId, namespaceStore.Namespace))
	}
	_, err := s.execer.DoQuery(createNewNamespaceQuery, namespaceStore.ServiceId, namespaceStore.Namespace)
	return err
}

func (s *namespaceRepository) retrieveAllNamespace(service_id int) ([]namespaceStore, error) {
	data, err := s.execer.DoQuery(retrieveAllNamespaceQuery, service_id)
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
func (s *namespaceRepository) getServiceId(serviceName string) (int, error) {
	ret, err := s.execer.DoQueryRow(findServiceIdFromName, serviceName)
	if err != nil {
		return 0, err
	}
	res, ok := ret["id"].(int64)
	if !ok {
		return 0, helper.NewHttpError(http.StatusNotFound, "Not Found")
	}
	return int(res), nil
}
