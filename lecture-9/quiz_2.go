package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// x := &struct {
	// 	Name string
	// 	age  int
	// }{"Dracula", 590}

	// data, _ := json.Marshal(x)
	// fmt.Println(string(data))

	// name := "Dracula"
	// x := &struct {
	// 	Name *string
	// 	Age  *int
	// }{&name, nil}

	// data, _ := json.Marshal(x)
	// fmt.Println(string(data))

	x := &struct {
		Name    string
		Victims []int
	}{"Dracula", nil}

	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
