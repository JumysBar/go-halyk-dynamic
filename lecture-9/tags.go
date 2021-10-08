package main

import (
	"encoding/json"
	"fmt"
)

type Computer struct {
	FieldName        string `json:"processor"`
	AnotherFieldName int    `json:"ram"`
	ThirdField       bool   `json:"ssd"`
}

func main() {
	x := `
	{
		"processor": "Intel core i7",
		"ram": 16
	}`

	// JSON unmarshal
	comp := &Computer{}
	json.Unmarshal([]byte(x), comp)
	fmt.Printf("Field: %s. Another field: %d\n", comp.FieldName, comp.AnotherFieldName)

	// JSON marshal
	comp.ThirdField = true
	data, _ := json.Marshal(comp)
	fmt.Printf("JSON: %s\n", data)
}
