package main

import (
	"fmt"
	"reflect"
)

type T1 int
type T2 int

// 覆盖相同底层类型的变量
func main() {

	s1 := []T1{1,2,3}
	s2 := make([]T2, len(s1))
	for k, v := range s1 {
		//s2[k] = v	// NOT ok
		s2[k] = T2(v)	//OK
	}
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(s2[0]))	//output: main.T2

	// 符合类型的元素是没法直接转换的
	//var st1 []T1
	//var st2 = ([]T2)(st1)		// not ok


}
