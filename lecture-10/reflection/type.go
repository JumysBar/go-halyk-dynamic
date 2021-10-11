package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	FieldOne string
	FieldTwo int
}

func PrintTypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case *MyStruct:
		fmt.Println("*MyStruct")
	}
}

func PrintType(x interface{}) {
	fmt.Printf("%T\n", x)
}

func PrintTypeReflect(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("%s\n", t.String())
}

func main() {
	x := 5
	// x := 3.14
	// x := &MyStruct{"123", 123}

	PrintTypeSwitch(x)
	PrintType(x)
	PrintTypeReflect(x)
}
