package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := `{	"Name": "Dracula", "Age": 590 }`

	// JSONobj := &struct {
	// 	Name         string
	// 	UnknownField int
	// }{}

	// json.Unmarshal([]byte(x), JSONobj)

	// fmt.Printf("Name: %s. Unknown field: %d\n", JSONobj.Name, JSONobj.UnknownField)

	// JSONobj := &struct {
	// 	name string
	// 	age  int
	// }{}

	// json.Unmarshal([]byte(x), JSONobj)

	// fmt.Printf("Name: %s. Age: %d\n", JSONobj.name, JSONobj.age)

	// JSONobj := &struct {
	// 	Name *string
	// 	Age  *int
	// }{}

	// json.Unmarshal([]byte(x), JSONobj)

	// fmt.Printf("Name: %s. Age: %d\n", *JSONobj.Name, *JSONobj.Age)

	JSONobj := &struct {
		Name         *string
		UnknownField *int
	}{}

	json.Unmarshal([]byte(x), JSONobj)

	fmt.Printf("Name: %s. Unknown field: %v\n", *JSONobj.Name, JSONobj.UnknownField)
}
