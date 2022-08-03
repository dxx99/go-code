package main

import (
	"fmt"
	"reflect"
)

func main() {
	//modifyVal()
	//modifyValV2()
	//modifyValV3()

	// 类型比较
	//diffType()

	// 判断两个类型深度相等
	equal()
}


// 通过反射修改变量的值
func modifyVal()  {
	x := 3.4
	p := reflect.ValueOf(x)
	p.SetFloat(3.1)		// panic: reflect: reflect.Value.SetFloat using unaddressable value
}

func modifyValV2()  {
	x := 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p", p.Type())				// type of p  *float64
	fmt.Println("setAbility of p", p.CanSet())	// setAbility of p  false
	p.SetFloat(3.1)									// panic: reflect: reflect.Value.SetFloat using unaddressable value
}

// 如果想要操作原变量，反射变量Value必须要hold住原变量的地址才行
func modifyValV3()  {
	x := 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(3.1)
	fmt.Println(v.Interface())		//3.1
	fmt.Println(x)					//3.1
}

// 不同的类型
func diffType() {
	var i int
	type MyInt int
	var j MyInt
	i = 1
	j = 1
	//fmt.Println(i == j)	// ./main.go:46:16: invalid operation: i == j (mismatched types int and MyInt)
	fmt.Println(i, j)
}

type MyInt int
type YourInt int
func equal()  {
	m, y :=  MyInt(1), YourInt(1)
	ans := reflect.DeepEqual(m, y)
	fmt.Println(ans)
}


