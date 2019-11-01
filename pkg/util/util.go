package util

import (
	"reflect"
	"encoding/json"
)

// Struct to map
func StructToMap(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// Struct Json to map
func StructJsonToMap(obj interface{}) (map[string]interface{}, error){
	jval, err := json.Marshal(obj);
	if err != nil {
		return nil, err
	}

	var data = make(map[string]interface{})
	if err := json.Unmarshal([]byte(jval), &data); err == nil {
		return data, err
	}

	return nil, err
}