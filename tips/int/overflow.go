package main

import (
	"fmt"
)

func main() {

	a := uint8(1)
	b := uint8(2)
	fmt.Println(a-b)	//output 255

	//a := uint(1) - uint(2)	//./overflow.go:6:15: constant -1 overflows uint

	c := uint8(0)
	c--
	fmt.Println(c)		//output: 255  出现溢出

	var d int8 = -1
	e := -128/d		//编译器会断言e的类型为-128 -- 127之间
	fmt.Println(e)	//output: -128  因为结果为128 然后溢出，得到-128
}
