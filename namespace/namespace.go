package namespace

import "time"

type namespaceStore struct {
	OrganizationId int       `json:"organization_id"`
	ServiceName    string    `json:"service_name"`
	Namespace      string    `json:"namespace"`
	ActiveVersion  int       `json:"version"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type namespaceView struct {
	Namespace     string                 `json:"namespace"`
	Configuration map[string]interface{} `json:"configurations"`
	CreatedBy     string                 `json:"created_by"`
}

func newNamespaceStruct(organizationId int, serviceName, name string, active int) *namespaceStore {
	return &namespaceStore{
		OrganizationId: organizationId,
		ServiceName:    serviceName,
		Namespace:      name,
		ActiveVersion:  active,
	}
}

func newNamespaceViewStruct(name string, configs map[string]interface{}) *namespaceView {
	return &namespaceView{
		Namespace:     name,
		Configuration: configs,
	}
}
