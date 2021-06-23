package main

import (
	"fmt"
	"sync"
)

// 对象池模式的并发安全实现
// 在较高的层次上，池模式是一种创建和提供固定数量可用对象的方式，它通常用于约束创建资源昂贵的事务(eg. 数据库连接)
func main() {
	myPool := &sync.Pool{New: func() interface{} {
		fmt.Println("create new instance.")
		return struct {
		}{}
	}}

	myPool.Get()	//1 这里我们调用get方法，将调用在池中定义的new函数，因为实例尚未实例化
	ins := myPool.Get()
	fmt.Println(ins)
	myPool.Put(ins) //2 将先前检索的实例放回池中，这时实例的可用数量为1
	myPool.Get()

	myPool.Get()  //3对象池中没有实例，这个时候就要创建实例
}
