package main_test

import (
	"fmt"
	"reflect"
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

	// 此时，即使z指针值所引用的int值的地址仍旧存储 ，在x2值中，但是此int值已经不再被使用了，所以垃圾 回收器认为可以回收它所占据的内存块了。
	//另一方面， x和x1各自所引用的int值仍旧将在下面被使用。

	// uintptr值可以参与算术运算
	x2 += 2
	fmt.Println(x2)	// 824633992922
	x2--
	x2--

	*x = 1
	*(*int)(x1) = 5			// 非类型安全的指针转换成类型安全的指针，然后赋值
	*(*int)(unsafe.Pointer(x2)) = 3

	fmt.Println(*x, *y, *z, *(*int)(x1), x2)
}


func TestUnsafeLiveArea(t *testing.T)  {
	type s struct{
		x int
		y *[1<<32]byte
	}
	a := s{
		x: 0,
		y: new([1<<32]byte),
	}
	p := uintptr(unsafe.Pointer(&a.y))
	fmt.Println("before: ", p)
	//... 使用a.x a.y

	// 一个聪明的编辑器能够察觉到值a.y将不会再被用到
	// 所以认为a.y值所占的内存块可以被回收了

	*(*byte)(unsafe.Pointer(p)) = 1	// 危险操作
	fmt.Println("after: ", p)
	fmt.Println(a.x)
}

func TestUnsafeMyStringToString(t *testing.T)  {
	type MyString string
	a := []MyString{"PHP", "Python", "Golang"}
	b := *(*[]string)(unsafe.Pointer(&a))
	b[1] = "Rust"


	fmt.Printf("a'type = %v, b'type = %v \n ", reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Println("a =", a, "b =",b)
}





