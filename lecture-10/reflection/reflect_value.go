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
	v := reflect.ValueOf(input)

	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("input is integer. Value: %d\n", v.Int())
	case reflect.String:
		fmt.Printf("input is string. Value: %d\n", v.String())
		// case reflect.Struct
	case reflect.Ptr:
		val := v.Elem()
		fmt.Printf("input is pointer on %s type\n", val.Type().String())
		switch val.Kind() {
		case reflect.Int:
			fmt.Printf("input is pointer on integer. Value: %d\n", val.Int())
			val.SetInt(123)
		case reflect.Struct:
			fmt.Println("input is pointer on struct")

			// call method
			method, _ := val.Type().MethodByName("GetGpu")
			ret := method.Func.Call([]reflect.Value{val})
			fmt.Printf("Return values: %v\n", ret)

			// change fields
			for i := 0; i < val.NumField(); i++ {
				field := val.Field(i)
				if field.CanSet() && field.Kind() == reflect.String {
					field.SetString("Changed string")
				}
			}
		}
	}
}

func main() {
	// x := 10
	x := GameMachine{"PS10", "Мощный процессор", "Мощная видеокарта", 1e10}

	fmt.Println(x)
	SomeFunction(&x)
	fmt.Println(x)
}
