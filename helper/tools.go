package helper

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

type MentItem struct {
	GameId  string   `json:"gameId"`
	MenuIds []string `json:"menuIds"`
}

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

func existGameIdItem(games []MentItem, gameId string) int {
	for i := 0; i < len(games); i++ {
		if games[i].GameId == gameId {
			return i + 1
		}
	}
	return 0
}

// 获取权限列表
func GetGameList(privs []interface{}) []MentItem {
	var i int
	acc := make([]MentItem, 0)
	for i = 0; i < len(privs); i++ {
		name := privs[i].(map[string]interface{})
		n := name["name"].(string)
		matched, _ := regexp.MatchString(`[\d]-[\d]-*`, n)
		if matched {
			nList := strings.Split(n, "-")
			var menuIds []string
			menuIds = append(menuIds, nList[1])
			menu := MentItem{nList[0], menuIds}
			existIndex := existGameIdItem(acc, nList[0])
			if existIndex != 0 {
				acc[existIndex-1].MenuIds = append(acc[existIndex-1].MenuIds, nList[1])
			} else {
				acc = append(acc, menu)
			}
		}
	}
	return acc
}
