package services

import "time"

type service struct {
	Id                 int       `json:"id"`
	Name               string    `json:"name"`
	AuthorizationToken string    `json:"authorization_token"`
	CreatedAt          time.Time `json:"created_at"`
}

type services []service

type servicev struct {
	Name string `json:"name"`
}

type serviceToken struct {
	Token string `json:"token"`
}
