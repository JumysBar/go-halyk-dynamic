package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := `
	{
		"key1": "value",
		"key2": 123
	}
	`
	// JSON unmarshaling
	var JSONmap map[string]interface{}
	err := json.Unmarshal([]byte(x), &JSONmap)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return
	}
	for key, value := range JSONmap {
		fmt.Printf("Key: %s. Value: %v\n", key, value)
	}

	// JSON marshaling
	JSONmap["key3"] = true
	JSONmap["key4"] = []int{1, 2, 3}
	JSONmap["key5"] = nil

	// data, err := json.Marshal(make(chan int))
	data, err := json.Marshal(JSONmap)
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
		return
	}

	fmt.Printf("JSON string: %s\n", data)
}
