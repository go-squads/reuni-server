package namespace

import "time"

type namespaceStore struct {
	Id            int        `json:"id"`
	ServiceId     int        `json:"service_id"`
	Namespace     string     `json:"namespace"`
	ActiveVersion int        `json:"version"`
	CreatedBy     string     `json:"created_by"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type namespaceView struct {
	Namespace     string                 `json:"namespace"`
	Configuration map[string]interface{} `json:"configurations"`
}

func newNamespaceStruct(id int, serviceId int, name string, active int) *namespaceStore {
	return &namespaceStore{
		Id:            id,
		ServiceId:     serviceId,
		Namespace:     name,
		ActiveVersion: active,
	}
}

func newNamespaceViewStruct(name string, configs map[string]interface{}) *namespaceView {
	return &namespaceView{
		Namespace:     name,
		Configuration: configs,
	}
}
