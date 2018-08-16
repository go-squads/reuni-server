package configuration

import (
	"encoding/json"
)

type Processor interface {
	getConfigurationProcess(serviceName, namespace string, version int) (*configView, error)
	getLatestVersionProcess(serviceName, namespace string) (int, error)
	createNewVersionProcess(serviceName, namespace string, config configView) (int, error)
	getConfigurationVersionsProcess(serviceName, namespace string) (string, error)
}

type mainProcessor struct {
	repo Repository
}

func (s *mainProcessor) getConfigurationProcess(serviceName, namespace string, version int) (*configView, error) {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return nil, err
	}
	config, err := s.repo.getConfiguration(serviceId, namespace, version)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (s *mainProcessor) getLatestVersionProcess(serviceName, namespace string) (int, error) {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return 0, err
	}
	version, err := s.repo.getLatestVersionForNamespace(serviceId, namespace)
	if err != nil {
		return 0, err
	}
	return version, nil
}

func (s *mainProcessor) createNewVersionProcess(serviceName, namespace string, config configView) (int, error) {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return 0, err
	}
	latestVersion, err := s.repo.getLatestVersionForNamespace(serviceId, namespace)
	if err != nil {
		return 0, err
	}

	err = s.repo.createNewVersion(serviceId, namespace, config, latestVersion+1)
	if err != nil {
		return 0, err
	}
	err = s.repo.updateNamespaceActiveVersion(serviceId, namespace, latestVersion+1)
	if err != nil {
		return 0, err
	}
	return latestVersion + 1, nil
}

func (s *mainProcessor) getConfigurationVersionsProcess(serviceName, namespace string) (string, error) {
	serviceId, err := s.repo.getServiceId(serviceName)
	if err != nil {
		return "", err
	}

	versions, err := s.repo.getVersions(serviceId, namespace)
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
