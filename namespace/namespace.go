package namespace

type configurationStore struct {
	Id             int               `json:"id"`
	ServiceId      int               `json:"service_id"`
	Namespace      string            `json:"namespace"`
	Version        int               `json:"version"`
	Configurations map[string]string `json:"configurations"`
}

type configurationView struct {
	Namespace     string            `json:"namespace"`
	Configuration map[string]string `json:"configurations"`
}
