package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

var (
	stop int32
	sum time.Duration
	count int64
)

func main() {
	f, _ := os.Create("./base/gc/memory/trace2.out")
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop) == 0 {
			t  = time.Now()
			runtime.GC()
			sum += time.Since(t)
			count++
		}
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count))
	}()

	//concat()
	concatV2()
	atomic.StoreInt32(&stop, 1)
}

func concat() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 8; j++ {
			go func() {
				s := "GO GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
 			}()
		}
	}
}

// 控制协程数
func concatV2()  {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(8)
		for j := 0; j < 8; j++ {
			go func() {
				defer wg.Done()
				s := "GO GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
			}()
		}
		wg.Wait()
	}
}
