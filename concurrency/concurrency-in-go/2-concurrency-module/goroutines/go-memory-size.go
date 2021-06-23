package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c //1 需要挂起goroutine, 让其永不退出
	}

	const numGoroutines = 1e4 //2

	wg.Add(numGoroutines)
	before := memConsumed() //3 创建分区之前的内存消耗量
	for i := 0; i < numGoroutines; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed() //4 消耗之后的内存量

	fmt.Printf("before=%.3fkb, after=%.3fkb, %.3fkb\n", float64(before)/numGoroutines/1e3, float64(after)/numGoroutines/1e3, float64(after-before)/numGoroutines/1000)

}
