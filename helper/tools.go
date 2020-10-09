package helper

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
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

// 获取权限游戏列表
func GetGameList(privs []interface{}) []interface{} {
	var i int
	acc := make([]interface{}, 0)
	for i = 0; i < len(privs); i++ {
		name := privs[i].(map[string]interface{})
		n := name["name"].(string)
		matched, _ := regexp.MatchString(`[\d]-[\d]-*`, n)
		if matched {
			nList := strings.Split(n, "-")
			acc = append(acc, nList)
		}
	}
	return acc
}
