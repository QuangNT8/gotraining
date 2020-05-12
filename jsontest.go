package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func main() {

	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})
	fmt.Println(v)
	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}

}

/*
func main() {
	var tests = []string{
		`"Platypus"`,
		`Platypus`,
		`{"id":"1"}`,
	}

	for _, t := range tests {
		fmt.Printf("isJSONString(%s) = %v\n", t, isJSONString(t))
		fmt.Printf("isJSON(%s) = %v\n\n", t, isJSON(t))
	}

} //*/
