package configurations

import (
	"encoding/json"

	context "github.com/go-squads/reuni-server/appcontext"
)

const createNewNamespaceQuery = "INSERT INTO configurations(service_id, namespace,config_store) VALUES ($1,$2,$3)"
const retrieveAllNamespaceQuery = "SELECT namespace,MAX(version) namespace FROM configurations WHERE service_id = $1 GROUP BY namespace"

func createNewNamespace(configStore configurationStore) error {
	db := context.GetDB()
	configjson, err := json.Marshal(configStore.Configurations)
	if err != nil {
		return err
	}
	_, err = db.Query(createNewNamespaceQuery, configStore.ServiceId, configStore.Namespace, configjson)
	return err
}

func retrieveAllNamespace(service_id int) ([]configurationStore, error) {
	db := context.GetDB()
	rows, err := db.Query(retrieveAllNamespaceQuery, service_id)
	if err != nil {
		return nil, err
	}
	var configurations []configurationStore
	for rows.Next() {
		var configuration configurationStore
		err = rows.Scan(&configuration.Namespace, &configuration.Version)
		if err != nil {
			return nil, err
		}
		configurations = append(configurations, configuration)
	}
	return configurations, nil
}
