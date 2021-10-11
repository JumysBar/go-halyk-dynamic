package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 string
	Field2 int
	field3 bool
}

func main() {
	x1 := &MyStruct{"Hello", 1, false}
	x2 := &MyStruct{"Hello", 1, false}
	// x2 := &MyStruct{"Hello", 2, false}
	// x2 := &MyStruct{"Hello", 1, true}

	ok := reflect.DeepEqual(x1, x2)
	if ok {
		fmt.Println("Objects are equal")
	} else {
		fmt.Println("Objects are not equal")
	}

	y1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	y2 := map[string]int{
		"key1": 1,
		// "key2": 2,
	}

	ok = reflect.DeepEqual(y1, y2)
	if ok {
		fmt.Println("Objects are equal")
	} else {
		fmt.Println("Objects are not equal")
	}

	z1 := []int{1, 2, 3}
	z2 := []int{1, 2, 3}
	// z2 := []float32{1, 2, 3}

	// z2 := make([]int, 0, 10)
	// z2 = append(z2, 1)
	// z2 = append(z2, 2)
	// z2 = append(z2, 3)

	ok = reflect.DeepEqual(z1, z2)
	if ok {
		fmt.Println("Objects are equal")
	} else {
		fmt.Println("Objects are not equal")
	}
}
