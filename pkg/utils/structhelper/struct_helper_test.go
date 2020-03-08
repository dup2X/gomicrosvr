package structhelper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	IntA    int              `json:"int_a,omitemtpy"`
	StringB string           `json:"string_b,omitempty"`
	Netted  TestNettedStruct `json:"netted"`
}

type TestTwoStruct struct {
	IntC *int `json:"int_c,omitempty"`
}

type TestNettedStruct struct {
	IntD int `json:"int_d"`
}

type TestBaseStruct struct {
	TestStruct
	TestTwoStruct
}

type TestBaseMoreStruct struct {
	*TestStruct
	*TestTwoStruct
	Sign string `json:"sign"`
}

func TestConvertStructToUrlValues(t *testing.T) {
	data := &TestStruct{
		IntA:    10,
		StringB: "hehehe",
	}
	params, err := ConvertStructToURLValues(data)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	fmt.Printf("params = %+v\n", params)
}

func TestConvertPointerStructToUrlValues(t *testing.T) {
	var intC = 3
	data := &TestBaseStruct{
		TestStruct{
			IntA:    90,
			StringB: "hehe",
		},
		TestTwoStruct{
			IntC: &intC,
		},
	}
	params, err := ConvertStructToURLValues(data)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	fmt.Printf("params = %+v\n", params)
}

func TestOmitEmpty(t *testing.T) {
	data := &TestTwoStruct{
		IntC: nil,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	fmt.Printf("json marshal point: %s\n", string(bytes))
	newData := TestTwoStruct{}
	err = json.Unmarshal(bytes, &newData)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	fmt.Printf("json unmarshal point: %+v\n", newData)
}

func TestConvertStructToMap(t *testing.T) {
	var intC = 3
	data := &TestBaseStruct{
		TestStruct{
			IntA:    90,
			StringB: "hehe",
			Netted: TestNettedStruct{
				IntD: 200,
			},
		},
		TestTwoStruct{
			IntC: &intC,
		},
	}
	ret, err := ConvertStructToMap(data)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	for key, val := range ret {
		switch t := val.(type) {
		default:
			fmt.Printf("key=%s, val=%v, type=%v\n", key, val, t)
		}
	}
	fmt.Printf("ret = %+v\n", ret)
}

func TestMarshUnMarshalBigInt(t *testing.T) {
	var intC = 324150239
	data := TestTwoStruct{
		IntC: &intC,
	}
	ret, err := ConvertStructToMap(data)
	if err != nil {
		t.Errorf("err = %v", err)
	}
	fmt.Printf("bigIntType : %v\n", reflect.TypeOf(ret["int_c"]))
	fmt.Printf("ret = %+v\n", ret)
}
