package configuration

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

type Repository interface {
	getConfiguration(organizationId int, serviceName, namespace string, version int) (*configView, error)
	getLatestVersionForNamespace(organizationId int, serviceName, namespace string) (int, error)
	createNewVersion(organizationId int, serviceName, namespace string, config configView, version int) error
	updateNamespaceActiveVersion(organizationId int, serviceName, namespace string, version int) error
	getVersions(organizationId int, serviceName, namespace string) ([]int, error)
	getOrganizationId(organizationName string) (int, error)
}

type mainRepository struct {
	execer helper.QueryExecuter
}

const (
	getConfigurationQuery             = "SELECT version,config_store as configs, created_by FROM configurations WHERE organization_id=$1 AND service_name=$2 AND namespace=$3 AND version=$4"
	getLatestVersionForNamespaceQuery = "SELECT MAX(version) as latest FROM configurations WHERE organization_id=$1 AND service_name=$2 AND namespace=$3"
	createNewVersionQuery             = "INSERT INTO configurations(organization_id, service_name, namespace, version, config_store, created_by) VALUES($1,$2,$3,$4,$5,$6)"
	updateNamespaceActiveVersionQuery = "UPDATE namespaces SET active_version=$1 WHERE organization_id=$2 AND service_name=$3 AND namespace=$4"
	getVersionsQuery                  = "SELECT version FROM configurations WHERE organization_id=$1 AND service_name=$2 AND namespace=$3"
	translateNameToIdQuery            = "SELECT id FROM organization WHERE name=$1"
)

func (s *mainRepository) getConfiguration(organizationId int, serviceName, namespace string, version int) (*configView, error) {
	var config configView
	data, err := s.execer.DoQueryRow(getConfigurationQuery, organizationId, serviceName, namespace, version)
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

func (s *mainRepository) getLatestVersionForNamespace(organizationId int, serviceName, namespace string) (int, error) {
	var latestVersion int
	data, err := s.execer.DoQueryRow(getLatestVersionForNamespaceQuery, organizationId, serviceName, namespace)
	if err != nil {
		return 0, err
	}
	latestVersion = int(data["latest"].(int64))
	return latestVersion, nil
}

func (s *mainRepository) createNewVersion(organizationId int, serviceName, namespace string, config configView, version int) error {
	configJSON, err := json.Marshal(config.Configuration)
	if err != nil {
		return err
	}
	log.Println("created BYYYYYY: " + config.Created_by)
	_, err = s.execer.DoQuery(createNewVersionQuery, organizationId, serviceName, namespace, version, configJSON, config.Created_by)
	if err != nil {
		return err
	}
	return nil
}

func (s *mainRepository) updateNamespaceActiveVersion(organizationId int, serviceName, namespace string, version int) error {

	_, err := s.execer.DoQuery(updateNamespaceActiveVersionQuery, version, organizationId, serviceName, namespace)
	if err != nil {
		return err
	}
	return nil
}

func (s *mainRepository) getVersions(organizationId int, serviceName, namespace string) ([]int, error) {
	data, err := s.execer.DoQuerySlice(getVersionsQuery, organizationId, serviceName, namespace)
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

func (s *mainRepository) getOrganizationId(organizationName string) (int, error) {
	ret, err := s.execer.DoQueryRow(translateNameToIdQuery, organizationName)
	if err != nil {
		return 0, err
	}
	res, ok := ret["id"].(int64)
	if !ok {
		return 0, helper.NewHttpError(http.StatusNotFound, "Not Found")
	}
	return int(res), nil
}
