package helper

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func ParseMap(data interface{}, dest interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	log.Println("Unmarshaling %v", string(jsonData))
	err = json.Unmarshal(jsonData, dest)
	if err != nil {
		return err
	}
	return nil
}
