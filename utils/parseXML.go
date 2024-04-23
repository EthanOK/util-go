package utils

import "encoding/xml"

func ParseXML(data []byte, struct_type interface{}) error {

	err := xml.Unmarshal(data, &struct_type)

	if err != nil {
		return err
	}

	return nil
}
