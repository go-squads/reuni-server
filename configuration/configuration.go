package configuration

type configView struct {
	Version       int               `json:"version"`
	Configuration map[string]string `json:"configuration "`
}
