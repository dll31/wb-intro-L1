// Задача: Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	collection := make([]interface{}, 0)

	collection = append(collection, "SomeString")
	collection = append(collection, false)
	collection = append(collection, 33)
	collection = append(collection, make(chan struct{}))
	collection = append(collection, 0.22)

	for _, v := range collection {
		switch reflect.ValueOf(v).Kind() {
		case reflect.Int:
			fmt.Printf("%v is int\n", v)
		case reflect.String:
			fmt.Printf("%v is string\n", v)
		case reflect.Chan:
			fmt.Printf("%v is chan\n", v)
		case reflect.Bool:
			fmt.Printf("%+v is bool\n", v)
		default:
			fmt.Printf("unknown type: %v", v)
		}
	}
}
