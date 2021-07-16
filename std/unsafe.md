## 非类型安全指针
- 主要目的
  - 类型转换
  - 通过uintptr, 支持算术运算

### 前言
- go指针的限制(也叫类型安全的指针)
    - 不支持算术运算
    - 任意两个指针值，很有可能不能相互转换
- 特点
    - 撸码心智成本低
    - 灵活性不够，导致不能撸出高效代码
- 非类型安全指针(unsafe)
    - 限制少
    - 出现bug的概率高
    - 不受[go1兼容性保护](https://golang.google.cn/doc/go1compat)，

### 关于unsafe标准库
一种特殊的类型，unsafe.Pointer 声明定义如下：
```
type Pointer *ArbitraryType
```
讲解：
- ArbitraryType 仅仅暗示unsafe.Pointer类型可以被转换为任意类型安全指针(反之也可)
- 类似于c语言的void*
- 非类型安全的指针是指底层类型为unsafe.Pointer的类型，零值用nil标识

go1.7以前版本标准库的三个函数
- Alignof
    - 用来取一个值在内存中地址对齐保证
    - 同一类型的值，在结构体或非结构体中对齐的保证可能不相同，与编译器实现有关
- Offsetof
    - 用来取一个结构体值的某个字段地址相对于此结构体值的地址偏移
    - 在一个程序中，对同一结构体类型的不同值的对应相同字段，此函数的返回是相同的
- Sizeof
    - 取值的size, 也叫类型size
    - 同一类型不同的值，此函数的返回值总数相同的
- 注意：
    - 返回值都是 uintptr
    - 这三个函数的调用总是在编译时刻被估值，结果为uintptr常量
    - 传给offsetof 函数的实参必须是一个字段选择器value.filed, 此选择器表示一个内嵌字段，不能使用隐似字段
  
go1.7新添加一个类型和两个函数
- 类型
```
type IntegerType int    
```

- 函数
```
func Add(ptr Pointer, len IntegerType) Pointer
func Slice(ptr *ArbitaryType, len IntegerType) []ArbitraryType
```
  - add 此函数在一个非安全指针表示的地址上添加一个偏移量，然后返回表示新地址的一个指针
  - slice 用来从一个任意安全指针派生一个指定长度的切片

### 非类型安全指针的转换
- 一个类型安全指针可以被显示转换为一个非类型安全指针类型，反之亦然
- 一个uintptr值可以被显示转换为一个非类型安全指针类型，反之亦然。

**注意**
- 一个nil非类型安全指针类型不应该被转换为uintptr并进行算术运算后再转换回来

**目的**
- 实现任意两个类型安全指针转换为对方的类型
- 也可以将一个阿全指针值和一个uintptr指转换为对方的类型

**危害**
- 这些转换在编译时刻是合法的，但在运行时并非合法和安全的
- 会摧毁go的类型系统(不包括非类型安全指针部分)精心设立的内存安全屏障

### 必须知道的一些事实
- 非类型安全指针值是指针但uintptr值是整数
  - 每一个非零安全或者不安全指针均引用这另一个值，但一个uintptr值并不引用任何值，被作为整数存储
  - uintptr值是一个整数，可以参与算术运算
  - 值与值之间和内存块与值之间的引用关系是通过指针来表征的
- 不再被使用的内存块的回收时间点是不确定的
  - [例子](./test/upsafe_test.go) - TestUnsafeMemoryCollect
- 一个值的地址在程序运行中可能改变
  - 一个协程的栈大小发送改变时，开辟在此栈上的内存块需要移动，从而相应的的值地址将改变
- 一个值的声明范围可能并没有代码中看上去大
  - [例子](./test/upsafe_test.go) - TestUnsafeLiveArea
- *unsafe.Pointer是一个类型安全指针类型

### 如何正确使用非类型安全指针
[标准库列出了六种使用模式](https://golang.google.cn/pkg/unsafe/#Pointer)
1. 将*T1的值转换成非类型安全指针，然后将此非类型安全指针转成类型*T2

```go
package unsafe_test
import (
  "fmt"
  "reflect"
  "testing"
  "unsafe"
)
func TestUnsafeMyStringToString(t *testing.T) {
  type MyString string
  a := []MyString{"PHP", "Python", "Golang"}
  b := *(*[]string)(unsafe.Pointer(&a))
  b[1] = "Rust"

  fmt.Printf("a'type = %v, b'type = %v \n ", reflect.TypeOf(a), reflect.TypeOf(b))
  fmt.Println("a =", a, "b =", b)
}
```

2. 将一个非类型安全的指针转换成一个uintptr值，然后使用uintptr
- 不是很有用，使用的比较少

3. 将一个非类型安全的指针转成uintptr, 然后用uintptr进行算术运算，再将算术运算的结构uintptr转成非类型安全指针
```go
package main

import "fmt"
import "unsafe"

type T struct {
	x bool
	y [3]int16
}

const N = unsafe.Offsetof(T{}.y)
const M = unsafe.Sizeof(T{}.y[0])

func main() {
	t := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	ty2 := (*int16)(unsafe.Pointer(uintptr(p)+N+M+M))
	fmt.Println(*ty2) // 789
}
```
- 1.7之后用的unsafe.Add函数来完成。

4. 将非类型安全指针值转换为uintptr值，并传递给syscall.Syscall函数调用
- 什么叫系统调用？

5. 将reflect.Value.Pointer 或 reflect.Value.UnsafeAddr的 uintptr的返回值转换成非类型安全的指针
- reflect标准库包中的Value类型的Pointer和UnsafeAddr方法都返回一个uintptr值，而不是一个unsafe.Pointer值
- 这样设计的目的是避免用户不引用unsafe标准库包就可以将这两个方法的返回值（如果是unsafe.Pointer类型）转换为任何类型安全指针类型

*这样的设计需要我们将这两个方法的调用的uintptr结果立即转换为非类型安全指针。
否则，将出现一个短暂的可能导致处于返回的地址处的内存块被回收掉的时间窗。
此时间窗是如此短暂以至于此内存块被回收掉的几率非常之低，
因而这样的编程错误造成的bug的重现几率亦十分得低。*
```
// 这样是安全的
p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))

u := reflect.ValueOf(new(int)).Pointer()
// 在这个时刻，处于存储在u中的地址处的内存块
// 可能会被回收掉。
p := (*int)(unsafe.Pointer(u))
```  

6. 将一个reflect.SliceHeader 或者 reflect.StringHeader值的Data字段转换为非类型安全指针，以及其逆转换
```go
package main

import "fmt"
import "unsafe"
import "reflect"

func main() {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s) // Golang
	// 现在，字符串s和切片a共享着底层的byte字节序列，
	// 从而使得此字符串中的字节变得可以修改。
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) // Google
}
```

```go
package main

import (
  "fmt"
  "reflect"
  "unsafe"
)

func main() {
  a := [6]byte{'G', 'o', '1', '0', '1'}
  bs := []byte("Golang")
  hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
  hdr.Data = uintptr(unsafe.Pointer(&a))

  hdr.Len = 2
  hdr.Cap = len(a)
  fmt.Printf("%s\n", bs) // Go
  bs = bs[:cap(bs)]
  fmt.Printf("%s\n", bs) // Go101
}
```



