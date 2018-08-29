package configuration

import (
	"encoding/json"
	"fmt"
)

type Processor interface {
	getConfigurationProcess(organizationName, serviceName, namespace string, version int) (*configView, error)
	getLatestVersionProcess(organizationName, serviceName, namespace string) (int, error)
	createNewVersionProcess(organizationName, serviceName, namespace string, config configView) (int, error)
	getConfigurationVersionsProcess(organizationName, serviceName, namespace string) (string, error)
	getDifferenceProcessor(organizationName, serviceName, namespace string, version, previous_version int) ([]byte, error)
}

type mainProcessor struct {
	repo Repository
}

func (s *mainProcessor) getConfigurationProcess(organizationName, serviceName, namespace string, version int) (*configView, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return nil, err
	}
	config, err := s.repo.getConfiguration(organizationId, serviceName, namespace, version)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (s *mainProcessor) getLatestVersionProcess(organizationName, serviceName, namespace string) (int, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return 0, err
	}
	version, err := s.repo.getLatestVersionForNamespace(organizationId, serviceName, namespace)
	if err != nil {
		return 0, err
	}
	return version, nil
}

func (s *mainProcessor) createNewVersionProcess(organizationName, serviceName, namespace string, config configView) (int, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return 0, err
	}
	latestVersion, err := s.repo.getLatestVersionForNamespace(organizationId, serviceName, namespace)
	if err != nil {
		return 0, err
	}

	err = s.repo.createNewVersion(organizationId, serviceName, namespace, config, latestVersion+1)
	if err != nil {
		return 0, err
	}
	err = s.repo.updateNamespaceActiveVersion(organizationId, serviceName, namespace, latestVersion+1)
	if err != nil {
		return 0, err
	}
	return latestVersion + 1, nil
}

func (s *mainProcessor) getConfigurationVersionsProcess(organizationName, serviceName, namespace string) (string, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return "", err
	}

	versions, err := s.repo.getVersions(organizationId, serviceName, namespace)
	versionsv := versionsView{
		Versions: versions,
	}
	if err != nil {
		return "", err
	}
	versionsJSON, err := json.Marshal(versionsv)
	if err != nil {
		return "", err
	}
	return string(versionsJSON), nil

}

func (s *mainProcessor) getDifferenceProcessor(organizationName, serviceName, namespace string, version, previous_version int) ([]byte, error) {
	organizationId, err := s.repo.getOrganizationId(organizationName)
	if err != nil {
		return nil, err
	}
	config, err := s.repo.getConfiguration(organizationId, serviceName, namespace, version)
	if err != nil {
		return nil, err
	}
	differenceCollector := make(map[string][]map[string]string)

	if config.Parent_version != 0 {
		parentVersionConfig, err := s.repo.getConfiguration(organizationId, serviceName, namespace, int(config.Parent_version))
		if err != nil {
			return nil, err
		}
		differenceWithParentVersion, err := getDifference(config.Configuration, parentVersionConfig.Configuration)
		if err != nil {
			return nil, err
		}
		differenceCollector["dif-parent-"+fmt.Sprint(config.Parent_version)] = differenceWithParentVersion
	}
	if previous_version != 0 {
		previousVersionConfig, err := s.repo.getConfiguration(organizationId, serviceName, namespace, previous_version)
		if err != nil {
			return nil, err
		}
		differenceWithPreviousVersion, err := getDifference(config.Configuration, previousVersionConfig.Configuration)
		if err != nil {
			return nil, err
		}
		differenceCollector["dif-previous-"+fmt.Sprint(previous_version)] = differenceWithPreviousVersion
	}

	return json.Marshal(differenceCollector)
}

func getDifference(currentVersionConfig map[string]string, previousVersionConfig map[string]string) ([]map[string]string, error) {
	deletedConfig := getConfigurationDeleted(currentVersionConfig, previousVersionConfig)
	createdConfig := getConfigurationCreated(currentVersionConfig, previousVersionConfig)
	changedConfig := getConfigurationChanged(currentVersionConfig, previousVersionConfig)
	var configDifference []map[string]string
	configDifference = append(configDifference, deletedConfig, createdConfig, changedConfig)
	return configDifference, nil
}

func getConfigurationDeleted(currentConfig, previousConfig map[string]string) map[string]string {
	return getKeydifference(previousConfig, currentConfig)
}
func getConfigurationCreated(currentConfig, previousConfig map[string]string) map[string]string {
	return getKeydifference(currentConfig, previousConfig)
}

func getConfigurationChanged(currentConfig, previousConfig map[string]string) map[string]string {
	mapB := map[string]bool{}
	for k, _ := range previousConfig {
		mapB[k] = true
	}

	res := make(map[string]string)
	for k, v := range currentConfig {
		if _, ok := mapB[k]; ok {
			if previousConfig[k] != v {
				res[k] = "from '" + previousConfig[k] + "' to '" + v + "'"
			}
		}
	}

	return res
}

func getKeydifference(configA, configB map[string]string) map[string]string {
	mapB := map[string]bool{}
	for k, _ := range configB {
		mapB[k] = true
	}

	res := make(map[string]string)
	for k, v := range configA {
		if _, ok := mapB[k]; !ok {
			res[k] = v
		}
	}

	return res
}
