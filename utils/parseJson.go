package utils

import "encoding/json"

func ParseJson(data []byte, struct_type interface{}) error {

	err := json.Unmarshal(data, &struct_type)

	if err != nil {
		return err
	}

	return nil
}
