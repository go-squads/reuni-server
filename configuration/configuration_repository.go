package configuration

import (
	"encoding/json"

	context "github.com/go-squads/reuni-server/appcontext"
)

const (
	getConfigurationQuery             = "SELECT version,config_store FROM configurations WHERE service_id=$1 AND namespace=$2 AND version=$3"
	getLatestVersionForNamespaceQuery = "SELECT MAX(version) FROM configurations WHERE service_id=$1 AND namespace=$2"
)

func getConfiguration(service_id int, namespace string, version int) (*configView, error) {
	var config configView
	var configJSON string
	row := context.GetDB().QueryRow(getConfigurationQuery, service_id, namespace, version)
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

func getLatestVersionForNamespace(service_id int, namespace string) (int, error) {
	var latestVersion int
	row := context.GetDB().QueryRow(getLatestVersionForNamespaceQuery, service_id, namespace)
	err := row.Scan(&latestVersion)
	if err != nil {
		return 0, err
	}
	return latestVersion, err
}
