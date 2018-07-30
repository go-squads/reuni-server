package configuration

import (
	"encoding/json"

	context "github.com/go-squads/reuni-server/appcontext"
)

const (
	getConfigurationQuery             = "SELECT version,config_store FROM configurations WHERE service_id=$1 AND namespace=$2 AND version=$3"
	getLatestVersionForNamespaceQuery = "SELECT MAX(version) FROM configurations WHERE service_id=$1 AND namespace=$2"
	createNewVersionQuery             = "INSERT INTO configurations(service_id, namespace, version, config_store) VALUES($1,$2,$3,$4)"
	updateNamespaceActiveVersionQuery = "UPDATE namespaces SET active_version=$1 WHERE service_id=$2 AND namespace=$3"
)

func getConfiguration(serviceId int, namespace string, version int) (*configView, error) {
	var config configView
	var configJSON string
	row := context.GetDB().QueryRow(getConfigurationQuery, serviceId, namespace, version)
	err := row.Scan(&config.Version, &configJSON)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(configJSON), &config.Configuration)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func getLatestVersionForNamespace(serviceId int, namespace string) (int, error) {
	var latestVersion int
	row := context.GetDB().QueryRow(getLatestVersionForNamespaceQuery, serviceId, namespace)
	err := row.Scan(&latestVersion)
	if err != nil {
		return 0, err
	}
	return latestVersion, nil
}

func createNewVersion(serviceId int, namespace string, config configView, version int) error {
	configJSON, err := json.Marshal(config.Configuration)
	if err != nil {
		return err
	}
	_, err = context.GetDB().Query(createNewVersionQuery, serviceId, namespace, version, configJSON)
	if err != nil {
		return err
	}
	return nil
}

func updateNamespaceActiveVersion(serviceId int, namespace string, version int) error {

	_, err := context.GetDB().Query(updateNamespaceActiveVersionQuery, version, serviceId, namespace)
	if err != nil {
		return err
	}
	return nil
}
