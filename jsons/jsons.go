package jsons

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"log"
	"reflect"
)

// ParseJson Json校验
func ParseJson(json string) (result gjson.Result, err error) {
	if !gjson.Valid(json) {
		err = errors.New("invalid json")
		return
	}
	result = gjson.Parse(json)
	return
}

// MapToJson map转为json字符串
func MapToJson(m map[string]interface{}) string {
	m2Json , _ := json.Marshal(m)
	map2String := string(m2Json)
	return map2String
}

// JsonToMap json 转map
func JsonToMap(jsonStr string) map[string]interface{} {
	var convert map[string]interface{}
	if jsonStr == "" {
		return convert
	}
	err := json.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		log.Println(err)
	}
	return convert
}


// JsonToMapArray json转map数组
func JsonToMapArray(jsonStr string) []map[string]interface{} {
	var convert []map[string]interface{}
	if jsonStr == "" {
		return convert
	}
	err := json.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		log.Println(err)
	}
	return convert
}

// StructToMap 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	typeOf := reflect.TypeOf(obj)
	valueOf := reflect.ValueOf(obj)
	convert := make(map[string]interface{})
	for i := 0; i < typeOf.NumField(); i ++ {
		convert[typeOf.Field(i).Name] = valueOf.Field(i).Interface()
	}
	return convert
}
