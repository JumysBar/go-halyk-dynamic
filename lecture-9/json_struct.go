package main

import (
	"encoding/json"
	"fmt"
)

type Credentials struct {
	Login    string
	Password string
}

type Language struct {
	Title string
	Words []string
}

type Profile struct {
	Name      string
	Age       int64
	Languages []Language
	Credentials
}

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
	JSONobj := &Profile{}
	if err := json.Unmarshal([]byte(x), JSONobj); err != nil {
		fmt.Printf("JSON unmarshal error: %v\n", err)
		return
	}
	fmt.Printf("Name: %s\n", JSONobj.Name)
	for _, lang := range JSONobj.Languages {
		fmt.Printf("Title: %s\n", lang.Title)
		for _, word := range lang.Words {
			fmt.Printf("Word: %s\n", word)
		}
		fmt.Println()
	}

	// json marshal
	JSONobj.Login = "admin"
	JSONobj.Password = "admin"

	data, _ := json.MarshalIndent(JSONobj, "", "  ")
	fmt.Printf("JSON: %s\n", data)
}
