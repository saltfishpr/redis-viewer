package util

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func JsonMarshalIndent(data interface{}) ([]byte, error) {
	return json.MarshalIndent(&data, "", "  ")
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
