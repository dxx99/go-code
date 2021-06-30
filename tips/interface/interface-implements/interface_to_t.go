package main

import (
	"fmt"
	"reflect"
)

// 将[]T类型覆盖到[]interface{}类型中
func main() {
	a := []int{1,2,3,4}

	sliceInterface := make([]interface{}, len(a))
	for key, item := range a {
		sliceInterface[key] = item
	}
	sliceInterface = append(sliceInterface, "hello")
	fmt.Println(sliceInterface)

	fmt.Println(reflect.TypeOf(sliceInterface[4]))
	fmt.Println(reflect.TypeOf(sliceInterface[3]))
}
