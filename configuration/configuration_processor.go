package configuration

import (
	"encoding/json"

	context "github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/services"
)

func getConfigurationProcess(serviceName, namespace string, version int) (*configView, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return nil, err
	}
	config, err := getConfiguration(context.GetHelper(), service.Id, namespace, version)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func getLatestVersionProcess(serviceName, namespace string) (int, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return 0, err
	}
	version, err := getLatestVersionForNamespace(context.GetHelper(), service.Id, namespace)
	if err != nil {
		return 0, err
	}
	return version, nil
}

func createNewVersionProcess(serviceName, namespace string, config configView) error {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return err
	}
	latestVersion, err := getLatestVersionForNamespace(context.GetHelper(), service.Id, namespace)
	if err != nil {
		return err
	}

	err = createNewVersion(context.GetHelper(), service.Id, namespace, config, latestVersion+1)
	if err != nil {
		return err
	}
	err = updateNamespaceActiveVersion(context.GetHelper(), service.Id, namespace, latestVersion+1)
	if err != nil {
		return err
	}
	return nil
}

func getConfigurationVersionsProcess(serviceName, namespace string) (string, error) {
	service, err := services.FindOneServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	versions, err := getVersions(context.GetHelper(), service.Id, namespace)
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
