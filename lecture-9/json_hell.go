package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := `
	{
		"Name": "Vladimir",
		"Age": 23,
		"Languages": [
			{
				"Title": "Russian",
				"Words": [
					"Привет",
					"Мир"
				]
			},
			{
				"Title": "English",
				"Words": [
					"Hello",
					"World"
				]
			},
			{
				"Title": "Korean",
				"Words": [
					"이봐",
					"평화"
				]
			}
		]
	}
	`
	var JSONmap map[string]interface{}
	if err := json.Unmarshal([]byte(x), &JSONmap); err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return
	}
	languages, ok := JSONmap["Languages"].([]interface{})
	if !ok {
		fmt.Println("Field 'Languages' has invalid type")
		return
	}
	for i, lang := range languages {
		mapLang, ok := lang.(map[string]interface{})
		if !ok {
			fmt.Printf("Array element with index %d has incorrect type", i)
			return
		}
		fmt.Printf("Title: %s\n", mapLang["Title"])
		words, ok := mapLang["Words"].([]interface{})
		// words, ok := mapLang["Words"].([]string)
		if !ok {
			fmt.Println("Field 'Words' has invalid type")
			return
		}
		for _, word := range words {
			strWord, ok := word.(string)
			if !ok {
				fmt.Printf("Array element with index %d has incorrect type", i)
				return
			}
			fmt.Println("Word:", strWord)
		}
		fmt.Println()
	}
}
