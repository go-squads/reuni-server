package namespace

import (
	"encoding/json"
	"errors"

	context "github.com/go-squads/reuni-server/appcontext"
)

const createNewNamespaceQuery = "INSERT INTO namespaces(service_id, namespace) VALUES ($1,$2)"
const createNewConfigurationsQuery = "INSERT INTO configurations(service_id, namespace, config_store) VALUES ($1,$2,$3)"
const retrieveAllNamespaceQuery = "SELECT namespace,active_version namespace FROM namespaces WHERE service_id = $1"
const countNamespaceNameByService = "SELECT count(namespace) FROM namespaces WHERE service_id=$1 AND namespace=$2"

func isNamespaceExist(service_id int, namespace string) (bool, error) {
	var count int
	db := context.GetDB()
	row := db.QueryRow(countNamespaceNameByService, service_id, namespace)

	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func createNewNamespace(configStore namespaceStore, configurations map[string]string) error {
	db := context.GetDB()
	configjson, err := json.Marshal(configurations)
	isNamespaceExist, err := isNamespaceExist(configStore.ServiceId, configStore.Namespace)
	if isNamespaceExist {
		return errors.New("Namespace already exist for the service")
	}
	if err != nil {
		return err
	}
	_, err = db.Query(createNewNamespaceQuery, configStore.ServiceId, configStore.Namespace)
	if err != nil {
		return err
	}
	_, err = db.Query(createNewConfigurationsQuery, configStore.ServiceId, configStore.Namespace, configjson)

	return err
}

func retrieveAllNamespace(service_id int) ([]namespaceStore, error) {
	db := context.GetDB()
	rows, err := db.Query(retrieveAllNamespaceQuery, service_id)
	if err != nil {
		return nil, err
	}
	var configurations []namespaceStore
	for rows.Next() {
		var configuration namespaceStore
		err = rows.Scan(&configuration.Namespace, &configuration.ActiveVersion)
		if err != nil {
			return nil, err
		}
		configurations = append(configurations, configuration)
	}
	return configurations, nil
}
