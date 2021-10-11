package main

import (
	"fmt"
	"reflect"
)

type GameMachine struct {
	ModelName string
	Processor string
	GPU       string
	Flops     float64
}

func (t *GameMachine) SetGpu(gpu string) {
	t.GPU = gpu
}

func (t *GameMachine) SetProcessor(proc string) {
	t.Processor = proc
}

func (t GameMachine) GetGpu() string {
	return t.GPU
}

func SomeFunction(input interface{}) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("input is integer")
	case reflect.String:
		fmt.Println("input is string")

	case reflect.Struct:
		fmt.Println("input is some struct")
		fmt.Println("------------")
		fmt.Printf("input has %d fields\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("Field %s has type %s\n", t.Field(i).Name, t.Field(i).Type.String())
		}
		fmt.Println("------------")
		fallthrough
	case reflect.Ptr:
		fmt.Printf("input has %d methods\n", t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Printf("Method %s has %d arguments\n", t.Method(i).Name, t.Method(i).Type.NumIn())
			for j := 0; j < t.Method(i).Type.NumIn(); j++ {
				fmt.Printf("Argument %d has type %s\n", j, t.Method(i).Type.In(j).String())
			}
		}
	}
}

func main() {
	x := 10
	// x := "Hello world!"
	// x := GameMachine{"PS10", "Мощный процессор", "Мощная видеокарта", 1e10}
	SomeFunction(x)
}
