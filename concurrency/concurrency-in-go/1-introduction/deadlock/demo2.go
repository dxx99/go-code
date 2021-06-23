package main

import (
	"fmt"
	"sync"
	"time"
)

type DemoVal struct {
	sync.Mutex
	Num int
}

var wg sync.WaitGroup

var printSum = func(v1, v2 *DemoVal) {
	defer wg.Done()
	v1.Lock()
	defer v1.Unlock()

	//模拟工作时间
	time.Sleep(time.Second * 2)

	v2.Lock()
	defer v2.Unlock()

	// 输出v1 + v2的结果
	fmt.Printf("v1 + v2 = %v\n", v1.Num+v2.Num)
}

func main() {
	var a, b DemoVal
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

}
