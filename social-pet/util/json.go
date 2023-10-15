package util

import (
	"encoding/json"
	"log"
)

func MapToStruct(input map[string]interface{}, obj interface{}) (err error) {
	var (
		b []byte
	)
	if b, err = json.Marshal(input); err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(b, &obj); err != nil {
		log.Println(err)
		return
	}
	return
}

//结构体转map
func StructToMap(input interface{}) (obj map[string]interface{}, err error) {
	var (
		b []byte
	)
	if b, err = json.Marshal(input); err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(b, &obj); err != nil {
		log.Println(err)
		return
	}
	return
}

func InterfaceToStruct(input interface{}, obj interface{}) (err error) {
	var (
		b []byte
	)
	if b, err = json.Marshal(input); err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(b, &obj); err != nil {
		return
	}
	return
}

/*结构体转json*/
func StructToJson(obj interface{}) (b []byte, err error) {

	if b, err = json.Marshal(obj); err != nil {
		return
	}
	return
}

/*结构体转json*/
func StructToJsonString(obj interface{}) (j string) {
	var (
		b   []byte
		err error
	)
	if b, err = json.Marshal(obj); err != nil {
		return
	}
	j = string(b)
	return
}

/*json字符串转结构体*/
func JsonStrToStruct(input string, obj interface{}) (err error) {
	if err = json.Unmarshal([]byte(input), &obj); err != nil {
		return
	}
	return
}
