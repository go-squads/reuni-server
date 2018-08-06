package namespace

import (
	"encoding/json"
	"errors"

	"github.com/go-squads/reuni-server/helper"
)

const createNewNamespaceQuery = "INSERT INTO namespaces(service_id, namespace) VALUES ($1,$2)"
const createNewConfigurationsQuery = "INSERT INTO configurations(service_id, namespace, config_store) VALUES ($1,$2,$3)"
const retrieveAllNamespaceQuery = "SELECT namespace,active_version namespace FROM namespaces WHERE service_id = $1"
const countNamespaceNameByService = "SELECT count(namespace) FROM namespaces WHERE service_id=$1 AND namespace=$2"

func isNamespaceExist(q helper.QueryExecuter, service_id int, namespace string) (bool, error) {
	var count int
	data, err := q.DoQuery(countNamespaceNameByService, service_id, namespace)

	err = helper.ParseMap(data[0]["count"], &count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func createNewNamespace(q helper.QueryExecuter, configStore namespaceStore, configurations map[string]string) error {
	configjson, err := json.Marshal(configurations)
	isNamespaceExist, err := isNamespaceExist(q, configStore.ServiceId, configStore.Namespace)
	if isNamespaceExist {
		return errors.New("Namespace already exist for the service")
	}
	if err != nil {
		return err
	}
	_, err = q.DoQuery(createNewNamespaceQuery, configStore.ServiceId, configStore.Namespace)
	if err != nil {
		return err
	}
	_, err = q.DoQuery(createNewConfigurationsQuery, configStore.ServiceId, configStore.Namespace, configjson)

	return err
}

func retrieveAllNamespace(q helper.QueryExecuter, service_id int) ([]namespaceStore, error) {
	data, err := q.DoQuery(retrieveAllNamespaceQuery, service_id)
	if err != nil {
		return nil, err
	}
	var configurations []namespaceStore
	err = helper.ParseMaps(data, &configurations)
	return configurations, nil
}
