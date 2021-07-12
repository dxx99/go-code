package main_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 同一类型不同的值，sizeof返回的值相同
func TestUnsafeSizeOf(t *testing.T)  {
	a := "a"
	b := "bbbbbbbbbbb"
	fmt.Printf("size of a is -> %v\n", unsafe.Sizeof(a))
	fmt.Printf("size of b is -> %v\n", unsafe.Sizeof(b))
	// output:
	//size of a is -> 16
	//size of b is -> 16
}

func TestUnsafeExample(t *testing.T)  {
	type myStruct struct {
		a int64
		b bool
		c string
	}

	var x myStruct
	m, n := unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(m, n)	//16 32

	// 在编译器做的处理
	fmt.Println(unsafe.Alignof(x.a))	// 8
	fmt.Println(unsafe.Alignof(x.b))	// 1
	fmt.Println(unsafe.Alignof(x.c))	// 8

	fmt.Println(unsafe.Offsetof(x.a))	// 0
	fmt.Println(unsafe.Offsetof(x.b))	// 8
	fmt.Println(unsafe.Offsetof(x.c))	// 16
}

func TestUnsafeExample2(t *testing.T)  {
	type T struct {
		c string
	}

	type S struct {
		b bool
	}

	var x struct{
		a int64
		*S	// 隐似指针类型
		T	// 隐似非指针
	}

	fmt.Println(unsafe.Offsetof(x.a))	// 0
	fmt.Println(unsafe.Offsetof(x.S))	// 8
	fmt.Println(unsafe.Offsetof(x.T))	// 16

	fmt.Println(unsafe.Offsetof(x.c))
}

func TestUnsafeMemoryCollect(t *testing.T)  {
	createInt := func() *int {
		return new(int)
	}

	x, y, z := createInt(), createInt(), createInt()
	var x1 = unsafe.Pointer(y)				// 和y一样引用着同一个值
	var x2 = uintptr(unsafe.Pointer(z))




	// uintptr值可以参与算术运算

	x2 += 2
	fmt.Println(x2)
	x2--
	x2--

	*x = 1
	*(*int)(x1) = 5			// 非类型安全的指针转换成类型安全的指针，然后赋值
	*(*int)(unsafe.Pointer(x2)) = 3

	fmt.Println(*x, *y, *z, *(*int)(x1), x2)


}

