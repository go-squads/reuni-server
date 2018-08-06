package services

import "time"

type service struct {
	Id                 int        `json:"id"`
	Name               string     `json:"name"`
	AuthorizationToken string     `json:"authorization_token"`
	CreatedAt          *time.Time `json:"created_at"`
}

type servicev struct {
	Name string `json:"name"`
}

type serviceToken struct {
	Token string `json:"token"`
}

func (s *service) IsEmpty() bool {
	return (s.Id == 0 && s.Name == "" && s.AuthorizationToken == "" && s.CreatedAt == nil)
}

func (s *serviceToken) IsEmpty() bool {
	return s.Token == ""
}
func isSliceEmpty(s []service) bool {
	for _, ss := range s {
		if !ss.IsEmpty() {
			return false
		}
	}
	return true
}
