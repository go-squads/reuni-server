package namespace

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/go-squads/reuni-server/helper"
)

const createNewNamespaceQuery = "INSERT INTO namespaces(service_id, namespace) VALUES ($1,$2)"
const createNewConfigurationsQuery = "INSERT INTO configurations(service_id, namespace, config_store) VALUES ($1,$2,$3)"
const retrieveAllNamespaceQuery = "SELECT namespace,active_version namespace FROM namespaces WHERE service_id = $1"
const countNamespaceNameByService = "SELECT count(namespace) as count FROM namespaces WHERE service_id=$1 AND namespace=$2"

func isNamespaceExist(q helper.QueryExecuter, service_id int, namespace string) (bool, error) {
	data, err := q.DoQueryRow(countNamespaceNameByService, service_id, namespace)
	if err != nil {
		return false, err
	}
	count, ok := data["count"].(int)
	if !ok {
		return false, errors.New("Query should result int")
	}

	return count > 0, nil
}
func createConfiguration(q helper.QueryExecuter, serviceId int, name string, configurations map[string]interface{}) error {
	configjson, err := json.Marshal(configurations)
	if err != nil {
		return err
	}
	_, err = q.DoQuery(createNewConfigurationsQuery, serviceId, name, configjson)
	return err
}
func createNewNamespace(q helper.QueryExecuter, namespaceStore namespaceStore) error {
	isNamespaceExist, err := isNamespaceExist(q, namespaceStore.ServiceId, namespaceStore.Namespace)
	if err != nil {
		return err
	}

	if isNamespaceExist {
		return errors.New("Namespace already exist for the service")
	}
	_, err = q.DoQuery(createNewNamespaceQuery, namespaceStore.ServiceId, namespaceStore.Namespace)
	return err
}

func retrieveAllNamespace(q helper.QueryExecuter, service_id int) ([]namespaceStore, error) {
	data, err := q.DoQuery(retrieveAllNamespaceQuery, service_id)
	if err != nil {
		return nil, err
	}
	var configurations []namespaceStore
	log.Print(data)
	err = helper.ParseMap(data, &configurations)
	return configurations, nil
}
