package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit chan int = make(chan int)

//由于多个协程同时往stdout写入， 纵然前面执行再快，最后还是要单线程输出到屏幕，导致的性能问题，如果fmt换为其他操作，就可以显示出GOMAXPROCS的好处。
func loop()  {
	for i := 0; i < 10000; i++ {
		if i % 10 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	quit <- 0
}

// 设置goMaxProcs带来的性能问题
func main() {

	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(1)
	beginTime := time.Now()
	for i := 0; i < 1000; i++ {
		go loop()
	}
	for i := 0; i < 1000; i++ {
		<-quit
	}
	useTime := time.Since(beginTime)
	fmt.Printf("cpuNum: %d, 运行时间: %v\n", cpuNum, useTime)

}
