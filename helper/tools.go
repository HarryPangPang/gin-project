package helper

import (
	"bytes"
	"encoding/json"
)

func String2Map(str string) map[string]interface{} {

	in := []byte(str)
	var personFromJSON interface{}

	decoder := json.NewDecoder(bytes.NewReader(in))
	decoder.UseNumber()
	decoder.Decode(&personFromJSON)

	result := personFromJSON.(map[string]interface{})
	return result
}

// 没有数字的时候可以使用
func JsonParse(str string) map[string]interface{} {
	var result map[string]interface{}
	json_body := []byte(str)
	json.Unmarshal(json_body, &result)
	return result
}
