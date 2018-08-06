package helper

import (
	"encoding/json"
)

func ParseMaps(data interface{}, dest interface{}) error {
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
