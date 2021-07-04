[TOC]
## [官方相关问题汇总](https://golang.org/doc/faq)

### 由来
#### 项目的目的？创建的原因？祖先是谁？
- 利用多核cpu, 为并发或并行提供支持，自动垃圾收集，严格依赖规范
- 高效编译，高效执行，易于编程
- c语言 + Newsqueak(limbo)并发

#### 设计原则？
- 实现效率，安全和流动性
- 没有前向声明和头文件
- 一切都只声明一次
- 初始化具有表现力，自动且易于使用
- 关键字的语法简洁明了
- 保持概念正交
  - 可以为任何类型实现方法
  - 结构代表数据，而接口代表抽象
  
  
### 设计
#### runtime
- 每个程序的一部分
- 实现go语言的垃圾收集，并发，堆栈管理和其他关键特性，类似于libc库

#### 为啥没有泛型
- [泛型建议](https://golang.org/issue/43651) 已被接纳
- 预计在1.18版本可以使用
- 泛型很方便，但是代价是类型系统和运行时的复杂性
- [泛型方案提案](https://github.com/golang/go/issues/15292)

#### 为啥没有异常处理
- 将异常耦合到控制结构(try-catch-finally)，会导致代码错综复杂，会将过多的普通错误(eg. 无法打开文件)标记为异常
- go通过多值返回在不重载返回值的情况下报告错误变得容易
- go还有内置函数，用于发送信号并从真正的异常情况中恢复
- 恢复机制： 在发生错误后被删除的函数状态的一部分执行，不需要额外的控制结构
- [错误处理](https://blog.golang.org/defer-panic-and-recover)
- [干净处理错误](https://blog.golang.org/errors-are-values)

#### 为啥没有断言
- 不希望程序员将断言用作依赖，以避免考虑正确的错误处理和报告
- 没有assert关键字，使用 val.(T) 来对变量进行断言
- reflect.TypeOf(val)  查看变量的类型

#### 为啥要使用csp上构建并发
- 由于其他语法的[复杂设计](https://en.wikipedia.org/wiki/POSIX_Threads)
- 过多的细节：互斥锁，条件变量，[内存屏障???](???)
- 为并发提供高级语言支持的最成功模型(csp)

#### 为啥使用goroutine而不是线程
- goroutine的设计会使并发易于使用
- 独立执行的函数-协程-复用到一组线程上，更加容易的处理阻塞
- 更加轻量，除了堆栈内存之外，几乎没有开销
- 为了使堆栈变小，go的runtime使用可调整大小的[有界堆栈???](???)
- runtime会自动增加(收缩)用于存储堆栈的内存
- 总结：轻量-并发容易

#### 为什么没有将map定义为原子的？
- 要求所有的map操作加互斥锁会减慢大多数程序的速度
- 不受控制的map会导致程序崩溃
- [原子操作](https://blog.csdn.net/codragon/article/details/112526621)

#### 为啥要内置map
- 它是一种强大而重要的内置数据结构，提供具有语法支持的出色实现，使编程更加愉快
- map足够强大，可以为绝大多数用途服务

#### 为啥map不允许切片作为键
- 映射查找需要一个相等运算符，切片没有实现
- slice，map，function 不可以，
-  数字、string、bool、array、channel、指针可以，以及 包含前面类型的 struct 可以
- 为结构和数组定义了相等性，因此此类类型可以用作映射键

#### 为啥map,slice,channel是引用，而array是值类型
- 指针与值的严格分离使语言更难使用，将这些类型改变为相对共享数据类型结构的引用
- 对性能产生一定的影响，但成为一种更有生产力，更舒适的语言
- ？？？

#### 为啥没有指针算法
- 安全
- 可以简化垃圾收集的实现


### 类型
#### go是面向对象语言吗？
- 是或否，有类型或方法，并允许面向对象编程风格，但没有类型层次结构
- 接口 类型嵌套
- 面向对象三大语言特点
  - 封装：通过首字母大小写来实现是否能被导出
  - 继承：通过组合实现
  - 多态：通过接口实现
  
#### 如何获取方法的动态调度
- 动态调度方法的唯一方法是通过接口。
- 结构体或任何其他具体类型上的方法总是静态解析的。

#### 为啥不支持重载
- 相同的函数名，不同的函数签名有时很有用，但实践会令人困惑和脆弱，仅按名称匹配并要求类型的一致性是 Go 类型系统中一个主要的简化决定。
- 由于重载不是必须的，没有，系统反而更加简单

#### 为啥没有implements声明
- go通过实现接口的所有方法签名来表示实现了该接口
- go接口的语义使go具有敏姐，轻量级的主要原因之一

#### 如果保证某个对象实现了某个接口
```go
package main
type T struct{}
type I interface {}
var _ I = T{}           // Verify that T implements I.
var _ I = (*T)(nil)     // Verify that *T implements I.
```

#### []T类型是否可以覆盖给[]interface{} 
- 可以
- [示例](../../tips/interface/interface-implements/interface_to_t.go)

#### 如何T1和T2有相同的底层结构，可以直接将[]T1转换成T2?
- 不可以
- 复合数据的类型转换没法直接使用转换
- [示例](../../tips/interface/interface-implements/convert_T_same_underlying_type.go)

#### 为什么我的nil错误值不等于nil
- 当一个interface的value和type都是unset的时候，它才等于nil
- 可以通过reflect来看类型和值
- [示例](../../tips/nil/nil_not_equal_nil.go)

#### 为啥不提供隐似类型转换
- C中数字类型之间自动转换的便利性被它引起的混乱所抵消
- 什么时候表达式是无符号的？价值有多大？会溢出吗？
- 结果是否可移植，独立于执行它的机器？它也使编译器复杂化；
- “通常的算术转换”不容易实现，并且跨架构不一致。

#### go常量是如何工作的？
- go对不同数值类型的变量之间的转换很严格，但语言中的常量要灵活得多
- [官方博客](https://blog.golang.org/constants)

### 指针与分配
#### 函数的参数什么时候按值传递
- go的一切都是按值传递，一个函数总是得到一个被传递的东西的副本，就想赋值语句赋值一样
- 传递指针值会生成一个指针的副本，但不会生成指针指向的数据副本
- map与slice的传递类似于指针
- 如果interface{}

#### 什么时候使用指向接口的指针？
- 几乎从不
- 编译器会很难理解这种错误

#### 我们应该定义值的方法还是指针的方法
```go
package main
type MyStruct struct {}
func (s *MyStruct) pointerMethod() { } // 指针方法
func (s MyStruct) valueMethod() { } // 值的方法
```
- 该方法是否需要修改接收器，如果是，则必须使用指针
  - slice与map有点特殊，如果要改切片长度，必须使用指针
- 效率考虑，如果结构体很大，接收是指针性能更好
- 一致性问题，如果使用指针接收器，其他的方法也要使用
- 除非方法的语义需要指针，否则值接收器是高效且清晰的

#### 为什么 T 和 *T 有不同的方法集
- 接口值包含指针*T, 方法调用可以通过取引用指针来获取值，反之不能
- 编译器可以将值的地址传递给方法的情况下，如果方法修改了值，则更改将在调用者中丢失。

#### new和make之间的区别
- new分配内存
- make初始化slice, map, channel
- [详解](https://golang.org/doc/effective_go#allocation_new)

#### int在64位机器上的大小
- int和uint为实现特定的，但给定的平台上彼此相同
- 为了平台的移植性，应使用显示的int64/int32

#### 如果知道一个变量是分配在堆上
- 看函数返回之后，该变量是否未被引用
- 如果变量非常大，将会存储到堆上
- go build -gcflags -m 去看编译的代码

#### 闭包作为 goroutine 运行会发生什么？
```go
package main

import "fmt"

func main() {
  done := make(chan bool)
  values := []string{"a", "b", "c"}
  for _, v := range values {
    go func() {
      fmt.Println(v)
      done <- true
    }()
  }
  // wait for all goroutines to complete before exiting
  for _ = range values {
    <-done
  }
}
```
- 因为循环的每次迭代都使用变量的相同实例v，因此每个闭包共享该单个变量
- [详细](https://golang.org/doc/faq#closures_and_goroutines)

#### 没啥没有?:运算符
- 不想过于频繁的创建难以理解的复杂表达式
- 通过if else就可以实现。if-else虽然长，但更清晰


### 包与测试
#### 如果创建多个文件包
- 将包的所有源文件自己放在一个目录中
- 源文件可以随意引用不同文件中的项目，不需要前向声明或头文件
- 除了被拆分成多个文件之外，该包将像单个文件包一样进行编译和测试。

#### 如何编写单元测试
- *_test.go与包源文件同一个目录
- 需要导入 import "testing" [testing框架](https://golang.org/pkg/testing/)
- 函数名要以Test*开头
- 运行时使用[go test](https://golang.org/cmd/go/#hdr-Test_packages)
- [详细文档](https://golang.org/doc/code)


### 并发
#### 哪些操作时原子的？互斥的？
- [内存模型](https://golang.org/ref/mem)
- 低级同步原语, [sync](https://golang.org/pkg/sync/), [atomic](https://golang.org/pkg/sync/atomic/)
- 高级别的同步原语, channel

#### 为什么我的程序使用更多的cpu时不能运行的更快
- 取决于它正在解决的问题，连续的问题不能通过更多的cpu来加速
- 在使用多个操作系统线程，花费更多的时间进行同步通信，而不是进行有效的计算，可能会遇到性能下降
- 线程间传递数据涉及到切换上下文，有显著的成本，并且该成本会随着cpu的增加而增加
- [参考](https://blog.golang.org/waza-talk)

#### 如何控制cpu的数量
- 可同时运行的goroutines由 [runtime.GOMAXPROCS](https://golang.org/pkg/runtime/#GOMAXPROCS) 来控制, 默认使用cpu内核数量

#### 为啥没有goroutine id
- 为了在编写并发代码时可以使用完整的go语言，防止使用者与该goroutine关联
- 特殊线程或 goroutine 的存在迫使程序员扭曲程序，以避免因无意中操作错误的线程而导致崩溃和其他问题。

#### 为啥要进行垃圾收集？成本会不会太高？
- 垃圾回收使接口更简单，因为他们不需要指定如何跨接口管理内存
- 当前的实现是一个标记和清除收集器
- 如果机器是多处理器，则收集器与主程序并行运行在单独的 CPU 内核上。
- [参考](https://blog.golang.org/ismmkeynote)

### 代码编写
#### 学习途径
- [标准库](https://golang.org/pkg/)
- [cmd指令](https://golang.org/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol)
- [高效go](https://golang.org/doc/effective_go)
- [如何成为贡献者](https://golang.org/doc/contribute)
- [如何code review](https://github.com/golang/go/wiki/CodeReviewComments)
- [mod使用](https://golang.org/doc/tutorial/create-module)




