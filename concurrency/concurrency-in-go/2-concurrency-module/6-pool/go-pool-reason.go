package main

import (
	"fmt"
	"sync"
)

// 使用池的原因：那么为什么要使用一个池，而不是实例化对象呢？ Go有一个垃圾收集器，所以实例化的对象将被自动清理。 重点是什么？ 考虑这个例子
//
func main() {
	var poolNums int
	calcPool := &sync.Pool{
		New: func() interface{} {
			poolNums++
			mem := make([]byte, 1024)
			return &mem //1 存储字节切片指针
		},
	}

	for i := 0; i < 4; i++ {
		calcPool.Put(calcPool.New())
	}

	const workerNums = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerNums)
	for i := 0; i < workerNums; i++ {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // 断言使用字节切片指针
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.\n", poolNums)
}
