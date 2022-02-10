// @description: json序列化与反序列化
// @file: json.go
// @date: 2022/02/10

package util

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
