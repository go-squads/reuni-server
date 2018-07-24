package namespace

type namespaceStore struct {
	Id            int    `json:"id"`
	ServiceId     int    `json:"service_id"`
	Namespace     string `json:"namespace"`
	ActiveVersion int    `json:"version"`
}

type namespaceView struct {
	Namespace     string            `json:"namespace"`
	Configuration map[string]string `json:"configurations"`
}
