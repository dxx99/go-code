package main

import (
	"sync"
	"testing"
)

// go 上下午切换
// 设置cpu为1 防止cpu上下文切换影响
// 基本上是操作系统上下文切换速度的10倍
//go:generate go test -bench=. -cpu=1 go-context-switch_test.go
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin //1 将两个协程阻塞在这里，不希望设置数据，影响goroutine上下文切换的度量
		for i := 0; i < b.N; i++ {
			c <- token //2 这里向接收者发送数据
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin //1
		for i := 0; i < b.N; i++ {
			<-c //3
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() //4 开启启动计时器
	close(begin)   //5  唤醒接收与发送者阻塞的协程
	wg.Wait()

	//output:
	//goos: darwin
	//goarch: amd64
	//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	//BenchmarkContextSwitch   8333444               141.0 ns/op
	//PASS
	//ok      command-line-arguments  1.328s
}
