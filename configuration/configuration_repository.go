package configuration

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

type Repository interface {
	getConfiguration(serviceId int, namespace string, version int) (*configView, error)
	getLatestVersionForNamespace(serviceId int, namespace string) (int, error)
	createNewVersion(serviceId int, namespace string, config configView, version int) error
	updateNamespaceActiveVersion(qserviceId int, namespace string, version int) error
	getVersions(serviceId int, namespace string) ([]int, error)
	getServiceId(serviceName string) (int, error)
}

type mainRepository struct {
	execer helper.QueryExecuter
}

const (
	getConfigurationQuery             = "SELECT version,config_store as configs FROM configurations WHERE service_id=$1 AND namespace=$2 AND version=$3"
	getLatestVersionForNamespaceQuery = "SELECT MAX(version) as latest FROM configurations WHERE service_id=$1 AND namespace=$2"
	createNewVersionQuery             = "INSERT INTO configurations(service_id, namespace, version, config_store) VALUES($1,$2,$3,$4)"
	updateNamespaceActiveVersionQuery = "UPDATE namespaces SET active_version=$1 WHERE service_id=$2 AND namespace=$3"
	getVersionsQuery                  = "SELECT version FROM configurations WHERE service_id=$1 AND namespace=$2"
	findServiceIdFromName             = "SELECT id FROM services WHERE name=$1"
)

func (s *mainRepository) getConfiguration(serviceId int, namespace string, version int) (*configView, error) {
	var config configView
	data, err := s.execer.DoQueryRow(getConfigurationQuery, serviceId, namespace, version)
	if err != nil {
		return nil, err
	}
	err = helper.ParseMap(data, &config)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, helper.NewHttpError(http.StatusNotFound, "Data not found")
	}
	bytes, ok := data["configs"].([]byte)
	if !ok {
		return nil, errors.New("Cannot parse config_store")
	}
	json.Unmarshal(bytes, &config.Configuration)
	return &config, nil
}

func (s *mainRepository) getLatestVersionForNamespace(serviceId int, namespace string) (int, error) {
	var latestVersion int
	data, err := s.execer.DoQueryRow(getLatestVersionForNamespaceQuery, serviceId, namespace)
	if err != nil {
		return 0, err
	}
	latestVersion = int(data["latest"].(int64))
	return latestVersion, nil
}

func (s *mainRepository) createNewVersion(serviceId int, namespace string, config configView, version int) error {
	configJSON, err := json.Marshal(config.Configuration)
	if err != nil {
		return err
	}
	_, err = s.execer.DoQuery(createNewVersionQuery, serviceId, namespace, version, configJSON)
	if err != nil {
		return err
	}
	return nil
}

func (s *mainRepository) updateNamespaceActiveVersion(serviceId int, namespace string, version int) error {

	_, err := s.execer.DoQuery(updateNamespaceActiveVersionQuery, version, serviceId, namespace)
	if err != nil {
		return err
	}
	return nil
}

func (s *mainRepository) getVersions(serviceId int, namespace string) ([]int, error) {
	data, err := s.execer.DoQuerySlice(getVersionsQuery, serviceId, namespace)
	if err != nil {
		return nil, err
	}
	var versions []int
	err = helper.ParseMap(data, &versions)
	if err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *mainRepository) getServiceId(serviceName string) (int, error) {
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
