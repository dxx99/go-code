package main

import (
	"fmt"
	"sync"
)


//go:generate go run go-once.go
// once.Do(func(){}) 只执行一次，通过原子操作，给对应的结构体加上一个1，然后判断是否已执行
//
// 通过grep统计sync.Once 在标准库中使用的次数
//grep -ir sync.Once $(go env GOROOT)/src | wc -l
func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()

	fmt.Println("count is ", count)
}
