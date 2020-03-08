package structhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

//ConvertStructToURLValues ...
func ConvertStructToURLValues(data interface{}) (url.Values, error) {
	ret := url.Values{}
	//convert to json string
	byteJSON, err := json.Marshal(data)
	if err != nil {
		return ret, err
	}
	//convert to map
	dataMap := map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewBuffer(byteJSON))
	decoder.UseNumber()
	err = decoder.Decode(&dataMap)
	if err != nil {
		return ret, err
	}

	// convert to url.Values
	for k, v := range dataMap {
		var jstr string
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Map:
			j, _ := json.Marshal(v)
			jstr = string(j)
		case reflect.Slice:
			j, _ := json.Marshal(v)
			jstr = string(j)
		default:
			jstr = fmt.Sprint(v)
		}
		ret.Add(k, jstr)
	}
	return ret, err
}

//ConvertStructToMap ...
func ConvertStructToMap(data interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	//convert to json string
	byteJSON, err := json.Marshal(data)
	if err != nil {
		return ret, err
	}
	//convert to map
	decoder := json.NewDecoder(bytes.NewBuffer(byteJSON))
	decoder.UseNumber()
	decoder.Decode(&ret)
	return ret, err
}
