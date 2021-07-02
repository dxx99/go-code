package main

import (
	"fmt"
	"reflect"
)

type A interface {

}

type B struct {

}

// 指针类型
var a *B

// nil not equal to nil
func main() {
	fmt.Println( a == nil)	//true
	fmt.Println( a == (*B)(nil))	//true
	fmt.Println( (A)(a) == (*B)(nil))	//true

	fmt.Println( (A)(a) == nil)	  //false ???  因为只有当一个interface的value和type都是unset的时候，它才等于nil, 上面的interface{}的类型是*B
	fmt.Println(reflect.TypeOf((A)(a)))	//*main.B
	fmt.Println(reflect.ValueOf((A)(a)))	//<nil>


}
