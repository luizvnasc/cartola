package util

import "encoding/json"

func MapToStruct(m map[string]interface{}, s interface{}) error {
	str, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(str, s)
}

func StructToMap(s interface{}) (m map[string]interface{}, err error) {
	str, err := json.Marshal(s)
	if err != nil {
		return
	}
	err = json.Unmarshal(str, &m)
	return
}
