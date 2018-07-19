package services

import "time"

type service struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type servicev struct {
	Name string `json:"name"`
}
