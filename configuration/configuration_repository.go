package configuration

import (
	"encoding/json"
	"log"

	"github.com/go-squads/reuni-server/helper"
)

const (
	getConfigurationQuery             = "SELECT version,config_store FROM configurations WHERE service_id=$1 AND namespace=$2 AND version=$3"
	getLatestVersionForNamespaceQuery = "SELECT MAX(version) FROM configurations WHERE service_id=$1 AND namespace=$2"
	createNewVersionQuery             = "INSERT INTO configurations(service_id, namespace, version, config_store) VALUES($1,$2,$3,$4)"
	updateNamespaceActiveVersionQuery = "UPDATE namespaces SET active_version=$1 WHERE service_id=$2 AND namespace=$3"
	getVersionsQuery                  = "SELECT version FROM configurations WHERE service_id=$1 AND namespace=$2"
)

func getConfiguration(q helper.QueryExecuter, serviceId int, namespace string, version int) (*configView, error) {
	var config configView
	data, err := q.DoQuery(getConfigurationQuery, serviceId, namespace, version)
	if err != nil {
		return nil, err
	}
	err = helper.ParseMap(data[0], &config.Configuration)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func getLatestVersionForNamespace(q helper.QueryExecuter, serviceId int, namespace string) (int, error) {
	var latestVersion int
	data, err := q.DoQuery(getLatestVersionForNamespaceQuery, serviceId, namespace)
	if err != nil {
		return 0, err
	}
	log.Println(data)
	err = helper.ParseMap(data[0]["version"], &latestVersion)
	return latestVersion, nil
}

func createNewVersion(q helper.QueryExecuter, serviceId int, namespace string, config configView, version int) error {
	configJSON, err := json.Marshal(config.Configuration)
	if err != nil {
		return err
	}
	_, err = q.DoQuery(createNewVersionQuery, serviceId, namespace, version, configJSON)
	if err != nil {
		return err
	}
	return nil
}

func updateNamespaceActiveVersion(q helper.QueryExecuter, serviceId int, namespace string, version int) error {

	_, err := q.DoQuery(updateNamespaceActiveVersionQuery, version, serviceId, namespace)
	if err != nil {
		return err
	}
	return nil
}

func getVersions(q helper.QueryExecuter, serviceId int, namespace string) ([]int, error) {
	data, err := q.DoQuery(getVersionsQuery, serviceId, namespace)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var versions []int
	log.Print(data)
	err = helper.ParseMap(data, &versions)
	return versions, nil
}
