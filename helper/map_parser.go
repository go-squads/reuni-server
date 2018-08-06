package helper

import (
	"encoding/json"
)

func ParseMap(data interface{}, dest interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, dest)
	if err != nil {
		return err
	}
	return nil
}
