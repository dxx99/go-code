## 非类型安全指针

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
- 一个值的地址在程序运行中可能改变
- 一个值的声明范围可能并没有代码中看上去大
- *unsafe.Pointer是一个类型安全指针类型