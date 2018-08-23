package configuration

import (
	"encoding/json"
)

type Processor interface {
	getConfigurationProcess(organizationName, serviceName, namespace string, version int) (*configView, error)
	getLatestVersionProcess(organizationName, serviceName, namespace string) (int, error)
	createNewVersionProcess(organizationName, serviceName, namespace string, config configView) (int, error)
	getConfigurationVersionsProcess(organizationName, serviceName, namespace string) (string, error)
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
