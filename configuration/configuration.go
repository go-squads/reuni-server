package configuration

type configView struct {
	Version       int               `json:"version"`
	Configuration map[string]string `json:"configuration"`
	Created_by    string            `json:"created_by"`
}

type versionView struct {
	Version int `json:"version"`
}

type versionsView struct {
	Versions []int `json:"versions"`
}
