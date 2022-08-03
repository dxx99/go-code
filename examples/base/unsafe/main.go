package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//limitOne()
	//limitTwo()

	//getSliceLen()
	//getMapLen()

	//getStructAttr()
	getStructAttrV2()
}

// 指针的一些限制

// 限制一：不能进行一些数学运算
func limitOne()  {
	a := 1
	p := &a

	//p++			//./main.go:16:3: invalid operation: p++ (non-numeric type *int)
	fmt.Println(p)
}


// 限制二：不同类型的指针不能相互转换
func limitTwo()  {
	a := 100
	var f *float64
	// f = &a	//./main.go:26:4: cannot use &a (type *int) as type *float64 in assignment

	fmt.Println(a, f)
}

//
/**
type slice struct {
	array unsafe.Pointer	// 元素指针
	len 	int				// 长度
	cap 	int				// 容量
}
 */
func getSliceLen() {
	s := make([]int, 9, 20)
	// 思路:
	//	先转换成uintptr, 进行指针偏移量加8，也就是移动到结构体len字段上
	// 	然后再将uintptr转换成 Pointer类型，这样通过类型强转成制定类型也就是(*int)指针类型
	//	然后通过*取到对应的数据
	l := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(l, len(s))
}

func getMapLen() {
	m := make(map[int]int, 10)
	m[1] =1
	m[2] =2

	// 分析
	// count为二级指针
	// &m -> Pointer -> **int -> int
	l := **(**int)(unsafe.Pointer(&m))
	fmt.Println(l, len(m))
}

type Programmer struct {
	name string
	language string
}



func getStructAttr()  {
	p := Programmer{
		name:     "dxx99",
		language: "go",
	}

	name := (*string)(unsafe.Pointer(&p))
	*name = "wahaha"
	fmt.Println(p)

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "golang"
	fmt.Println(p)
}


type ProgrammerV2 struct {
	name string
	age int
	language string
}

// 通过unsafe.Sizeof()函数处理
func getStructAttrV2()  {
	p := ProgrammerV2{
		name:     "dxx99",
		age:      33,
		language: "go",
	}

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string("")) ))
	*lang = "golang"
	fmt.Println(p)
}
