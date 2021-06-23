package main

import (
	"fmt"
	"sync"
)

// 多个goroutine进行加减操作，通过互斥锁来保证数据安全性
// 被锁定部分是程序的性能瓶颈，进入和退出锁定的成本有点高，因此人们通常尽量减少锁定涉及的范围
func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	var count int
	increment := func() {
		l.Lock()
		defer l.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		l.Lock()
		defer l.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Printf("increment and decrement finish. count is %d\n", count)

}
